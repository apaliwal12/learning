package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// ─────── 9. STRUCTS & MEMORY LAYOUT ───────
// A struct is a typed collection of fields. They're useful for grouping data together.
// Memory Alignment: Go structures are padded to align fields in memory for faster CPU access.

type User struct {
	ID        int
	IsActive  bool
	// 7 bytes of padding inserted here on 64-bit systems!
	Username  string
}

// OptimizedUser reorders fields to minimize padding.
type OptimizedUser struct {
	ID       int    // 8 bytes
	Username string // 16 bytes
	IsActive bool   // 1 byte
	// 7 bytes of padding at the END (to align the total size to a multiple of 8)
}

func demonstrateStructs() {
	fmt.Println("\n--- Structs & Alignment ---")

	// Struct initialization
	u1 := User{
		ID:       1,
		IsActive: true,
		Username: "gopher", // Note the trailing comma. It's required!
	}
	fmt.Printf("User: %+v\n", u1) // %+v prints field names

	// Anonymous struct (useful for one-off data grouping, like test cases)
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Config: %v\n", config)

	// Memory Alignment Demo
	fmt.Printf("Size of User: %d bytes\n", unsafe.Sizeof(User{}))
	fmt.Printf("Size of OptimizedUser: %d bytes\n", unsafe.Sizeof(OptimizedUser{}))
}

// ─────── 10. METHODS & POINTER VS VALUE RECEIVERS ───────
// Go doesn't have classes, but you can define methods on types.
// A method is a function with a special "receiver" argument.

type Counter struct {
	Value int
}

// Value Receiver: Operates on a COPY of the struct.
// Does NOT modify the original struct.
func (c Counter) IncrementValue() {
	c.Value++
}

// Pointer Receiver: Operates on a POINTER to the struct.
// Modifies the original struct.
// Use pointer receivers when:
// 1. You need to modify the state of the receiver.
// 2. The struct is large, and copying it would be expensive.
// 3. For consistency (if some methods use pointers, use pointers for all).
func (c *Counter) IncrementPointer() {
	c.Value++
}

func demonstrateMethods() {
	fmt.Println("\n--- Pointer vs Value Receivers ---")
	
	c := Counter{Value: 0}
	
	c.IncrementValue()
	fmt.Printf("After IncrementValue: %d (Unchanged!)\n", c.Value)
	
	// Go automatically takes the address of 'c' here (&c).
	c.IncrementPointer()
	fmt.Printf("After IncrementPointer: %d (Changed!)\n", c.Value)
}

// ─────── 11. STRUCT TAGS ───────
// Struct tags provide metadata about fields.
// They are heavily used for JSON serialization, database ORMs (like GORM), and validation.

type APIResponse struct {
	// `json:"user_id"` tells the JSON encoder to use "user_id" as the key.
	UserID    int    `json:"user_id"`
	// `omitempty` tells the encoder to omit the field if it has a zero value.
	AuthToken string `json:"auth_token,omitempty"` 
	// `-` tells the encoder to ignore this field entirely.
	InternalData string `json:"-"`
}

func demonstrateTags() {
	fmt.Println("\n--- Struct Tags ---")
	
	resp := APIResponse{
		UserID:       42,
		AuthToken:    "", // Zero value, will be omitted
		InternalData: "secret",
	}

	bytes, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(bytes))
}
