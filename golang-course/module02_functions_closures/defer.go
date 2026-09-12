package main

import (
	"fmt"
)

// ─────── 6. DEFER ───────
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not
// executed until the surrounding function returns.
// Deferred function calls are pushed onto a stack. When a function returns, its
// deferred calls are executed in Last-In-First-Out (LIFO) order.

func demonstrateDefer() {
	fmt.Println("\n--- Defer ---")
	fmt.Println("Counting down:")

	// These will print 3, 2, 1 because of LIFO order.
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("Lift off!")
}

// ─────── 7. DEFER & NAMED RETURNS ───────
// Deferred functions can read and modify named return values.
// This is the most legitimate use case for named return values.

func addAndDouble(a, b int) (result int) {
	defer func() {
		result = result * 2 // Modifies the return value *after* the function body completes
	}()
	result = a + b
	return // Returns 5 (if a=2, b=3), but then the defer doubles it to 10.
}

// ─────── 8. PANIC AND RECOVER ───────
// Go does not have exceptions. It has `panic` and `recover`.
// panic: typically means something went unexpectedly wrong. It stops normal execution,
// runs all deferred functions, and crashes the program.
// recover: regains control of a panicking goroutine. It is only useful inside deferred functions.
//
// The Rule of Thumb:
// Only panic for programmer errors (e.g., index out of bounds, nil pointer).
// Return an `error` for expected failure cases (e.g., file not found, bad input).

func safeOperation() {
	// The recover idiom
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from a panic! The panic was: %v\n", r)
		}
	}()

	fmt.Println("Doing something safe...")
	doSomethingDangerous()
	fmt.Println("This line will never execute because of the panic.")
}

func doSomethingDangerous() {
	panic("I am panicking!") // Initiates the panic
}

func demonstratePanicRecover() {
	fmt.Println("\n--- Panic & Recover ---")
	safeOperation()
	fmt.Println("Program execution continued normally after recovery.")
}
