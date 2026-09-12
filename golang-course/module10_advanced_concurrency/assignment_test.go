package main

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunAll(t *testing.T) {
	var counter int32

	tasks := []func(){
		func() {
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt32(&counter, 1)
		},
		func() {
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt32(&counter, 1)
		},
		func() {
			time.Sleep(30 * time.Millisecond)
			atomic.AddInt32(&counter, 1)
		},
	}

	start := time.Now()
	RunAll(tasks)
	elapsed := time.Since(start)

	if counter != 3 {
		t.Errorf("Expected counter to be 3, got %d. Did not wait for all tasks?", counter)
	}

	// If it ran sequentially, it would take >= 60ms.
	// If it ran concurrently, it should take ~30ms.
	if elapsed >= 60*time.Millisecond {
		t.Errorf("Took too long (%v). Tasks were probably not run concurrently.", elapsed)
	}
}

func TestConcurrentMap(t *testing.T) {
	t.Run("Concurrent access", func(t *testing.T) {
		t.Skip("Uncomment when ConcurrentMap is implemented")
		/*
		m := &ConcurrentMap{data: make(map[string]string)} // Assume standard init
		// To truly test this without compile errors for stub code, we'd use interfaces.
		// For simplicity, we just skip it initially.
		*/
	})
}

func TestFetchWithContext(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		fastFetch := func() string {
			time.Sleep(10 * time.Millisecond)
			return "data"
		}

		res, err := FetchWithContext(ctx, fastFetch)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if res != "data" {
			t.Errorf("Expected 'data', got %q", res)
		}
	})

	t.Run("Timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		slowFetch := func() string {
			time.Sleep(100 * time.Millisecond)
			return "data"
		}

		res, err := FetchWithContext(ctx, slowFetch)
		if err == nil {
			t.Error("Expected context error, got nil")
		}
		if res != "" {
			t.Errorf("Expected empty string on error, got %q", res)
		}
	})
}
