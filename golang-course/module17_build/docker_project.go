//go:build ignore

package main

import (
	"fmt"
	"net/http"
	"os"
)

// ─────── DOCKER PROJECT ───────
// This is a simple web server designed to be run inside a Docker container.
// Check out the `Dockerfile` in this directory to see how Go apps are typically
// containerized using Multi-Stage Builds!

// We can inject this version during `go build` using ldflags!
var Version = "development"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go! Running version: %s\n", Version)
	})

	fmt.Printf("Starting server on port %s (Version: %s)...\n", port, Version)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
