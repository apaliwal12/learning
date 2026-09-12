package main

import (
	"fmt"
)

// ─────── 1. FUNCTION SIGNATURES & MULTIPLE RETURNS ───────
// Functions in Go are first-class citizens.
// Go supports multiple return values, which is heavily used for the (result, error) idiom.

// Divide returns both the quotient and the remainder.
func Divide(a, b int) (int, int) {
	return a / b, a % b
}

// Named return values can act as local variables.
// Use them sparingly. They are mainly useful for documentation or 
// when you need to capture return values in a `defer` block.
func DivideNamed(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	// Naked return: returns the current values of quotient and remainder.
	// Generally discouraged in long functions as it hurts readability.
	return
}

// ─────── 2. VARIADIC FUNCTIONS ───────
// A variadic function can accept a variable number of arguments of a specific type.
// Inside the function, the variadic parameter `nums` acts as a slice of integers (`[]int`).

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// ─────── 3. INIT FUNCTION ───────
// init() functions execute automatically before main() starts.
// They are used for package-level initialization state.
// You can have multiple init() functions in a single package or even in a single file.
// Order of execution: imported packages -> package level vars -> init() -> main()

func init() {
	// This will print before anything in main() runs.
	// fmt.Println("init() in functions.go executing...")
}

func demonstrateFunctions() {
	fmt.Println("\n--- Functions & Signatures ---")
	q, r := Divide(10, 3)
	fmt.Printf("10 / 3 = %d, remainder %d\n", q, r)

	q2, r2 := DivideNamed(14, 4)
	fmt.Printf("14 / 4 = %d, remainder %d (using named returns)\n", q2, r2)

	total := Sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum(1, 2, 3, 4, 5) = %d\n", total)

	// Spreading a slice into a variadic function:
	numbers := []int{10, 20, 30}
	total2 := Sum(numbers...) // The '...' unwraps the slice
	fmt.Printf("Sum(numbers...) = %d\n", total2)
}
