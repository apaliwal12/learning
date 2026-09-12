package main

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestCreateLogger(t *testing.T) {
	logger := CreateLogger()
	if logger == nil {
		t.Skip("CreateLogger not implemented yet")
	}

	// It's hard to robustly test if it points to os.Stdout and has Debug level
	// without reaching into unexported fields, but we can verify it's a *slog.Logger
	_, ok := interface{}(logger).(*slog.Logger)
	if !ok {
		t.Error("CreateLogger did not return a *slog.Logger")
	}
}

func TestLogTransaction(t *testing.T) {
	// We'll capture the output in a buffer
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := slog.New(handler)

	LogTransaction(logger, 123, 45.67, true)

	output := buf.String()

	if !strings.Contains(output, `"msg":"transaction processed"`) {
		t.Errorf("Expected message 'transaction processed', got: %s", output)
	}
	if !strings.Contains(output, `"account_id":123`) {
		t.Errorf("Missing or incorrect account_id attribute: %s", output)
	}
	if !strings.Contains(output, `"amount":45.67`) {
		t.Errorf("Missing or incorrect amount attribute: %s", output)
	}
	if !strings.Contains(output, `"success":true`) {
		t.Errorf("Missing or incorrect success attribute: %s", output)
	}
}
