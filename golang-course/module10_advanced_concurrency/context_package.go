package main

import (
	"context"
	"fmt"
	"time"
)

// ─────── 4. CONTEXT ───────
// The `context` package is the standard way to propagate cancellation signals,
// deadlines, and request-scoped values across API boundaries and between goroutines.
// It is ubiquitous in Go. Almost every I/O or network function takes a context.Context
// as its FIRST parameter.

// simulateLongDBQuery takes a context. If the context is canceled before the
// query finishes, the function aborts early.
func simulateLongDBQuery(ctx context.Context) error {
	fmt.Println("[DB] Query started (will take 50ms)")
	
	select {
	case <-time.After(50 * time.Millisecond): // Simulate work
		fmt.Println("[DB] Query completed successfully")
		return nil
	case <-ctx.Done(): // Context was canceled or timed out!
		// ctx.Err() tells us WHY it was canceled (e.g., context.Canceled or context.DeadlineExceeded)
		fmt.Printf("[DB] Query aborted: %v\n", ctx.Err())
		return ctx.Err()
	}
}

func demonstrateContextCancel() {
	fmt.Println("\n--- context.WithCancel ---")
	
	// Create a base context (context.Background() is the root of all contexts)
	// WithCancel returns a derived context and a cancel function.
	ctx, cancel := context.WithCancel(context.Background())
	
	// Launch the query
	go simulateLongDBQuery(ctx)
	
	// Main goroutine decides to cancel the query early
	time.Sleep(20 * time.Millisecond)
	fmt.Println("Main: We don't need the result anymore, canceling context...")
	
	// Calling cancel() broadcasts the cancellation signal to the context and ALL its children.
	cancel() 
	
	time.Sleep(10 * time.Millisecond) // Give DB goroutine time to print its abort message
}

func demonstrateContextTimeout() {
	fmt.Println("\n--- context.WithTimeout ---")
	
	// WithTimeout automatically calls the cancel function when the duration expires.
	// You STILL must `defer cancel()` to release resources if the function returns
	// BEFORE the timeout expires!
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	
	err := simulateLongDBQuery(ctx)
	if err != nil {
		fmt.Println("Main: DB query failed:", err)
	}
}

// ─────── 5. CONTEXT VALUES ───────
// Contexts can carry request-scoped data (e.g., Trace IDs, Authentication tokens).
// USE WITH CAUTION: Do not use context values for passing optional function parameters.
// Only use them for metadata that passes *through* layers transparently.

// We define a custom unexported type for the context key to prevent collisions.
// If we used a string like "traceID", another package might also use "traceID"
// and overwrite our value!
type contextKey string
const traceIDKey contextKey = "traceID"

func demonstrateContextValues() {
	fmt.Println("\n--- context.WithValue ---")
	
	ctx := context.WithValue(context.Background(), traceIDKey, "req-12345")
	
	// Passing it down...
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// Extract the value. It returns an `any`, so we must type-assert.
	traceID, ok := ctx.Value(traceIDKey).(string)
	if ok {
		fmt.Printf("Processing request with Trace ID: %s\n", traceID)
	} else {
		fmt.Println("No Trace ID found in context")
	}
}
