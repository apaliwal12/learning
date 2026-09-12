package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// ─────── PROJECT D: KEY-VALUE STORE WITH WAL ───────
// This project synthesizes:
// - File I/O (Module 7)
// - Concurrency / sync.RWMutex (Module 9)
// - Database Concepts (Module 12)
//
// A Write-Ahead Log (WAL) is how real databases (like Postgres or Redis AOF)
// survive crashes. Every mutation is written to an append-only file BEFORE
// it is applied to the in-memory state.

type KVStore struct {
	mu   sync.RWMutex
	data map[string]string
	wal  *os.File
}

// NewKVStore creates or recovers a store from the WAL file.
func NewKVStore(walPath string) (*KVStore, error) {
	store := &KVStore{
		data: make(map[string]string),
	}

	// 1. Open the file in Append/Create mode
	file, err := os.OpenFile(walPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	store.wal = file

	// 2. Recover state from the WAL
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ",", 3)
		if len(parts) < 2 {
			continue
		}
		
		op := parts[0]
		key := parts[1]
		
		switch op {
		case "SET":
			if len(parts) == 3 {
				store.data[key] = parts[2]
			}
		case "DEL":
			delete(store.data, key)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to recover from WAL: %w", err)
	}

	log.Printf("Recovered %d keys from WAL.\n", len(store.data))
	return store, nil
}

// Set saves a value. It MUST write to the WAL first!
func (s *KVStore) Set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 1. Write to WAL
	entry := fmt.Sprintf("SET,%s,%s\n", key, value)
	if _, err := s.wal.WriteString(entry); err != nil {
		return err
	}
	// Force write to disk (sync). In a real DB, you might batch these for performance.
	if err := s.wal.Sync(); err != nil {
		return err
	}

	// 2. Apply to memory
	s.data[key] = value
	return nil
}

// Get retrieves a value.
func (s *KVStore) Get(key string) (string, bool) {
	s.mu.RLock() // Use RLock because we are only reading!
	defer s.mu.RUnlock()
	
	val, exists := s.data[key]
	return val, exists
}

// Close gracefully closes the WAL.
func (s *KVStore) Close() error {
	return s.wal.Close()
}

func main() {
	fmt.Println("--- Project D: Key-Value Store with WAL ---")
	
	walPath := "database.wal"
	
	// Create/Recover the store
	store, err := NewKVStore(walPath)
	if err != nil {
		log.Fatalf("Failed to initialize store: %v", err)
	}
	defer store.Close()
	
	// Print current state
	val, ok := store.Get("user_1")
	if ok {
		fmt.Printf("Recovered state: user_1 = %s\n", val)
	} else {
		fmt.Println("No previous state for user_1.")
	}
	
	// Write new data
	fmt.Println("Setting user_1 = Alice...")
	if err := store.Set("user_1", "Alice"); err != nil {
		log.Fatalf("Failed to set: %v", err)
	}
	
	fmt.Println("Check the `database.wal` file! Next time you run this, it will recover Alice automatically.")
}
