package main

import "fmt"

func main() {
	fmt.Println("Welcome to Module 10: Advanced Concurrency & Context")

	demonstrateWaitGroup()
	demonstrateMutex()
	demonstrateOnce()
	
	demonstrateContextCancel()
	demonstrateContextTimeout()
	demonstrateContextValues()
}
