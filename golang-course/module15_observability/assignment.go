package main

import (
	"log/slog"
)

// TODO 1: Implement CreateLogger
// It should return a new *slog.Logger that outputs JSON.
// The output destination should be os.Stdout.
// The default log level should be set to Debug.
func CreateLogger() *slog.Logger {
	return nil // Fix me
}

// TODO 2: Implement LogTransaction
// It takes a logger, an account ID, an amount, and a boolean indicating success.
// It should log an Info message: "transaction processed"
// It must include three attributes:
// 1. "account_id" (int)
// 2. "amount" (float64)
// 3. "success" (bool)
func LogTransaction(logger *slog.Logger, accountID int, amount float64, success bool) {
	// Fix me
}
