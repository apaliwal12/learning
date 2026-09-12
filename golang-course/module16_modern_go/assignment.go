package main

import (
	"iter"
)

// TODO 1: Implement Filter
// It should take a slice of ANY type T, and a predicate function `keep func(T) bool`.
// It should return a NEW slice containing only the elements for which `keep` returns true.
// func Filter[T any](s []T, keep func(T) bool) []T { ... }

// TODO 2: Implement Map
// It should take a slice of type T1, and a transform function `fn func(T1) T2`.
// It should return a NEW slice of type T2.
// func Map[T1, T2 any](s []T1, transform func(T1) T2) []T2 { ... }

// TODO 3: Implement Range Iterator
// It should return an iter.Seq[int] that yields integers from `start` up to (but not including) `end`.
func Range(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// Fix me
	}
}
