package main

import (
	"fmt"
	"iter"
)

// ─────── 3. ITERATORS (Go 1.23+) ───────
// In Go 1.23, the `iter` package and "range over func" were introduced.
// This allows you to write custom iterators that can be used directly in a standard
// `for ... range` loop!
//
// Before 1.23, if you built a custom Tree data structure, you had to either:
// 1. Return a slice of all items (high memory usage)
// 2. Use a channel (slow and requires goroutine management)
// 3. Write a clunky Next() method.
//
// Now, you can return a function of type `iter.Seq` or `iter.Seq2`.

// Backward returns an iterator that yields elements of a slice in reverse order.
// iter.Seq2[int, E] means it yields TWO values: an int (index) and an E (the element).
func Backward[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			// The `yield` function is what the `for` loop body actually runs!
			// If `yield` returns false (e.g., if the user used `break` in their loop),
			// we must stop yielding values.
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// Fibonacci returns an iterator that yields the Fibonacci sequence infinitely.
// It yields just one value, so it uses iter.Seq[int].
func Fibonacci() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func demonstrateIterators() {
	fmt.Println("\n--- Iterators (Go 1.23+) ---")
	
	names := []string{"Alice", "Bob", "Charlie"}
	
	fmt.Println("Backward:")
	// WE ARE RANGING OVER A FUNCTION!
	for i, name := range Backward(names) {
		fmt.Printf("[%d] %s\n", i, name)
	}
	
	fmt.Println("\nFibonacci (first 10):")
	count := 0
	for val := range Fibonacci() {
		fmt.Print(val, " ")
		count++
		if count >= 10 {
			break // This causes the `yield` inside the iterator to return false!
		}
	}
	fmt.Println()
}
