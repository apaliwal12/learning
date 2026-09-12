package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ─────── 1. HTTP CLIENT ───────
// Go's `net/http` package provides a robust, production-ready HTTP client.
// By default, `http.Get` uses `http.DefaultClient`, which DOES NOT HAVE A TIMEOUT.
// NEVER use `http.DefaultClient` in production. Always configure your own!

func demonstrateHTTPClient() {
	fmt.Println("\n--- HTTP Client ---")
	
	// Create a client with a sane timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	
	// We use httpbin.org which is great for testing HTTP requests
	req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	
	// Add custom headers
	req.Header.Add("User-Agent", "Go-Mastery-Guide/1.0")
	req.Header.Add("Accept", "application/json")
	
	// Execute the request
	fmt.Println("Sending GET request to httpbin.org...")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return
	}
	
	// CRITICAL: You must close the response body, or you will leak connections!
	defer resp.Body.Close()
	
	fmt.Printf("Status Code: %d %s\n", resp.StatusCode, resp.Status)
	
	// Read the body (usually you would use json.NewDecoder here!)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err == nil {
		// Truncate output for demo
		out := string(bodyBytes)
		if len(out) > 200 {
			out = out[:200] + "...\n}"
		}
		fmt.Printf("Response Body:\n%s\n", out)
	}
}

// ─────── 2. HTTP CLIENT WITH CONTEXT ───────
// We can use contexts to cancel requests midway if they take too long,
// overriding the Client's default timeout.

func demonstrateHTTPContext() {
	fmt.Println("\n--- HTTP Request with Context ---")
	
	// This endpoint delays for 3 seconds
	req, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/delay/3", nil)
	
	// Create a context with a 1-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	// Attach the context to the request
	req = req.WithContext(ctx)
	
	client := &http.Client{} // No timeout here, relying entirely on context
	
	fmt.Println("Sending request (server delays 3s, our context timeout is 1s)...")
	_, err := client.Do(req)
	if err != nil {
		fmt.Printf("Expected error due to context timeout: %v\n", err)
	}
}
