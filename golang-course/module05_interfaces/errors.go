package main

import (
	"errors"
	"fmt"
)

// ─────── 6. THE ERROR INTERFACE ───────
// In Go, `error` is just a built-in interface:
// type error interface {
//     Error() string
// }
//
// You can define custom error types by implementing this interface.
// This allows you to attach additional context or state to your errors.

type HTTPError struct {
	StatusCode int
	Path       string
	Message    string
}

// Implement the error interface
func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d on %s: %s", e.StatusCode, e.Path, e.Message)
}

// ─────── 7. ERROR WRAPPING (Go 1.13+) ───────
// You often want to add context to an error while preserving the original error
// so it can be inspected later. We use %w in fmt.Errorf to "wrap" the error.

func fetchResource(path string) error {
	// Simulate an error
	baseErr := errors.New("connection reset by peer")
	
	// Wrap the error with context
	return fmt.Errorf("failed to fetch %s: %w", path, baseErr)
}

func demonstrateCustomErrors() {
	fmt.Println("\n--- Custom Errors & Wrapping ---")
	
	// Using a custom error type
	err := HTTPError{
		StatusCode: 404,
		Path:       "/api/users/1",
		Message:    "user not found",
	}
	
	fmt.Println("Custom error string:", err.Error())
	
	// Let's look at error wrapping
	wrappedErr := fetchResource("/data")
	fmt.Println("Wrapped error string:", wrappedErr.Error())
	
	// We can use errors.Is to check if a specific error is anywhere in the chain
	baseErr := errors.New("connection reset by peer")
	// Note: Because errors.New creates a NEW error pointer every time, this will
	// actually return false in this specific demo unless we exported baseErr as a package variable.
	// We will see proper usage of errors.Is and errors.As in Module 14.
	fmt.Printf("Is it a connection reset? %v (Requires exact pointer match)\n", errors.Is(wrappedErr, baseErr))
	
	// Better to use errors.As to extract our custom error type from a wrapped chain
	var httpErr HTTPError
	
	// Let's simulate a wrapped HTTP error
	complexErr := fmt.Errorf("database query failed: %w", err)
	
	if errors.As(complexErr, &httpErr) {
		fmt.Printf("Extracted HTTP Error - Status: %d\n", httpErr.StatusCode)
	} else {
		fmt.Println("Not an HTTP error")
	}
}
