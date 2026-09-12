package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ─────── MIDDLEWARE PROJECT ───────
// Middleware in Go is just a function that takes an http.Handler and returns an http.Handler.
// It wraps around your core business logic, allowing you to run code BEFORE and AFTER
// the main handler executes (e.g., Logging, Authentication, Rate Limiting).

// 1. A basic handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate some work
	time.Sleep(20 * time.Millisecond)
	w.Write([]byte("Hello, World!"))
}

// 2. Logging Middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Run code BEFORE the next handler
		log.Printf("--> %s %s", r.Method, r.URL.Path)
		
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
		
		// Run code AFTER the next handler
		duration := time.Since(start)
		log.Printf("<-- %s %s took %v", r.Method, r.URL.Path, duration)
	})
}

// 3. Authentication Middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return // Short-circuit! Do NOT call next.ServeHTTP
		}
		
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("--- Middleware Pattern Demo ---")
	
	mux := http.NewServeMux()
	
	// Our raw handler
	var handler http.Handler = http.HandlerFunc(helloHandler)
	
	// Wrap it in middlewares (from inside out)
	// First Auth, then Logging.
	// When a request comes in: Logging -> Auth -> helloHandler -> Auth -> Logging
	handler = AuthMiddleware(handler)
	handler = LoggingMiddleware(handler)
	
	mux.Handle("/secure", handler)
	
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	
	go func() {
		fmt.Println("Server listening on http://localhost:8081")
		srv.ListenAndServe()
	}()
	
	time.Sleep(500 * time.Millisecond)
	
	// Test unauthorized
	fmt.Println("\nAttempting unauthorized request...")
	resp1, _ := http.Get("http://localhost:8081/secure")
	fmt.Println("Response Status:", resp1.Status)
	
	// Test authorized
	fmt.Println("\nAttempting authorized request...")
	req, _ := http.NewRequest("GET", "http://localhost:8081/secure", nil)
	req.Header.Set("Authorization", "secret-token")
	resp2, _ := http.DefaultClient.Do(req)
	fmt.Println("Response Status:", resp2.Status)
}
