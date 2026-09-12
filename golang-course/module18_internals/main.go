package main

import "fmt"

func main() {
	fmt.Println("Welcome to Module 18: Advanced Internals")

	stackAllocation()
	_ = heapAllocation()
	
	demonstrateMemory()
	demonstrateReflection()
}
