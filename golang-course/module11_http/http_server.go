package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ─────── 3. HTTP SERVER ───────
// Go's `net/http` package includes a production-ready HTTP server.
// Handlers in Go implement the `http.Handler` interface, which has one method:
//   ServeHTTP(ResponseWriter, *Request)
//
// The `ResponseWriter` is an interface you use to write the response back to the client.
// The `*Request` is a struct containing all information about the incoming request.

// A simple handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// You can read query parameters
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	
	// Write the response
	w.WriteHeader(http.StatusOK) // 200 is default, but good to be explicit
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

// A JSON API handler
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	
	data := map[string]string{
		"message": "Welcome to the API",
		"status":  "success",
		"time":    time.Now().Format(time.RFC3339),
	}
	
	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")
	
	// Encode data directly to the ResponseWriter stream
	json.NewEncoder(w).Encode(data)
}

func demonstrateHTTPServer() {
	fmt.Println("\n--- HTTP Server (Simulated) ---")
	
	// We use the default multiplexer (router) provided by net/http
	// In production, people often use external routers like chi or gorilla/mux
	// for better path variable matching (e.g. /users/{id}), though Go 1.22
	// improved the standard library router significantly.
	
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/api/info", apiHandler)
	
	// Create a Server struct to configure timeouts (CRITICAL for production!)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  // Max time to read request headers/body
		WriteTimeout: 10 * time.Second, // Max time to write response
		IdleTimeout:  120 * time.Second, // Max time to keep Keep-Alive connections open
	}
	
	// We launch the server in a goroutine so it doesn't block our demo script
	go func() {
		fmt.Printf("Server listening on http://localhost%s\n", srv.Addr)
		// ListenAndServe always returns an error (usually http.ErrServerClosed)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()
	
	// Give the server a moment to start up
	time.Sleep(500 * time.Millisecond)
	
	// Now we'll use our client to make a request to our own server!
	fmt.Println("Making a request to our local server...")
	resp, err := http.Get("http://localhost:8080/api/info")
	if err == nil {
		defer resp.Body.Close()
		fmt.Printf("Received status: %s\n", resp.Status)
	}
}
