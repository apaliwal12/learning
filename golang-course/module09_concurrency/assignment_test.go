package main

import (
	"testing"
	"time"
)

func TestAsyncSquare(t *testing.T) {
	ch := AsyncSquare(5)
	if ch == nil {
		t.Fatal("AsyncSquare returned nil channel")
	}

	select {
	case val, ok := <-ch:
		if !ok {
			t.Fatal("Channel closed before sending value")
		}
		if val != 25 {
			t.Errorf("Expected 25, got %d", val)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for value")
	}

	// Verify it was closed
	select {
	case _, ok := <-ch:
		if ok {
			t.Error("Channel was not closed after sending")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Timeout waiting for channel close")
	}
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	out := FanIn(ch1, ch2)
	if out == nil {
		t.Fatal("FanIn returned nil channel")
	}

	go func() {
		ch1 <- 1
		ch2 <- 2
		ch1 <- 3
	}()

	results := make(map[int]bool)
	for i := 0; i < 3; i++ {
		select {
		case val := <-out:
			results[val] = true
		case <-time.After(1 * time.Second):
			t.Fatal("Timeout waiting for values")
		}
	}

	if !results[1] || !results[2] || !results[3] {
		t.Errorf("Did not receive all expected values, got: %v", results)
	}
}

func TestExecuteWithTimeout(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		fastWork := func() {
			time.Sleep(10 * time.Millisecond)
		}
		success := ExecuteWithTimeout(fastWork, 100*time.Millisecond)
		if !success {
			t.Error("Expected success=true for fast work")
		}
	})

	t.Run("Timeout case", func(t *testing.T) {
		slowWork := func() {
			time.Sleep(200 * time.Millisecond)
		}
		success := ExecuteWithTimeout(slowWork, 50*time.Millisecond)
		if success {
			t.Error("Expected success=false for slow work")
		}
	})
}
