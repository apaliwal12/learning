package main

import (
	"os"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"hello", 0},
		{"hello\n", 1},
		{"hello\nworld\n", 2},
		{strings.Repeat("a\n", 2000), 2000}, // Tests chunking logic
	}

	for _, tt := range tests {
		t.Run(tt.input[:min(len(tt.input), 10)], func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := CountLines(r)
			if err != nil {
				t.Fatalf("CountLines returned error: %v", err)
			}
			if got != tt.want {
				t.Errorf("CountLines() = %d; want %d", got, tt.want)
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestWritePeopleToJSON(t *testing.T) {
	// Again, wrapping to avoid compile errors before struct is defined
	t.Run("JSON Encoding", func(t *testing.T) {
		t.Skip("Uncomment this test once Person and WritePeopleToJSON are implemented")
		/*
			people := []Person{
				{Name: "Alice", Age: 30},
				{Name: "Bob", Age: 25},
			}

			var buf bytes.Buffer
			err := WritePeopleToJSON(people, &buf)
			if err != nil {
				t.Fatalf("WritePeopleToJSON error: %v", err)
			}

			got := strings.TrimSpace(buf.String())
			want := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
			if got != want {
				// json.NewEncoder might add a trailing newline, which TrimSpace handles
				t.Errorf("WritePeopleToJSON() wrote %q; want %q", got, want)
			}
		*/
	})
}

func TestAppendToFile(t *testing.T) {
	filename := "test_append.txt"
	defer os.Remove(filename) // Cleanup

	err := AppendToFile(filename, "Line 1\n")
	if err != nil {
		t.Fatalf("AppendToFile(1) error: %v", err)
	}

	err = AppendToFile(filename, "Line 2\n")
	if err != nil {
		t.Fatalf("AppendToFile(2) error: %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	got := string(data)
	want := "Line 1\nLine 2\n"
	if got != want {
		t.Errorf("File contents = %q; want %q", got, want)
	}
}
