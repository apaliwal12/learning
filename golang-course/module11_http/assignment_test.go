package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchStatus(t *testing.T) {
	// Create a test server that returns 201 Created
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer ts.Close()

	status, err := FetchStatus(ts.URL)
	if err != nil {
		t.Fatalf("FetchStatus error: %v", err)
	}
	if status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, status)
	}
}

func TestFetchWithContextAndTimeout(t *testing.T) {
	// Create a test server that sleeps for 50ms before responding
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	t.Run("Timeout triggers", func(t *testing.T) {
		// Set timeout to 10ms (shorter than server sleep)
		_, err := FetchWithContextAndTimeout(context.Background(), ts.URL, 10*time.Millisecond)
		if err == nil {
			t.Error("Expected timeout error, got nil")
		}
	})

	t.Run("Success", func(t *testing.T) {
		// Set timeout to 100ms (longer than server sleep)
		status, err := FetchWithContextAndTimeout(context.Background(), ts.URL, 100*time.Millisecond)
		if err != nil {
			t.Fatalf("Expected success, got error: %v", err)
		}
		if status != http.StatusOK {
			t.Errorf("Expected status 200, got %d", status)
		}
	})
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HealthHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}
	
	// Optional: we can check the body
	/*
	body, _ := io.ReadAll(res.Body)
	want := `{"status":"ok"}`
	if strings.TrimSpace(string(body)) != want {
		t.Errorf("Expected body %q, got %q", want, body)
	}
	*/
}
