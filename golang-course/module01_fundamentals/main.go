// Package main demonstrates the fundamentals of Go.
//
// ─────── 1. ORIGIN STORY & PHILOSOPHY ───────
// Go was conceived in late 2007 at Google by Robert Griesemer, Rob Pike, and Ken Thompson.
// The motivation? Frustration.
// - C++ took too long to compile.
// - Java was too verbose and heavy.
// - Python was too slow and lacked type safety.
//
// The vision was to create a language with the efficiency of a statically typed, compiled
// language, but with the ease of programming of an interpreted language.
//
// Guiding Philosophies:
// 1. Simplicity over complexity: "Less is exponentially more." (Rob Pike)
// 2. Concurrency is not parallelism: Go was designed for multi-core machines from day one.
// 3. Composition over inheritance: No class hierarchies.
// 4. Errors are values: Explicit error handling over implicit exceptions.
package main

import (
	"fmt"
	"os"
)

// ─────── 2. VARIABLES & ZERO VALUES ───────

// Package-level variables can be declared using the 'var' keyword.
// If no initial value is provided, Go assigns the "zero value" for the type.
var (
	globalCounter int    // zero value is 0
	globalName    string // zero value is ""
	isActive      bool   // zero value is false
)

func variablesDemo() {
	fmt.Println("--- Variables & Zero Values ---")
	// The short variable declaration ':=' can only be used inside functions.
	// It infers the type based on the right-hand side.
	localCount := 10
	name := "Gopher"

	// You can declare multiple variables at once.
	x, y, z := 1, 2.5, "three"
	fmt.Printf("x: %T, y: %T, z: %T\n", x, y, z)

	// The blank identifier '_' is used to discard values you don't need.
	_, b := 10, 20
	fmt.Printf("localCount: %d, name: %s, b: %d\n", localCount, name, b)
}

// ─────── 3. CONSTANTS & IOTA ───────

// Constants are evaluated at compile time. They cannot be changed.
const Pi = 3.14159265359

// `iota` is a magical identifier used in const declarations.
// It simplifies definitions of incrementing numbers (enums).
const (
	ReadPermission  = 1 << iota // 1 (1 << 0)
	WritePermission             // 2 (1 << 1)
	ExecPermission              // 4 (1 << 2)
)

func constantsDemo() {
	fmt.Println("\n--- Constants & Iota ---")
	fmt.Printf("Pi: %f\n", Pi)
	fmt.Printf("Read: %d, Write: %d, Exec: %d\n", ReadPermission, WritePermission, ExecPermission)
}

// ─────── 4. CONTROL STRUCTURES ───────

func controlStructuresDemo() {
	fmt.Println("\n--- Control Structures ---")

	// if statement with an initialization statement.
	// `v` is only in scope for the if/else blocks.
	if v := 10; v > 5 {
		fmt.Printf("%d is greater than 5\n", v)
	}

	// The 'for' loop is Go's only looping construct.
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Printf("Sum of 0..4 is %d\n", sum)

	// While-style 'for' loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Printf("n is now %d\n", n)

	// Switch statement (no 'break' needed, it doesn't fall through by default)
	os := "linux"
	switch os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Other OS")
	}
}

// ─────── 5. POINTERS & MEMORY (STACK vs HEAP) ───────

// Pointers hold the memory address of a value.
// Go has pointers but no pointer arithmetic (by default, for safety).
func pointersDemo() {
	fmt.Println("\n--- Pointers ---")

	i := 42
	p := &i // p is a pointer to i (type *int)

	fmt.Printf("Value of i: %d, Address of i: %p\n", i, p)

	*p = 21 // Dereferencing p to mutate the underlying value
	fmt.Printf("New value of i: %d\n", i)

	// 'new' vs 'make'
	// 'new(T)' allocates memory for a zeroed T and returns a pointer to it (*T).
	// It's rarely used because 'var v T' or 'v := &T{}' are preferred.
	ptr := new(int)
	fmt.Printf("Value at ptr: %d\n", *ptr)

	// Escape Analysis:
	// Go decides whether to put variables on the Stack or the Heap.
	// If a variable's reference "escapes" the function (e.g., returned as a pointer),
	// the compiler allocates it on the Heap so it survives the function return.
	// You can see this with: `go build -gcflags="-m"`
}

// ─────── 6. ERRORS ARE VALUES ───────

// In Go, errors are just values that implement the `error` interface.
// We do not throw exceptions. We return errors and handle them explicitly.
// This forces the programmer to think about failure states.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func errorsDemo() {
	fmt.Println("\n--- Errors as Values ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Expected to print the error
	}
}

func main() {
	fmt.Println("Welcome to Go Fundamentals!")
	variablesDemo()
	constantsDemo()
	controlStructuresDemo()
	pointersDemo()
	errorsDemo()

	// Interacting with OS args (Command Line)
	if len(os.Args) > 1 {
		fmt.Println("\nPassed arguments:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("Arg %d: %s\n", i+1, arg)
		}
	}
}
