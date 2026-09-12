package main

import "fmt"

// ─────── 1. GENERICS (Go 1.18+) ───────
// For over a decade, Go developers begged for Generics (Type Parameters).
// Before 1.18, if you wanted a function to reverse a slice, you had to write one
// for []int, one for []string, or use the slow and unsafe `reflect` package with `any`.

// PrintSlice uses a Type Parameter `[T any]`
// It means: "This function works for a slice of ANY type, and we'll call that type T".
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// ─────── 2. TYPE CONSTRAINTS ───────
// Sometimes `any` is too broad. What if we want to add two things together?
// We can't use `any` because we can't use the `+` operator on structs or booleans.
// We must CONSTRAIN the type.

// We can define custom constraints using interfaces.
type Number interface {
	int | int32 | int64 | float32 | float64
}

// Sum works for any type T that satisfies the Number constraint.
func Sum[T Number](numbers []T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

// Go provides standard constraints in the `cmp` package (like cmp.Ordered)
// for types that support <, >, <=, >=.

func demonstrateGenerics() {
	fmt.Println("\n--- Generics ---")
	
	ints := []int{1, 2, 3}
	strs := []string{"a", "b", "c"}
	
	// The compiler INFERS the type T based on the argument we pass!
	PrintSlice(ints) // T is int
	PrintSlice(strs) // T is string
	
	fmt.Println("Sum of ints:", Sum(ints))
	
	floats := []float64{1.5, 2.5}
	fmt.Println("Sum of floats:", Sum(floats))
	
	// Sum(strs) // COMPILER ERROR: string does not satisfy Number
}
