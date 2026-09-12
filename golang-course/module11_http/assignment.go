package main

import (
	"context"
	"net/http"
	"time"
)

// TODO 1: Implement FetchStatus
// It should take a URL string and return the HTTP status code (int) and any error.
// You MUST use a custom http.Client with a 2-second timeout, not the default client.
func FetchStatus(url string) (int, error) {
	return 0, nil // Fix me
}

// TODO 2: Implement FetchWithContextAndTimeout
// It should take a context, a URL, and a timeout duration.
// It should create a new context with the given timeout derived from the passed context.
// It should make an HTTP GET request to the URL using the new context.
// Return the status code.
// Make sure to close the response body!
func FetchWithContextAndTimeout(ctx context.Context, url string, timeout time.Duration) (int, error) {
	return 0, nil // Fix me
}

// TODO 3: Implement a simple HTTP Handler `HealthHandler`
// It should implement http.HandlerFunc (so it takes a ResponseWriter and Request).
// It should write the status code 200 OK.
// It should write the response body: `{"status":"ok"}`
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Fix me
}
