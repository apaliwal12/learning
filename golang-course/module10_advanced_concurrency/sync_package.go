package main

import (
	"fmt"
	"sync"
	"time"
)

// ─────── 1. WAITGROUP ───────
// sync.WaitGroup is the standard way to wait for a collection of goroutines to finish.
// It is essentially a thread-safe counter.
// - Add(delta int): Increments the counter.
// - Done(): Decrements the counter (same as Add(-1)).
// - Wait(): Blocks until the counter becomes 0.

func demonstrateWaitGroup() {
	fmt.Println("\n--- sync.WaitGroup ---")
	
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Always Add *before* launching the goroutine!
		
		go func(id int) {
			defer wg.Done() // Ensure Done is called even if the function panics
			
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(20 * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
	
	fmt.Println("Main: waiting for workers to finish...")
	wg.Wait()
	fmt.Println("Main: all workers finished!")
}

// ─────── 2. MUTEX (MUTUAL EXCLUSION) ───────
// While channels are preferred for orchestrating data flow, sometimes you just
// need to protect a shared piece of state (like a map or a counter) from concurrent access.
// This is where sync.Mutex comes in.

type SafeCounter struct {
	mu sync.Mutex // Protects the value
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock() // Always defer Unlock to prevent deadlocks on panics
	
	c.v[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func demonstrateMutex() {
	fmt.Println("\n--- sync.Mutex ---")
	
	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup
	
	// Launch 100 goroutines to increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("somekey")
		}()
	}
	
	wg.Wait()
	fmt.Printf("Counter value: %d (Expected: 100)\n", c.Value("somekey"))
}

// ─────── 3. SYNC.ONCE ───────
// sync.Once ensures that a piece of code is executed exactly ONCE, regardless of
// how many goroutines try to call it simultaneously.
// Often used for lazy initialization of singletons or expensive resources.

var (
	once sync.Once
	expensiveResource string
)

func initResource() {
	fmt.Println("Initializing expensive resource... (This should only print ONCE)")
	time.Sleep(50 * time.Millisecond)
	expensiveResource = "Ready!"
}

func demonstrateOnce() {
	fmt.Println("\n--- sync.Once ---")
	
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Multiple goroutines race to call this, but initResource only runs once.
			once.Do(initResource)
			fmt.Printf("Goroutine %d sees: %s\n", id, expensiveResource)
		}(i)
	}
	
	wg.Wait()
}
