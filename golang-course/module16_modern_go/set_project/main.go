package main

import (
	"fmt"
	"iter"
)

// ─────── GENERIC SET PROJECT ───────
// Go does not have a built-in Set data structure.
// Historically, developers used `map[string]struct{}` for string sets, and
// `map[int]struct{}` for integer sets.
// With Generics, we can write a Set ONCE and use it for ANY type that is comparable!

// comparable is a built-in constraint that allows == and != operations.
// It is required for map keys.
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates a new initialized Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Add inserts an item into the set.
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Contains returns true if the item is in the set.
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Items returns an iterator over the items in the set.
// This combines Go 1.18 Generics with Go 1.23 Iterators!
func (s *Set[T]) Items() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s.items {
			if !yield(item) {
				return
			}
		}
	}
}

func main() {
	fmt.Println("--- Generic Set Demo ---")
	
	// String Set
	stringSet := NewSet[string]()
	stringSet.Add("apple")
	stringSet.Add("banana")
	stringSet.Add("apple") // Duplicate ignored!
	
	fmt.Printf("Contains banana? %v\n", stringSet.Contains("banana"))
	
	fmt.Println("String Set items:")
	for item := range stringSet.Items() {
		fmt.Printf("- %s\n", item)
	}
	
	// Integer Set
	intSet := NewSet[int]()
	intSet.Add(10)
	intSet.Add(20)
	intSet.Add(30)
	
	fmt.Println("\nInteger Set items:")
	for item := range intSet.Items() {
		fmt.Printf("- %d\n", item)
	}
}
