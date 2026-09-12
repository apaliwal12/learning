package main

import (
	"fmt"
	"time"
)

// ─────── 5. TIMERS & TICKERS ───────
// These are used for scheduling future events.
// They return a struct that contains a Channel (C). The Go runtime will send
// the current time on that channel when the timer expires.

func demonstrateTimers() {
	fmt.Println("\n--- Timers & Tickers ---")
	
	// time.After is a convenience wrapper around time.NewTimer.
	// It returns a channel that receives the time after the duration.
	// CAUTION: Using time.After in a loop can cause memory leaks if the
	// timer is not reached, because time.After cannot be stopped manually.
	fmt.Println("Waiting 100ms...")
	<-time.After(100 * time.Millisecond)
	fmt.Println("Done waiting!")
	
	// A Timer can be stopped or reset.
	// This is the preferred way if you might need to cancel it.
	timer := time.NewTimer(1 * time.Second)
	
	go func() {
		// Wait for timer to fire
		<-timer.C
		fmt.Println("Timer fired!")
	}()
	
	// Stop the timer before it fires
	if timer.Stop() {
		fmt.Println("Timer was successfully stopped before firing.")
	}

	// Tickers fire repeatedly at a given interval.
	ticker := time.NewTicker(50 * time.Millisecond)
	
	// It's crucial to stop tickers to release resources!
	defer ticker.Stop()
	
	count := 0
	for t := range ticker.C {
		fmt.Printf("Tick at %v\n", t.Format("15:04:05.000"))
		count++
		if count >= 3 {
			break
		}
	}
}
