package main

import (
	"errors"
	"strings"
	"testing"
)

func TestFetchResource(t *testing.T) {
	err := FetchResource("123")
	
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	expectedPrefix := "FetchResource failed:"
	if !strings.HasPrefix(err.Error(), expectedPrefix) {
		t.Errorf("Error string %q does not start with %q", err.Error(), expectedPrefix)
	}

	if !errors.Is(err, ErrNotFound) {
		t.Error("The returned error does NOT wrap ErrNotFound (did you use %w?)")
	}
}

func TestIsNotFoundError(t *testing.T) {
	// True case
	wrappedErr := errors.Join(errors.New("some context"), ErrNotFound)
	if !IsNotFoundError(wrappedErr) {
		t.Error("IsNotFoundError returned false for a wrapped ErrNotFound")
	}

	// False case
	otherErr := errors.New("a completely different error")
	if IsNotFoundError(otherErr) {
		t.Error("IsNotFoundError returned true for an unrelated error")
	}
}
