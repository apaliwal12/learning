package main

import (
	"fmt"
)

// ─────── 4. CLOSURES & ANONYMOUS FUNCTIONS ───────
// A closure is an anonymous function that references variables from outside its body.
// The function may access and assign to the referenced variables; in this sense
// the function is "bound" to the variables.

// CounterFactory returns a function that increments and returns a counter.
// The returned function acts as a closure over the `count` variable.
// Because the closure captures the reference to `count`, the variable escapes
// to the heap and its lifetime is extended to match the closure's lifetime.
func CounterFactory() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func demonstrateClosures() {
	fmt.Println("\n--- Closures ---")
	counterA := CounterFactory()
	counterB := CounterFactory() // Completely separate state

	fmt.Printf("Counter A: %d\n", counterA()) // 1
	fmt.Printf("Counter A: %d\n", counterA()) // 2
	
	fmt.Printf("Counter B: %d\n", counterB()) // 1 (independent state)

	// Anonymous function executed immediately (IIFE: Immediately Invoked Function Expression)
	func(msg string) {
		fmt.Printf("IIFE says: %s\n", msg)
	}("Hello from an anonymous function!")

	// ─────── 5. THE LOOP-VARIABLE TRAP ───────
	// HISTORICAL NOTE (Pre-Go 1.22):
	// When creating closures inside a loop, capturing the loop variable used to be
	// a massive source of bugs. Prior to 1.22, the loop variable was reused for every iteration.
	// Therefore, closures would capture the *same memory address* and all closures
	// would see the *last* value of the loop variable.
	//
	// Go 1.22 FIX:
	// Go 1.22 changes this behavior so each iteration of a "for" loop creates a
	// *new* variable. The code below now works intuitively as expected.
	
	fmt.Println("\nLoop-Variable Capture:")
	var funcs []func()
	for i := 0; i < 3; i++ {
		// In Go < 1.22, you had to do: i := i 
		// Now it just works!
		funcs = append(funcs, func() {
			fmt.Printf("Value captured: %d\n", i)
		})
	}

	for _, f := range funcs {
		f() // Will print 0, 1, 2 in Go 1.22+ (would print 3, 3, 3 in Go <= 1.21)
	}
}
