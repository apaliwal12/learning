package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

// ─────── 1. STANDARD LOGGING ───────
// Go's standard `log` package is simple and prints to stderr by default.
// It automatically adds timestamps.
// However, it is unstructured (just plain text strings), which makes it hard
// to parse in modern log aggregation systems like Datadog, Splunk, or ELK.

func demonstrateStandardLog() {
	fmt.Println("\n--- Standard 'log' Package ---")
	
	// You can configure the default logger
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	
	log.Println("This is a standard log message.")
	log.Printf("User %d logged in from %s", 42, "127.0.0.1")
}

// ─────── 2. STRUCTURED LOGGING (slog) ───────
// In Go 1.21, the `log/slog` package was introduced.
// This is the MODERN way to log in Go.
// It outputs structured data (like JSON), which is easily indexable and searchable.

func demonstrateSlog() {
	fmt.Println("\n--- Structured Logging 'log/slog' ---")
	
	// 1. Create a JSON handler
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Log everything from Debug and above
	})
	
	// 2. Create the logger
	logger := slog.New(jsonHandler)
	
	// 3. (Optional) Set it as the default logger for the whole app
	slog.SetDefault(logger)
	
	// 4. Log with attributes!
	slog.Info("user logged in", 
		slog.Int("user_id", 42),
		slog.String("ip_address", "192.168.1.1"),
	)
	
	slog.Error("failed to connect to database",
		slog.String("db_host", "db.local"),
		slog.Int("attempt", 3),
	)
	
	// 5. Creating child loggers with pre-bound attributes
	// This is amazing for passing a logger down into a specific component (e.g. a worker)
	workerLogger := logger.With(slog.String("component", "background_worker"))
	workerLogger.Debug("job started", slog.Int("job_id", 999))
}
