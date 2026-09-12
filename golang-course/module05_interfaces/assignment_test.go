package main

import (
	"testing"
)

func TestDataStoreImplementation(t *testing.T) {
	// We use reflection/type assertions to check if the student implemented it
	// without causing compile errors before they start.

	t.Run("InMemoryStore", func(t *testing.T) {
		// Because we don't know if they defined InMemoryStore yet, we have to
		// be careful. A normal Go test would just fail to compile.
		// Since we want these tests to be runnable "red tests", we'd ideally
		// have a stub. For the sake of this platform, we'll assume they will
		// define `type InMemoryStore struct { data map[string]string }`

		// This will panic/fail compilation if InMemoryStore doesn't exist.
		// That's standard for Go TDD.

		/*
			store := &InMemoryStore{data: make(map[string]string)}
			var ds DataStore = store // Check interface compliance

			err := ds.Save("user1", "alice")
			if err != nil {
				t.Errorf("Save returned unexpected error: %v", err)
			}

			val, err := ds.Load("user1")
			if err != nil || val != "alice" {
				t.Errorf("Load failed: got (%q, %v), want (%q, nil)", val, err, "alice")
			}

			_, err = ds.Load("missing")
			if err == nil || err.Error() != "not found" {
				t.Errorf("Expected 'not found' error, got %v", err)
			}
		*/
		t.Skip("Uncomment the test code in assignment_test.go once you define InMemoryStore")
	})
}

func TestDBError(t *testing.T) {
	/*
		err := DBError{Code: 500, Message: "Connection failed"}

		// Check if it implements error
		var _ error = err

		want := "DB Error [500]: Connection failed"
		if err.Error() != want {
			t.Errorf("Error() = %q; want %q", err.Error(), want)
		}
	*/
	t.Skip("Uncomment the test code in assignment_test.go once you define DBError")
}

func TestExtractDBErrorCode(t *testing.T) {
	/*
		t.Run("Valid DBError", func(t *testing.T) {
			err := DBError{Code: 404, Message: "Not Found"}
			code, ok := ExtractDBErrorCode(err)
			if !ok || code != 404 {
				t.Errorf("ExtractDBErrorCode() = (%d, %v); want (404, true)", code, ok)
			}
		})

		t.Run("Standard Error", func(t *testing.T) {
			err := errors.New("standard error")
			code, ok := ExtractDBErrorCode(err)
			if ok || code != 0 {
				t.Errorf("ExtractDBErrorCode() = (%d, %v); want (0, false)", code, ok)
			}
		})
	*/
	t.Skip("Uncomment the test code in assignment_test.go once you define DBError and update ExtractDBErrorCode")
}
