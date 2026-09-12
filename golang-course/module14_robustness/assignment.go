package main

import (
	"errors"
)

// ErrNotFound is a sentinel error.
var ErrNotFound = errors.New("resource not found")

// TODO 1: Implement FetchResource
// It should return the ErrNotFound error, but WRAPPED with additional context.
// The error string should be: "FetchResource failed: resource not found"
// You MUST use fmt.Errorf with the %w verb.
func FetchResource(id string) error {
	return nil // Fix me
}

// TODO 2: Implement IsNotFoundError
// It should take an error.
// It should return true ONLY IF the error, or any error in its chain, is ErrNotFound.
// Use errors.Is.
func IsNotFoundError(err error) bool {
	return false // Fix me
}
