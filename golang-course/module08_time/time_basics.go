package main

import (
	"fmt"
	"time"
)

// ─────── 1. TIME BASICS & FORMATTING ───────
// Time in Go is represented by the `time.Time` struct.
//
// FORMATTING QUIRK: Go does NOT use strftime (like %Y-%m-%d).
// Instead, Go uses a specific reference date:
// Mon Jan 2 15:04:05 MST 2006
// (Mnemonic: 01/02 03:04:05PM '06 -0700)
// To format a time, you literally show how the reference date would look!

func demonstrateTimeBasics() {
	fmt.Println("\n--- Time Basics & Formatting ---")

	now := time.Now()
	
	// Pre-defined formats
	fmt.Println("RFC3339:", now.Format(time.RFC3339))
	
	// Custom formats
	// YYYY-MM-DD
	fmt.Println("Custom Date:", now.Format("2006-01-02"))
	
	// 12-hour clock with PM
	fmt.Println("Custom Time:", now.Format("03:04 PM"))

	// ─────── 2. TIMEZONES & PARSING ───────
	// Never assume local time in distributed systems. Use UTC!
	utcTime := now.UTC()
	fmt.Println("UTC Time:", utcTime.Format(time.RFC3339))

	// Parsing strings into time.Time
	dateStr := "2024-10-31"
	// Parse requires the layout and the string
	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Error parsing:", err)
	} else {
		fmt.Printf("Parsed %s: %v\n", dateStr, parsed)
	}
	
	// Gotcha: time.Parse uses UTC by default if no timezone is in the string.
	// If you want to parse it into a specific timezone, use time.ParseInLocation.
}
