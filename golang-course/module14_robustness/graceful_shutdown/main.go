package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ─────── GRACEFUL SHUTDOWN PROJECT ───────
// When a Go server receives a SIGTERM (e.g., from Kubernetes scaling down)
// or SIGINT (Ctrl+C), it should NOT just die immediately.
// It should stop accepting NEW requests, finish processing CURRENT requests,
// and THEN exit.

func main() {
	fmt.Println("--- Graceful Shutdown Demo ---")
	
	// 1. Setup a basic HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simulate a slow request
		fmt.Println("Handling request...")
		time.Sleep(3 * time.Second)
		w.Write([]byte("Request finished successfully\n"))
		fmt.Println("Request complete.")
	})

	srv := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	// 2. Start the server in a goroutine
	go func() {
		fmt.Printf("Server listening on %s. Try sending a request, then pressing Ctrl+C.\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 3. Setup signal catching
	// We create a channel to receive OS signals.
	quit := make(chan os.Signal, 1)
	// We tell the os/signal package to route SIGINT and SIGTERM to our channel.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 4. Block the main goroutine until a signal is received
	sig := <-quit
	fmt.Printf("\nReceived signal: %v. Initiating graceful shutdown...\n", sig)

	// 5. Create a context with a timeout for the shutdown
	// We give the server 5 seconds to finish active requests.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 6. Call srv.Shutdown(ctx)
	// This immediately stops accepting new connections, and waits for active
	// connections to finish (up to the context timeout).
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server gracefully stopped.")
}
