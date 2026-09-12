package main

import (
	"fmt"
	"slices"
	"unsafe"
)

// ─────── 1. ARRAYS ───────
// Arrays in Go are fixed-size values. The size is part of the type (e.g., [4]int).
// Because they are values, assigning an array to a new variable copies all elements.
// Arrays are primarily used as backing storage for slices.

func demonstrateArrays() {
	fmt.Println("\n--- Arrays ---")
	var a [3]int
	a[0] = 10
	a[1] = 20
	a[2] = 30
	fmt.Printf("Array a: %v, type: %T, size: %d bytes\n", a, a, unsafe.Sizeof(a))

	// [...] infers the size from the number of elements
	b := [...]string{"Go", "Rust", "Zig"}
	fmt.Printf("Array b: %v, type: %T\n", b, b)

	// Copying arrays copies the values
	c := a
	c[0] = 999
	fmt.Printf("After modifying copy - original: %v, copy: %v\n", a, c)
}

// ─────── 2. SLICE HEADER & INTERNALS ───────
// Slices are dynamic, flexible views into the elements of an array.
// Internally, a slice is a 24-byte struct (on 64-bit systems) containing:
// { ptr *T, len int, cap int }
// - ptr: Pointer to the first element of the backing array.
// - len: Number of elements currently in the slice.
// - cap: Maximum capacity of the backing array from the ptr offset.

func demonstrateSlices() {
	fmt.Println("\n--- Slices & Internals ---")
	
	// Creating a slice with make(Type, len, cap)
	s := make([]int, 2, 5)
	s[0] = 1
	s[1] = 2
	fmt.Printf("s: %v, len: %d, cap: %d, size: %d bytes\n", s, len(s), cap(s), unsafe.Sizeof(s))

	// Creating a slice from an array
	arr := [5]int{10, 20, 30, 40, 50}
	s2 := arr[1:4] // From index 1 (inclusive) to 4 (exclusive)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	// Modifying the slice modifies the backing array
	s2[0] = 999
	fmt.Printf("Modified s2: %v, underlying arr: %v\n", s2, arr)
}

// ─────── 3. APPEND & GROWTH ───────
// The append function adds elements to the end of a slice.
// If the backing array has enough capacity, append modifies it in place and returns a new slice header with a larger len.
// If the backing array is full, append allocates a NEW, larger array, copies the elements, and returns the new slice.

func demonstrateAppend() {
	fmt.Println("\n--- Append & Growth ---")
	
	var s []int // nil slice, len=0, cap=0
	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Printf("Appended %d: len=%d, cap=%d, address=%p\n", i, len(s), cap(s), s)
	}
	// Notice how the capacity doubles when it needs to grow (1 -> 2 -> 4 -> 8).
	// Starting in Go 1.18, the growth factor for large slices smoothly transitions from 2x to 1.25x.
}

// ─────── 4. SLICE GOTCHAS & 3-INDEX SLICING ───────

func demonstrateSliceGotchas() {
	fmt.Println("\n--- Slice Gotchas ---")

	// The Memory Leak Trap:
	// Taking a small slice of a massive array keeps the ENTIRE massive array in memory.
	massive := make([]byte, 10*1024*1024) // 10MB
	tiny := massive[0:10]
	_ = tiny // massive cannot be garbage collected because tiny references its backing array.

	// The Fix: Use slices.Clone() or copy()
	tinyFixed := slices.Clone(massive[0:10])
	_ = tinyFixed // Now massive can be garbage collected (if no other references exist).

	// The Appending to Sub-slices Trap:
	a := []int{1, 2, 3, 4, 5}
	b := a[:3] // [1, 2, 3], cap=5
	
	// Because b has capacity, appending to it overwrites the elements in 'a'!
	b = append(b, 99)
	fmt.Printf("After appending to b - a: %v, b: %v\n", a, b)

	// The Fix: 3-Index Slicing [low:high:max]
	// This artificially restricts the capacity of the new slice to prevent overwriting.
	a2 := []int{1, 2, 3, 4, 5}
	b2 := a2[0:3:3] // [1, 2, 3], cap=3 (max-low)
	b2 = append(b2, 99) // Forces a reallocation because cap is reached
	fmt.Printf("After appending to b2 (3-index) - a2: %v, b2: %v\n", a2, b2)
}

// ─────── 5. SLICES PACKAGE (Go 1.21+) ───────

func demonstrateSlicesPkg() {
	fmt.Println("\n--- Slices Package ---")
	
	names := []string{"Bob", "Alice", "Eve", "Charlie"}
	
	slices.Sort(names)
	fmt.Printf("Sorted: %v\n", names)

	idx, found := slices.BinarySearch(names, "Charlie")
	fmt.Printf("Found Charlie? %v at index %d\n", found, idx)

	contains := slices.Contains(names, "Dave")
	fmt.Printf("Contains Dave? %v\n", contains)

	// Custom sorting
	numbers := []int{5, 2, 8, 1, 9}
	slices.SortFunc(numbers, func(a, b int) int {
		return b - a // Descending order
	})
	fmt.Printf("Sorted Descending: %v\n", numbers)
}
