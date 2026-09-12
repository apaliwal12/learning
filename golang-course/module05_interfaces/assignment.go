package main

// TODO 1: Implement the DataStore interface
// It should have two methods:
// - Save(key string, value string) error
// - Load(key string) (string, error)
type DataStore interface {
	// Fix me
}

// TODO 2: Implement InMemoryStore
// Create a struct InMemoryStore that has a map[string]string field.
// It must implement the DataStore interface.
// Save should store the key-value pair.
// Load should return the value if it exists, otherwise return an error "not found".

// TODO 3: Implement DBError
// Create a custom error struct DBError with fields:
// - Code int
// - Message string
// It must implement the error interface. The Error() method should return:
// "DB Error [Code]: Message"

// TODO 4: Implement ExtractDBErrorCode
// Write a function that takes an `error` as an argument.
// If the error is of type DBError (use a type assertion), return its Code and true.
// Otherwise, return 0 and false.
func ExtractDBErrorCode(err error) (int, bool) {
	return 0, false // Fix me
}
