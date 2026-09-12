package main

import (
	"fmt"
)

// ─────── 9. HIGHER-ORDER FUNCTIONS ───────
// A higher-order function is a function that does at least one of the following:
// - takes one or more functions as arguments
// - returns a function as its result

// Map applies a function to every element of a slice and returns a new slice.
// Note: In Go 1.18+, this would typically be written using generics (Module 16).
// For now, we use a specific type (int to int).
func MapInts(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate.
func FilterInts(input []int, predicate func(int) bool) []int {
	var result []int
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func demonstrateHigherOrder() {
	fmt.Println("\n--- Higher-Order Functions ---")
	
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Example of passing an anonymous function as an argument to Map
	doubled := MapInts(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("Doubled: %v\n", doubled)

	// Example of passing an anonymous function to Filter
	evens := FilterInts(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Evens: %v\n", evens)
}
