package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// ─────── OBSERVABILITY PROJECT ───────
// This project ties together HTTP Middlewares, Contexts, and Structured Logging.
// We will extract a "Trace ID" from the incoming HTTP request, put it into the
// Context, and then create a Logger that automatically includes that Trace ID
// in every log message it prints!

type contextKey string
const traceIDKey contextKey = "X-Trace-ID"

// 1. The Middleware
func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to get trace ID from headers
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			// In reality, we'd generate a UUID here if one was missing
			traceID = fmt.Sprintf("generated-%d", time.Now().UnixNano())
		}
		
		// Put it in the context
		ctx := context.WithValue(r.Context(), traceIDKey, traceID)
		
		// Create a new request with the updated context
		r = r.WithContext(ctx)
		
		next.ServeHTTP(w, r)
	})
}

// 2. The Logger Injector
// We can wrap our slog.Logger so it always looks for a TraceID in the context!
// However, standard slog doesn't pull from context automatically. We have to do it manually,
// or write a custom Handler. For simplicity, we'll write a helper function.

func getLogger(ctx context.Context) *slog.Logger {
	logger := slog.Default()
	
	if traceID, ok := ctx.Value(traceIDKey).(string); ok {
		return logger.With(slog.String("trace_id", traceID))
	}
	
	return logger
}

// 3. The Business Logic
func processOrder(ctx context.Context, orderID string) {
	log := getLogger(ctx)
	
	log.Info("Starting to process order", slog.String("order_id", orderID))
	time.Sleep(50 * time.Millisecond)
	log.Info("Order processed successfully")
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	// The logger here will have the trace_id attached!
	log := getLogger(r.Context())
	
	log.Info("Received order request")
	
	// Pass the context down!
	processOrder(r.Context(), "ORD-789")
	
	w.Write([]byte("Order processed\n"))
}

func main() {
	fmt.Println("--- Observability Demo ---")
	
	// Setup JSON logging
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	
	mux := http.NewServeMux()
	mux.Handle("/order", TraceMiddleware(http.HandlerFunc(orderHandler)))
	
	go http.ListenAndServe(":8083", mux)
	
	time.Sleep(500 * time.Millisecond)
	
	// Make a request with a specific trace ID
	req, _ := http.NewRequest("POST", "http://localhost:8083/order", nil)
	req.Header.Set("X-Trace-ID", "TRACE-12345-ABC")
	
	fmt.Println("Sending request with Trace ID...")
	http.DefaultClient.Do(req)
	
	time.Sleep(100 * time.Millisecond)
}
