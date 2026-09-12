package main

import (
	"fmt"
	"time"
)

// ─────── 1. GOROUTINES ───────
// "Concurrency is about dealing with lots of things at once.
//  Parallelism is about doing lots of things at once." - Rob Pike
//
// A goroutine is a lightweight thread managed by the Go runtime.
// They start with a tiny stack (typically 2KB) that grows and shrinks as needed.
// You can comfortably run hundreds of thousands of goroutines on a standard laptop.
//
// To start a goroutine, simply prefix a function call with the keyword `go`.

func printNumbers(prefix string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		// Yielding control back to the scheduler
		time.Sleep(10 * time.Millisecond)
	}
}

func demonstrateGoroutines() {
	fmt.Println("\n--- Goroutines ---")
	
	// This runs synchronously on the main goroutine
	printNumbers("Sync")
	
	// These run concurrently in the background
	go printNumbers("Async A")
	go printNumbers("Async B")
	
	// If the main goroutine exits, all other goroutines are abruptly terminated!
	// For this demo, we sleep to give them time to finish.
	// (In real code, we use sync.WaitGroup or Channels to wait for them).
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Main goroutine exiting.")
}
