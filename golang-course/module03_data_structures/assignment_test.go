package main

import (
	"maps"
	"slices"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		index int
		want  []int
	}{
		{"Remove middle", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{"Remove first", []int{1, 2, 3}, 0, []int{2, 3}},
		{"Remove last", []int{1, 2, 3}, 2, []int{1, 2}},
		{"Index out of bounds positive", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"Index out of bounds negative", []int{1, 2, 3}, -1, []int{1, 2, 3}},
		{"Empty slice", []int{}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clone original to ensure it's not modified
			orig := slices.Clone(tt.slice)

			got := RemoveAtIndex(tt.slice, tt.index)

			if !slices.Equal(got, tt.want) {
				t.Errorf("RemoveAtIndex() = %v, want %v", got, tt.want)
			}

			// Verify original wasn't mutated
			if !slices.Equal(tt.slice, orig) {
				t.Errorf("Original slice was mutated! %v != %v", tt.slice, orig)
			}
		})
	}
}

func TestFilterMap(t *testing.T) {
	m := map[string]int{
		"a": 10,
		"b": 20,
		"c": 5,
		"d": 30,
	}

	got := FilterMap(m, 15)
	want := map[string]int{"b": 20, "d": 30}

	if !maps.Equal(got, want) {
		t.Errorf("FilterMap() = %v, want %v", got, want)
	}
}

func TestCountFrequencies(t *testing.T) {
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	got := CountFrequencies(words)
	want := map[string]int{
		"apple":  3,
		"banana": 2,
		"cherry": 1,
	}

	if !maps.Equal(got, want) {
		t.Errorf("CountFrequencies() = %v, want %v", got, want)
	}
}

// We use type assertions in the test to ensure the student actually
// implemented the methods correctly, even if they aren't exported interfaces.
func TestRectangle(t *testing.T) {
	// A small trick to test methods if the struct is implemented
	type AreaCalculator interface {
		Area() float64
	}
	type Scaler interface {
		Scale(float64)
	}

	// This code will fail to compile if Rectangle or its methods are missing.
	// We wrap it in a defer/recover just in case, but compilation failure is the actual test here.
	t.Run("Area", func(t *testing.T) {
		r := Rectangle{
			// Width:  5, // TODO: Uncomment this
			// Height: 10, // TODO: Uncomment this
		}

		// Interface conversion check
		var ac AreaCalculator
		// We have to use a workaround to check if r implements AreaCalculator without compile errors
		// if the user hasn't written it yet.
		// Since we want the code to compile even before they do the assignment,
		// we skip strict compilation checks and just verify if they implemented it.
		// Actually, standard Go tests will just fail to compile if Rectangle isn't defined.
		// Since we stubbed `type Rectangle struct{}`, it will compile, but interface assertion will fail.

		ac, ok := any(r).(AreaCalculator)
		if !ok {
			t.Fatalf("Rectangle does not implement Area() float64")
		}

		got := ac.Area()
		if got != 50.0 {
			t.Errorf("Area() = %f; want 50.0", got)
		}
	})

	t.Run("Scale", func(t *testing.T) {
		r := Rectangle{
			// Width:  5, // TODO: Uncomment this
			// Height: 10, // TODO: Uncomment this
		}

		s, ok := any(&r).(Scaler)
		if !ok {
			t.Fatalf("*Rectangle does not implement Scale(float64)")
		}

		s.Scale(2.0)

		// To check fields without knowing if they exist, we just hope the assignment works.
		// If they named fields differently, Area() will have to reflect that anyway.
		ac := any(r).(AreaCalculator)
		if ac.Area() != 200.0 { // (5*2) * (10*2) = 10 * 20 = 200
			t.Errorf("After Scale(2), Area() = %f; want 200.0", ac.Area())
		}
	})
}
