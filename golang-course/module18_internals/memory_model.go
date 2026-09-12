package main

import (
	"fmt"
	"runtime"
)

// ─────── 1. THE GO MEMORY MODEL ───────
// Go is a garbage-collected language, but UNDERSTANDING how it manages memory
// separates average Go developers from masters.
//
// Stack vs Heap:
// - STACK: Fast, thread-specific memory. Used for variables that don't "escape"
//   the function they are declared in. Memory is cleaned up instantly when the func returns.
// - HEAP: Slower, global memory. Used for variables that are shared across boundaries,
//   or returned as pointers. Memory is cleaned up by the Garbage Collector (GC).

// This variable will likely stay on the STACK because it is not returned as a pointer
// and its lifecycle is contained entirely within this function or those it calls synchronously.
func stackAllocation() int {
	x := 42
	return x 
}

// This variable MUST ESCAPE to the HEAP.
// Why? Because we are returning a pointer to it. If it lived on the stack,
// the memory would be wiped when the function returned, and the pointer would be invalid!
// The Go compiler's "Escape Analysis" detects this automatically.
// You can view escape analysis by building with: `go build -gcflags="-m"`
func heapAllocation() *int {
	x := 42
	return &x
}

// ─────── 2. GARBAGE COLLECTION ───────
// Go uses a Concurrent Mark-and-Sweep Garbage Collector.
// It prioritizes LOW LATENCY (sub-millisecond pauses) over maximum throughput.
// You rarely need to tune it, but you can use the `GOGC` environment variable.
// GOGC=100 (default) means: "Run a GC cycle when the heap size doubles since the last collection."
// GOGC=off turns off the GC!

func demonstrateMemory() {
	fmt.Println("\n--- Memory Model & GC ---")
	
	// We can force a garbage collection cycle (rarely needed in production)
	runtime.GC()
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Allocated Heap: %d bytes\n", m.Alloc)
	fmt.Printf("Total Allocations (cumulative): %d bytes\n", m.TotalAlloc)
	fmt.Printf("Number of GC Cycles: %d\n", m.NumGC)
}
