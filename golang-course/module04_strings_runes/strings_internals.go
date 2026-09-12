package main

import (
	"fmt"
	"unsafe"
)

// ─────── 1. STRING INTERNALS ───────
// In Go, a string is a read-only slice of bytes.
// Internally, a string is a 16-byte struct (on 64-bit systems) containing:
// { ptr *byte, len int }
// 
// Because strings are read-only, passing them around is extremely cheap.
// You are only ever passing the 16-byte header, never copying the underlying bytes.

func demonstrateStringInternals() {
	fmt.Println("\n--- String Internals ---")
	
	s1 := "hello"
	s2 := s1 // Copies the 16-byte header (ptr, len). The backing bytes are shared.
	
	fmt.Printf("s1 size: %d bytes, s2 size: %d bytes\n", unsafe.Sizeof(s1), unsafe.Sizeof(s2))
	
	// Since strings are read-only byte slices, we can index into them.
	// But be careful! Indexing a string returns the BYTE at that position, not the character.
	fmt.Printf("s1[0] (the byte): %v, as string: %q\n", s1[0], string(s1[0]))

	// s1[0] = 'H' // COMPILER ERROR: cannot assign to s1[0]
}

// ─────── 2. STRINGS AS BYTE SLICES ───────
// You can freely convert between `string` and `[]byte`.
// However, doing so allocates a NEW backing array and copies the data.
// Why? Because strings are immutable, but `[]byte` is mutable. If they shared memory,
// you could mutate the `[]byte` and accidentally change the immutable string!

func demonstrateStringByteConversion() {
	fmt.Println("\n--- String <-> []byte Conversion ---")
	
	s := "immutable"
	
	// Convert to []byte (ALLOCATION + COPY)
	b := []byte(s)
	b[0] = 'I' // We can mutate the byte slice
	
	// Convert back to string (ALLOCATION + COPY)
	s2 := string(b)
	
	fmt.Printf("Original: %s, Mutated: %s\n", s, s2)
	
	// Advanced Note: In highly performance-sensitive code, there are ways to do
	// zero-allocation conversions using the `unsafe` package, but it is dangerous 
	// and violates Go's memory safety guarantees if you mutate the bytes.
}
