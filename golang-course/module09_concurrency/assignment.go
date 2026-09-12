package main

import (
	"time"
)

// TODO 1: Implement AsyncSquare
// It should take an integer and return a channel.
// It should launch a goroutine that squares the integer, sends the result
// to the channel, and then closes the channel.
func AsyncSquare(n int) <-chan int {
	return nil // Fix me
}

// TODO 2: Implement FanIn
// It should take two integer channels and return a single integer channel.
// It should launch a goroutine that reads from both input channels and
// forwards any received values to the output channel.
// Use a `select` statement inside a `for` loop.
// Note: For this basic assignment, you don't need to worry about closing the
// output channel when the inputs close.
func FanIn(ch1, ch2 <-chan int) <-chan int {
	return nil // Fix me
}

// TODO 3: Implement ExecuteWithTimeout
// It should take a function `work` and a `timeout` duration.
// It should run the work function in a goroutine.
// If the work completes before the timeout, return true.
// If the timeout expires before the work completes, return false.
// Hint: use a channel to signal completion of the work, and select with time.After.
func ExecuteWithTimeout(work func(), timeout time.Duration) bool {
	return false // Fix me
}
