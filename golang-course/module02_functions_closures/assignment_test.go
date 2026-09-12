package main

import (
	"errors"
	"strings"
	"testing"
)

func TestMultipleReturns(t *testing.T) {
	tests := []struct {
		s       string
		wantLen int
	}{
		{"hello", 5},
		{"", 0},
		{"gopher", 6},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			l, s := MultipleReturns(tt.s)
			if l != tt.wantLen || s != tt.s {
				t.Errorf("MultipleReturns(%q) = %d, %q; want %d, %q", tt.s, l, s, tt.wantLen, tt.s)
			}
		})
	}
}

func TestSumVariadic(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Empty", []int{}, 0},
		{"One", []int{5}, 5},
		{"Multiple", []int{1, 2, 3, 4}, 10},
		{"Negative", []int{-1, -2, 3}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumVariadic(tt.nums...)
			if got != tt.want {
				t.Errorf("SumVariadic(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMultiplyBy(t *testing.T) {
	doubler := MultiplyBy(2)
	tripler := MultiplyBy(3)

	if got := doubler(5); got != 10 {
		t.Errorf("MultiplyBy(2)(5) = %d; want 10", got)
	}
	if got := tripler(5); got != 15 {
		t.Errorf("MultiplyBy(3)(5) = %d; want 15", got)
	}
}

func TestExecuteWithCleanup(t *testing.T) {
	cleanupCalled := false
	mainFunc := func() {
		panic("boom")
	}
	cleanupFunc := func() {
		cleanupCalled = true
	}

	// We have to recover here because ExecuteWithCleanup doesn't suppress the panic itself
	func() {
		defer func() {
			recover()
		}()
		ExecuteWithCleanup(mainFunc, cleanupFunc)
	}()

	if !cleanupCalled {
		t.Error("Cleanup function was not called after panic")
	}

	cleanupCalled = false
	ExecuteWithCleanup(func() {}, cleanupFunc)
	if !cleanupCalled {
		t.Error("Cleanup function was not called after normal execution")
	}
}

func TestFilterWords(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "date"}

	containsA := func(s string) bool { return strings.Contains(s, "a") }
	has5Letters := func(s string) bool { return len(s) == 5 }

	t.Run("Contains A", func(t *testing.T) {
		got := FilterWords(words, containsA)
		want := []string{"apple", "banana", "date"}
		if len(got) != len(want) {
			t.Fatalf("Got %v; want %v", got, want)
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("Got %v; want %v", got, want)
				break
			}
		}
	})

	t.Run("5 Letters", func(t *testing.T) {
		got := FilterWords(words, has5Letters)
		want := []string{"apple"}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("Got %v; want %v", got, want)
		}
	})
}

func TestSafeExecute(t *testing.T) {
	t.Run("Normal execution", func(t *testing.T) {
		err := SafeExecute(func() {})
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})

	t.Run("Panic string", func(t *testing.T) {
		err := SafeExecute(func() { panic("database connection failed") })
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
		if !strings.Contains(err.Error(), "database connection failed") {
			t.Errorf("Error message %q does not contain panic value", err.Error())
		}
	})

	t.Run("Panic error", func(t *testing.T) {
		err := SafeExecute(func() { panic(errors.New("benchmark error")) }) // Just an example error
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func TestMemoizeFibonacci(t *testing.T) {
	fib := MemoizeFibonacci()

	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{10, 55},
		{40, 102334155}, // This would be very slow without memoization
	}

	for _, tt := range tests {
		got := fib(tt.n)
		if got != tt.want {
			t.Errorf("Fib(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}
