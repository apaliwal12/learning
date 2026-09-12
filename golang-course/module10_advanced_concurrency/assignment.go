package main

import (
	"context"
)

// TODO 1: Implement RunAll
// It should take a slice of functions `tasks`.
// It must run ALL of them concurrently.
// It MUST block and wait until all of them have completely finished before returning.
// Hint: use sync.WaitGroup.
func RunAll(tasks []func()) {
	// Fix me
}

// TODO 2: Define a thread-safe string map `ConcurrentMap`
// It should have methods:
// - Set(key, value string)
// - Get(key string) (string, bool)
// Hint: use a struct with a map and a sync.RWMutex (or sync.Mutex).
type ConcurrentMap struct {
	// Fix me
}

// func (m *ConcurrentMap) Set(key, value string) { ... }
// func (m *ConcurrentMap) Get(key string) (string, bool) { ... }

// TODO 3: Implement FetchWithContext
// It takes a context and a `fetch` function (which simulates a slow network request).
// It should run the `fetch` function in a goroutine.
// If `fetch` completes normally, return its result and nil error.
// If the context is canceled or times out before `fetch` completes, 
// return an empty string and the context's error (ctx.Err()).
func FetchWithContext(ctx context.Context, fetch func() string) (string, error) {
	return "", nil // Fix me
}
