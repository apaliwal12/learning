package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Café", "éfaC"}, // Multibyte test
		{"🌍🌎🌏", "🌏🌎🌍"}, // Emoji test
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := ReverseString(tt.s)
			if got != tt.want {
				t.Errorf("ReverseString(%q) = %q; want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"Racecar", true},
		{"hello", false},
		{"taco cat", true},
		{"A man a plan a canal Panama", true}, // Need to handle spaces and case
		{"", true},
		{"a", true},
		{"ab", false},
		{"Käse", false}, // Multibyte
		{"éé", true},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := IsPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestBuildGreeting(t *testing.T) {
	tests := []struct {
		names []string
		want  string
	}{
		{[]string{}, "Hello there!"},
		{[]string{"Alice"}, "Hello Alice!"},
		{[]string{"Alice", "Bob"}, "Hello Alice and Bob!"},
		{[]string{"Alice", "Bob", "Charlie"}, "Hello Alice, Bob, and Charlie!"},
		{[]string{"A", "B", "C", "D"}, "Hello A, B, C, and D!"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := BuildGreeting(tt.names)
			if got != tt.want {
				t.Errorf("BuildGreeting(%v) = %q; want %q", tt.names, got, tt.want)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"hello", 2},
		{"", 0},
		{"xyz", 0},
		{"A E I O U", 5},
		{"aeiou", 5},
		{"bAnAnA", 3},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := CountVowels(tt.s)
			if got != tt.want {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestFirstNRunes(t *testing.T) {
	tests := []struct {
		s    string
		n    int
		want string
	}{
		{"hello", 2, "he"},
		{"hello", 10, "hello"},
		{"", 5, ""},
		{"Café", 3, "Caf"},
		{"Café", 4, "Café"},
		{"🌍🌎🌏", 2, "🌍🌎"},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := FirstNRunes(tt.s, tt.n)
			if got != tt.want {
				t.Errorf("FirstNRunes(%q, %d) = %q; want %q", tt.s, tt.n, got, tt.want)
			}
		})
	}
}
