package main

import (
	"fmt"
	"strings"
	"time"
)

// Compose takes two functions and returns a new function that applies
// the second function to the result of the first function.
// f(g(x))
func Compose(f, g func(string) string) func(string) string {
	return func(s string) string {
		return f(g(s))
	}
}

// Pipeline takes a slice of functions and returns a single function that
// applies all functions in sequence.
func Pipeline(funcs ...func(string) string) func(string) string {
	return func(s string) string {
		result := s
		for _, fn := range funcs {
			result = fn(result)
		}
		return result
	}
}

// RetryWithBackoff creates a closure that retries an operation up to maxRetries times.
// If the operation fails, it waits for an exponentially increasing delay before retrying.
// (In a real app, you'd use a context here, but we haven't learned that yet!)
func RetryWithBackoff(maxRetries int, op func() error) error {
	delay := 10 * time.Millisecond // start with small delay for fast testing

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := op()
		if err == nil {
			return nil // Success!
		}
		
		fmt.Printf("Attempt %d failed: %v. Retrying in %v...\n", attempt, err, delay)
		
		if attempt < maxRetries {
			time.Sleep(delay)
			delay *= 2 // Exponential backoff
		} else {
			return fmt.Errorf("operation failed after %d attempts: %w", maxRetries, err)
		}
	}
	return nil
}

// ResourceTracker uses defer to track resource usage.
// It returns a function that must be called to "release" the resource.
func ResourceTracker(name string) func() {
	fmt.Printf("[Resource] Acquiring %s\n", name)
	
	// Return a cleanup function
	return func() {
		fmt.Printf("[Resource] Releasing %s\n", name)
	}
}

func main() {
	fmt.Println("--- Functional Pipeline Demo ---")

	// 1. String transformation pipeline
	trim := strings.TrimSpace
	upper := strings.ToUpper
	exclaim := func(s string) string { return s + "!!!" }

	processString := Pipeline(trim, upper, exclaim)
	
	input := "   hello world   "
	fmt.Printf("Input:  %q\n", input)
	fmt.Printf("Output: %q\n", processString(input))

	// 2. Resource Management Demo
	fmt.Println("\n--- Resource Tracking Demo ---")
	func() {
		// Acquire resource and defer its release
		releaseDB := ResourceTracker("DatabaseConnection")
		defer releaseDB()

		releaseFile := ResourceTracker("TempFile")
		defer releaseFile()

		fmt.Println("Doing work with resources...")
		// When this function ends, the deferred releases will happen in reverse order (LIFO)
	}()

	// 3. Retry Demo
	fmt.Println("\n--- Retry Closure Demo ---")
	failures := 0
	flakyOperation := func() error {
		if failures < 2 {
			failures++
			return fmt.Errorf("network timeout")
		}
		fmt.Println("Operation succeeded!")
		return nil
	}

	err := RetryWithBackoff(3, flakyOperation)
	if err != nil {
		fmt.Printf("Final error: %v\n", err)
	}
}
