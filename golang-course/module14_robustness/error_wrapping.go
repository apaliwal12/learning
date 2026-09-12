package main

import (
	"errors"
	"fmt"
)

// ─────── 1. ERROR WRAPPING ───────
// In Go 1.13, error wrapping was introduced. It allows you to add context
// to an error without losing the original underlying error.
// You wrap an error using fmt.Errorf with the `%w` verb.

var ErrDatabaseConnection = errors.New("database connection failed")

func connectToDB() error {
	// Simulating a low-level failure
	return ErrDatabaseConnection
}

func fetchUsers() error {
	err := connectToDB()
	if err != nil {
		// We wrap the low-level error with higher-level context
		return fmt.Errorf("fetchUsers: failed to connect: %w", err)
	}
	return nil
}

func demonstrateErrorWrapping() {
	fmt.Println("\n--- Error Wrapping ---")
	
	err := fetchUsers()
	fmt.Println("Print Error:", err)
	
	// ─────── 2. ERRORS.IS ───────
	// errors.Is checks if ANY error in the wrapped chain matches a target error.
	if errors.Is(err, ErrDatabaseConnection) {
		fmt.Println("Result: The underlying cause WAS a database connection issue!")
	}
	
	// ─────── 3. ERRORS.AS ───────
	// errors.As checks if ANY error in the wrapped chain matches a specific TYPE,
	// and if so, it extracts it into the provided variable.
	// (We'll just demonstrate the concept here)
	
	// type CustomDBError struct { Code int }
	// var dbErr *CustomDBError
	// if errors.As(err, &dbErr) {
	//     fmt.Println("Custom error code:", dbErr.Code)
	// }
}
