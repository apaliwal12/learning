package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-5, 5, 0},
		{10, -5, 5},
	}

	for _, tt := range tests {
		t.Run("Add", func(t *testing.T) {
			got := AddTwoNumbers(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("AddTwoNumbers(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{-5, false},
	}

	for _, tt := range tests {
		t.Run("IsEven", func(t *testing.T) {
			got := IsEven(tt.n)
			if got != tt.want {
				t.Errorf("IsEven(%d) = %v; want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		wantValue int
		wantErr   bool
	}{
		{"Normal division", 10, 2, 5, false},
		{"Division by negative", 10, -2, -5, false},
		{"Zero numerator", 0, 5, 0, false},
		{"Division by zero", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := SafeDivide(tt.a, tt.b)
			hasErr := err != nil

			if hasErr != tt.wantErr {
				t.Errorf("SafeDivide(%d, %d) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
			}
			if !hasErr && gotValue != tt.wantValue {
				t.Errorf("SafeDivide(%d, %d) = %d; want %d", tt.a, tt.b, gotValue, tt.wantValue)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	a, b := 10, 20
	Swap(&a, &b)
	if a != 20 || b != 10 {
		t.Errorf("Swap failed: expected a=20, b=10; got a=%d, b=%d", a, b)
	}
}

func TestSumUpTo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 15},
		{10, 55},
		{1, 1},
		{0, 0},
		{-5, 0},
	}

	for _, tt := range tests {
		t.Run("SumUpTo", func(t *testing.T) {
			got := SumUpTo(tt.n)
			if got != tt.want {
				t.Errorf("SumUpTo(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestDetermineGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{90, "A"},
		{85, "B"},
		{80, "B"},
		{75, "C"},
		{70, "C"},
		{65, "D"},
		{60, "D"},
		{50, "F"},
		{0, "F"},
		{-10, "Invalid"},
		{110, "Invalid"},
	}

	for _, tt := range tests {
		t.Run("DetermineGrade", func(t *testing.T) {
			got := DetermineGrade(tt.score)
			if got != tt.want {
				t.Errorf("DetermineGrade(%d) = %s; want %s", tt.score, got, tt.want)
			}
		})
	}
}
