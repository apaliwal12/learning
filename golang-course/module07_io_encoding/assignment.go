package main

import (
	"io"
)

// TODO 1: Implement CountLines
// It should take an io.Reader and return the number of newline characters ('\n') in it.
// You should read the data in chunks (e.g., 1024 bytes at a time) rather than
// reading the entire thing into memory at once.
func CountLines(r io.Reader) (int, error) {
	return 0, nil // Fix me
}

// TODO 2: Define a struct named `Person`
// It should have two fields: Name (string) and Age (int).
// Add struct tags so it marshals to JSON as: `{"name": "...", "age": ...}`
// type Person struct { ... }

// TODO 3: Implement WritePeopleToJSON
// It should take a slice of Person and an io.Writer.
// It must encode the slice to JSON and write it directly to the io.Writer.
// Hint: Use json.NewEncoder
// func WritePeopleToJSON(people []Person, w io.Writer) error { ... }

// TODO 4: Implement AppendToFile
// It should take a filename and a string of data.
// It must append the data to the file, creating the file if it doesn't exist.
// Permissions should be 0644.
func AppendToFile(filename, data string) error {
	return nil // Fix me
}
