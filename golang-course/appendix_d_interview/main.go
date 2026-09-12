package main

import "fmt"

// ─────── APPENDIX D: INTERVIEW PREP ───────
// A rapid-fire Q&A of the most common Go interview questions.

func demonstrateInterviewQuestions() {
	fmt.Println("\n--- Go Interview Q&A ---")
	
	fmt.Println("\nQ1: What is the difference between an Array and a Slice?")
	fmt.Println("A: An array has a FIXED size determined at compile time (e.g., [3]int). " +
		"A slice is a dynamically-sized view into an underlying array. " +
		"A slice header contains 3 things: a pointer to the array, a length, and a capacity.")
	
	fmt.Println("\nQ2: How does `defer` work?")
	fmt.Println("A: `defer` pushes a function call onto a stack. When the surrounding function " +
		"returns, the deferred calls are executed in Last-In-First-Out (LIFO) order. " +
		"It is most commonly used for closing files, unlocking mutexes, and recovering from panics.")
	
	fmt.Println("\nQ3: What happens if you read from a closed channel? What if you write?")
	fmt.Println("A: Reading from a closed channel succeeds immediately, returning the zero value " +
		"of the channel's type and a boolean `false`. " +
		"Writing to a closed channel causes a PANIC.")
	
	fmt.Println("\nQ4: What is an empty interface (`interface{}` or `any`)?")
	fmt.Println("A: It is an interface with zero methods. Since every type implements zero methods, " +
		"an empty interface can hold a value of ANY type. It's often used when the type is " +
		"unknown at compile time (e.g., fmt.Print).")
	
	fmt.Println("\nQ5: How do you check the actual type of an interface variable at runtime?")
	fmt.Println("A: Using a Type Assertion: `val, ok := myVar.(string)`. " +
		"Or using a Type Switch: `switch v := myVar.(type) { ... }`.")
	
	fmt.Println("\nQ6: Why is Go fast?")
	fmt.Println("A: 1. Compiles directly to machine code (no VM). " +
		"2. Very efficient garbage collector tuned for low latency. " +
		"3. Goroutines are lightweight (starting at ~2KB) compared to OS threads (~1MB). " +
		"4. Excellent support for value types (stack allocation) which reduces GC pressure.")
}

func main() {
	demonstrateInterviewQuestions()
}
