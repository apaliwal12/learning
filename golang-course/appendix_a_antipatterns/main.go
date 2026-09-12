package main

import (
	"fmt"
	"sync"
)

// ─────── APPENDIX A: GO ANTI-PATTERNS ───────
// Knowing what NOT to do is just as important as knowing what to do.

// 1. Goroutine Leaks
// A goroutine leak happens when a goroutine is started but never finishes.
// It will sit in memory forever, slowly exhausting your server's RAM.
func antiPatternGoroutineLeak() {
	ch := make(chan int)

	go func() {
		// This goroutine is blocked forever waiting to send,
		// because nobody is reading from `ch`!
		ch <- 1
	}()

	// FIX: Use a context with a timeout, or a buffered channel,
	// or ensure there is always a reader!
}

// 2. Loop Variable Capture (Pre-Go 1.22)
// Before Go 1.22, the `val` variable in a `for` loop was reused for every iteration.
// If you spawned a goroutine inside the loop, it would likely print the LAST value
// multiple times!
func antiPatternLoopCapture() {
	fmt.Println("--- Loop Capture ---")
	words := []string{"apple", "banana", "cherry"}
	var wg sync.WaitGroup

	for _, word := range words {
		wg.Add(1)
		// ANTI-PATTERN in <1.22 (Though fixed in 1.22+, it's still good to know!)
		// The safe way historically was to pass it as an argument: func(w string)
		go func(w string) {
			defer wg.Done()
			fmt.Println("Processing:", w)
		}(word)
	}
	wg.Wait()
}

// 3. Shadowing Errors
// Using `:=` can accidentally declare a NEW `err` variable that hides the outer one!
func antiPatternShadowing() error {
	var finalErr error

	// Do something...
	if true {
		// DANGER: `:=` creates a NEW error variable scoped ONLY to this `if` block.
		// It does NOT update `finalErr`.
		// result, finalErr := someFunctionThatErrors()
		result, err := func() (int, error) { return 0, fmt.Errorf("hidden error") }()
		_ = result
		// FIX: Use `=` if you want to update the outer variable, or define `err` explicitly.
		finalErr = err
	}

	return finalErr
}

// 4. Naked returns with named return variables in long functions
// It makes the code incredibly hard to read.
func antiPatternNakedReturn() (result int, err error) {
	result = 42

	// ... 100 lines of code ...

	// What are we returning here?! You have to scroll up to find out.
	// FIX: Always be explicit: `return result, err`
	return
}

func main() {
	fmt.Println("Welcome to Appendix A: Anti-Patterns")
	antiPatternLoopCapture()

	if err := antiPatternShadowing(); err != nil {
		fmt.Println("Shadowing error:", err)
	}
}
