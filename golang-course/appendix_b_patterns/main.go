package main

import (
	"fmt"
	"time"
)

// ─────── APPENDIX B: PATTERNS COOKBOOK ───────
// Go has a unique set of idiomatic patterns. Here are the most famous ones.

// 1. FUNCTIONAL OPTIONS PATTERN
// Problem: You have a struct with many optional configuration parameters.
// Solution: Pass functions that modify the struct, rather than a giant config object.

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

// Option is a function that modifies a Server
type Option func(*Server)

// WithPort is an Option builder
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

// WithTimeout is an Option builder
func WithTimeout(d time.Duration) Option {
	return func(s *Server) {
		s.Timeout = d
	}
}

// NewServer takes a variadic number of Options!
func NewServer(host string, options ...Option) *Server {
	// 1. Set defaults
	s := &Server{
		Host:    host,
		Port:    8080,
		Timeout: 30 * time.Second,
	}

	// 2. Apply all provided options
	for _, opt := range options {
		opt(s)
	}

	return s
}

func demonstrateFunctionalOptions() {
	fmt.Println("\n--- Functional Options ---")
	
	// Create with defaults
	s1 := NewServer("localhost")
	fmt.Printf("Default Server: %+v\n", s1)
	
	// Create with custom overrides
	s2 := NewServer("api.example.com", WithPort(443), WithTimeout(5*time.Second))
	fmt.Printf("Custom Server: %+v\n", s2)
}

func main() {
	fmt.Println("Welcome to Appendix B: Patterns Cookbook")
	demonstrateFunctionalOptions()
	
	// (Other patterns to research: The Singleton using sync.Once, The Worker Pool, The Generator)
}
