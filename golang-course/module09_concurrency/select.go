package main

import (
	"fmt"
	"time"
)

// ─────── 5. THE SELECT STATEMENT ───────
// The `select` statement lets a goroutine wait on multiple communication operations.
// A `select` blocks until one of its cases can run, then it executes that case.
// If multiple are ready, it chooses one at random.
//
// This is the bedrock of advanced concurrent patterns in Go (timeouts, cancellations).

func demonstrateSelect() {
	fmt.Println("\n--- Select Statement ---")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch1 <- "Data from ch1"
	}()
	
	go func() {
		time.Sleep(10 * time.Millisecond) // This one will finish first
		ch2 <- "Data from ch2"
	}()
	
	// We wait on both channels simultaneously
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
	
	// ─────── 6. SELECT TIMEOUTS ───────
	// By combining `select` with `time.After`, we can create highly robust
	// systems that never block forever.
	
	fmt.Println("\n--- Select Timeout ---")
	slowCh := make(chan string)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		slowCh <- "This is too slow"
	}()
	
	select {
	case msg := <-slowCh:
		fmt.Println("Received:", msg)
	case <-time.After(30 * time.Millisecond): // Timeout fires first!
		fmt.Println("Operation timed out!")
	}
	
	// ─────── 7. DEFAULT CASE (NON-BLOCKING CHANNELS) ───────
	// A `default` case in a `select` is run if no other case is ready.
	// This turns a blocking channel operation into a non-blocking one.
	
	fmt.Println("\n--- Default Case (Non-blocking) ---")
	emptyCh := make(chan int)
	
	select {
	case val := <-emptyCh:
		fmt.Println("Received:", val)
	default:
		fmt.Println("Nothing to receive right now, moving on...")
	}
}
