package main

import (
	"fmt"
	"strings"
)

// ─────── 5. STRING BUILDER & EFFICIENT CONCATENATION ───────
// Because strings are immutable, concatenating strings using `+` inside a loop
// is highly inefficient. It allocates a new string and copies both strings into
// it on EVERY iteration. O(N^2) time complexity.
// 
// To build strings efficiently, always use `strings.Builder`.
// It maintains an internal `[]byte` buffer and only allocates a new string once
// at the very end when you call `String()`.

func demonstrateStringBuilder() {
	fmt.Println("\n--- strings.Builder ---")
	
	// Create a builder
	var builder strings.Builder
	
	// Optional but recommended: Grow the internal buffer if you know the rough size.
	// This prevents reallocation during appending.
	builder.Grow(100)
	
	// Append strings, bytes, or runes
	builder.WriteString("Building a string ")
	builder.WriteString("is fast ")
	builder.WriteByte('!')
	builder.WriteRune('🚀')
	
	// Generate the final string (Zero-allocation conversion!)
	// strings.Builder uses an unsafe pointer cast internally to convert its []byte
	// into a string without copying the bytes, because it knows the []byte won't
	// be mutated after String() is called.
	result := builder.String()
	
	fmt.Printf("Final string: %s\n", result)
}

// ─────── 6. USEFUL STRINGS PACKAGE FUNCTIONS ───────

func demonstrateStringsPackage() {
	fmt.Println("\n--- strings Package Utilities ---")
	
	s := "  Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.  "
	
	fmt.Printf("Original: %q\n", s)
	
	// Trimming
	trimmed := strings.TrimSpace(s)
	fmt.Printf("Trimmed: %q\n", trimmed)
	
	// Contains, Prefix, Suffix
	fmt.Printf("Contains 'reliable': %v\n", strings.Contains(trimmed, "reliable"))
	fmt.Printf("Starts with 'Go': %v\n", strings.HasPrefix(trimmed, "Go"))
	
	// Splitting and Joining
	words := strings.Split(trimmed, " ")
	fmt.Printf("Word count: %d\n", len(words))
	
	joined := strings.Join(words[:5], "-") // "Go-is-an-open-source"
	fmt.Printf("Joined first 5: %s\n", joined)
	
	// Replacements (n=-1 means replace all)
	replaced := strings.Replace(trimmed, "simple", "complex (just kidding, simple)", 1)
	fmt.Printf("Replaced: %s\n", replaced)
}
