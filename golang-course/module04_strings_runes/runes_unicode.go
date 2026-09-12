package main

import (
	"fmt"
	"unicode/utf8"
)

// ─────── 3. RUNES & UTF-8 ───────
// Go's source code is defined to be UTF-8 text; therefore, string literals are UTF-8 text.
// 
// A 'rune' is an alias for 'int32'. It represents a single Unicode code point.
// While a string is a sequence of bytes (uint8), when you iterate over a string
// using a `for range` loop, Go automatically decodes the UTF-8 bytes into runes!

func demonstrateRunes() {
	fmt.Println("\n--- Runes & UTF-8 ---")
	
	// 'e' with an acute accent (é) takes 2 bytes in UTF-8.
	// The world emoji (🌍) takes 4 bytes.
	s := "Café 🌍" 

	// len() returns the number of BYTES, not the number of characters!
	fmt.Printf("String: %s\n", s)
	fmt.Printf("len(s) in bytes: %d\n", len(s)) 
	
	// To get the actual number of characters (runes), use utf8.RuneCountInString
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s))

	// ─────── 4. ITERATING OVER STRINGS ───────
	
	fmt.Println("\nIterating by BYTE (using standard for loop):")
	for i := 0; i < len(s); i++ {
		// This will print the raw bytes. Multi-byte characters will be split!
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	fmt.Println("\nIterating by RUNE (using for-range loop):")
	for index, runeValue := range s {
		// Notice how the index skips! (0, 1, 2, 3, 5, 6, 7...)
		// That's because the 'é' starts at byte 3 and ends at byte 4.
		fmt.Printf("Byte %d: %c (type: %T, value: %U)\n", index, runeValue, runeValue, runeValue)
	}
	
	// Converting string to []rune
	// This allocates a new array and decodes the entire UTF-8 string into int32s.
	// Useful if you need to access characters by index (e.g., getting the 5th character).
	r := []rune(s)
	fmt.Printf("\nAs []rune: %v, length: %d\n", r, len(r))
}
