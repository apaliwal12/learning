package main

import (
	"fmt"
	"maps"
)

// ─────── 6. MAPS & INTERNALS ───────
// Maps in Go are hash tables. They provide unordered O(1) average time complexity for lookups, inserts, and deletes.
// A map is a pointer to a `runtime.hmap` struct internally.
// 
// When you pass a map to a function, you are passing a pointer to the `hmap`.
// Thus, modifications made inside the function are visible to the caller.
// 
// Map Keys: Must be a comparable type (booleans, numeric, string, pointer, channel, interface, structs/arrays containing only comparable types).
// Slices, maps, and functions cannot be used as map keys.

func demonstrateMaps() {
	fmt.Println("\n--- Maps ---")

	// Creating a map: map[KeyType]ValueType
	// Using make allows you to pre-allocate capacity if you know the rough size, which avoids rehashing.
	m := make(map[string]int, 10) 
	
	m["apple"] = 5
	m["banana"] = 10

	// Literal syntax
	scores := map[string]int{
		"Alice": 95,
		"Bob":   85,
	}
	fmt.Printf("Scores: %v\n", scores)

	// Reading from a map
	// If the key doesn't exist, it returns the zero value for the value type.
	fmt.Printf("Alice's score: %d\n", scores["Alice"])
	fmt.Printf("Charlie's score: %d\n", scores["Charlie"]) // Prints 0

	// The "comma ok" idiom to check if a key actually exists
	charlieScore, ok := scores["Charlie"]
	if ok {
		fmt.Printf("Charlie exists with score: %d\n", charlieScore)
	} else {
		fmt.Println("Charlie not found in map")
	}

	// Deleting a key
	delete(scores, "Bob")
	fmt.Printf("Scores after deleting Bob: %v\n", scores)

	// Iterating over a map (Order is randomized by design!)
	fmt.Println("Iterating scores:")
	for key, val := range scores {
		fmt.Printf(" - %s: %d\n", key, val)
	}
}

// ─────── 7. MAPS PACKAGE (Go 1.21+) ───────

func demonstrateMapsPkg() {
	fmt.Println("\n--- Maps Package ---")
	
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	m3 := map[string]int{"a": 1, "b": 3}

	fmt.Printf("m1 == m2? %v\n", maps.Equal(m1, m2))
	fmt.Printf("m1 == m3? %v\n", maps.Equal(m1, m3))

	// Clone a map
	m4 := maps.Clone(m1)
	m4["c"] = 3
	fmt.Printf("Cloned & modified: m1=%v, m4=%v\n", m1, m4)

	// Copy from src to dst (overwrites existing keys)
	maps.Copy(m1, map[string]int{"b": 99, "d": 4})
	fmt.Printf("After maps.Copy into m1: %v\n", m1)
}

// ─────── 8. MAP GOTCHAS & CONCURRENCY ───────

func demonstrateMapGotchas() {
	fmt.Println("\n--- Map Gotchas ---")

	// 1. Uninitialized Maps (nil map panic)
	var m map[string]int // m is nil
	// m["key"] = 10 // PANIC: assignment to entry in nil map
	
	// You can READ from a nil map (returns zero values), but you cannot WRITE to it.
	fmt.Printf("Reading from nil map: %d\n", m["key"]) 

	// 2. Concurrency
	// Maps are NOT safe for concurrent use. 
	// If one goroutine writes to a map while another is reading/writing, 
	// Go will crash the program with a "fatal error: concurrent map read and map write".
	// (We will cover sync.RWMutex and sync.Map in the Concurrency modules).

	// 3. Addressability
	// You cannot take the address of a map element.
	type User struct{ Name string }
	users := map[int]User{1: {"Alice"}}
	
	// _ = &users[1] // COMPILER ERROR: cannot take the address of users[1]
	// users[1].Name = "Bob" // COMPILER ERROR: cannot assign to struct field users[1].Name in map

	// Fix: Replace the whole struct or use pointers as values
	u := users[1]
	u.Name = "Bob"
	users[1] = u // Replace

	usersPtr := map[int]*User{1: {"Alice"}}
	usersPtr[1].Name = "Bob" // This works!
}
