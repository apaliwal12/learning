package main

import (
	"fmt"
	"time"
)

// ─────── 3. MONOTONIC CLOCKS VS WALL CLOCKS ───────
// OS clocks can jump backwards (e.g., NTP sync, Daylight Saving Time).
// If you measure duration using the "wall clock" and it jumps backwards,
// you might get a negative duration!
//
// To solve this, Go 1.9+ transparently includes a "Monotonic Clock" reading
// inside `time.Time` values returned by `time.Now()`.
// Monotonic clocks ALWAYS move forward and are not affected by OS time changes.
//
// RULE: Use `time.Sub()` or `time.Since()` to measure durations. They use the
// monotonic clock automatically if it is present.

func demonstrateMonotonic() {
	fmt.Println("\n--- Monotonic Clocks ---")
	
	start := time.Now() // Contains both wall clock and monotonic clock
	
	// Simulate work
	time.Sleep(50 * time.Millisecond)
	
	// time.Since uses the monotonic reading to calculate duration perfectly,
	// even if the OS time was changed during the Sleep.
	elapsed := time.Since(start)
	fmt.Printf("Work took: %v\n", elapsed)
	
	// Stripping the monotonic clock:
	// If you serialize a time (to JSON, DB, etc.), the monotonic reading is lost!
	// Only the wall clock is saved.
	strippedStart := start.Round(0)
	
	// Comparing times
	fmt.Printf("start == strippedStart? %v\n", start == strippedStart) // false, because one has monotonic!
	
	// Correct way to compare if they represent the same instant:
	fmt.Printf("start.Equal(strippedStart)? %v\n", start.Equal(strippedStart)) // true
}

// ─────── 4. DURATIONS ───────
// time.Duration is an int64 representing nanoseconds.
// ALWAYS multiply by the predefined constants (time.Second, time.Millisecond).

func demonstrateDurations() {
	fmt.Println("\n--- Durations ---")
	
	d := 5 * time.Second
	fmt.Printf("Duration: %v (as float seconds: %f)\n", d, d.Seconds())
	
	// Adding duration to time
	now := time.Now()
	future := now.Add(2 * time.Hour)
	fmt.Printf("2 hours from now: %s\n", future.Format("15:04:05"))
}
