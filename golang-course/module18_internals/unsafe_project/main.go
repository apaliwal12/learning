package main

import (
	"fmt"
	"unsafe"
)

// ─────── UNSAFE PACKAGE PROJECT ───────
// Go's type system guarantees memory safety. You cannot treat an int like a string,
// or read memory outside the bounds of a slice.
// However, occasionally, for extreme performance optimization or C-interop,
// you need to break the rules. The `unsafe` package lets you do this.
// WARNING: USING UNSAFE MEANS YOU GIVE UP GO'S GUARANTEES. YOU CAN CAUSE SEGFAULTS!

func main() {
	fmt.Println("--- Unsafe Package Demo ---")
	
	// 1. Sizeof
	// Returns the size of a variable in bytes.
	var x int64 = 10
	fmt.Printf("Size of int64: %d bytes\n", unsafe.Sizeof(x))
	
	// 2. Unsafe Pointers and Casting
	// In safe Go, you cannot cast a *int to a *float64.
	// `unsafe.Pointer` is equivalent to `void*` in C.
	// You can cast ANY pointer to unsafe.Pointer, and cast unsafe.Pointer to ANY pointer!
	
	var i int64 = 42
	
	// Safe: Get pointer to i
	ptrToI := &i 
	
	// Unsafe: Cast to unsafe.Pointer
	unsafePtr := unsafe.Pointer(ptrToI)
	
	// Unsafe: Cast unsafe.Pointer to *float64
	ptrToFloat := (*float64)(unsafePtr)
	
	// Read it! It will be garbage because the bit pattern of int(42) is not a valid float.
	fmt.Printf("Int value: %d\n", i)
	fmt.Printf("Float interpretation of those bytes: %f\n", *ptrToFloat)
	
	// 3. String to Byte Slice (Zero Allocation!)
	// Normally, `[]byte("string")` allocates new memory and copies the data.
	// We can use unsafe to force a slice to point to the string's read-only memory.
	// (Note: Modifying this byte slice will cause a panic!)
	
	str := "Hello, Unsafe World!"
	
	// We get the pointer to the underlying string data
	strDataPtr := unsafe.StringData(str)
	
	// We construct a byte slice pointing to that data
	byteSlice := unsafe.Slice(strDataPtr, len(str))
	
	fmt.Printf("Zero-allocation byte slice: %s\n", string(byteSlice))
}
