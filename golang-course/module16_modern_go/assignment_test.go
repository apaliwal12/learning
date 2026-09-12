package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("Generics", func(t *testing.T) {
		t.Skip("Uncomment once Filter is implemented")
		/*
			ints := []int{1, 2, 3, 4, 5}
			evens := Filter(ints, func(x int) bool { return x%2 == 0 })
			if !reflect.DeepEqual(evens, []int{2, 4}) {
				t.Errorf("Filter failed for ints. Got: %v", evens)
			}

			strs := []string{"apple", "banana", "cherry"}
			hasA := Filter(strs, func(s string) bool { return s[0] == 'a' })
			if !reflect.DeepEqual(hasA, []string{"apple"}) {
				t.Errorf("Filter failed for strings. Got: %v", hasA)
			}
		*/
	})
}

func TestMap(t *testing.T) {
	t.Run("Generics", func(t *testing.T) {
		t.Skip("Uncomment once Map is implemented")
		/*
			ints := []int{1, 2, 3}
			strs := Map(ints, func(x int) string {
				return string(rune('A' + x - 1)) // 1->A, 2->B, 3->C
			})
			if !reflect.DeepEqual(strs, []string{"A", "B", "C"}) {
				t.Errorf("Map failed. Got: %v", strs)
			}
		*/
	})
}

func TestRangeIterator(t *testing.T) {
	seq := Range(2, 6)

	var got []int
	for val := range seq {
		got = append(got, val)
	}

	want := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Range(2, 6) yielded %v, want %v", got, want)
	}

	// Test early break
	var gotEarly []int
	for val := range seq {
		gotEarly = append(gotEarly, val)
		if val == 3 {
			break
		}
	}
	wantEarly := []int{2, 3}
	if !reflect.DeepEqual(gotEarly, wantEarly) {
		t.Errorf("Early break yielded %v, want %v", gotEarly, wantEarly)
	}
}
