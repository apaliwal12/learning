package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// ─────── PROJECT B: TCP LOAD BALANCER ───────
// This project synthesizes:
// - Low-level networking (net.Listener, net.Conn)
// - Concurrency & sync.Mutex (for round-robin state)
// - io.Copy (for fast data streaming)
//
// A TCP Load Balancer accepts incoming connections on a single port,
// and proxies the raw byte stream to one of several backend servers.

// 1. Mock Backend Servers
// We need servers to balance traffic to!
func startMockBackend(port string, name string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start backend %s: %v", name, err)
	}
	
	log.Printf("Backend %s listening on %s\n", name, port)
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		
		// When a client connects, we just send a greeting and close it.
		// (In a real app, this would be an HTTP server or database)
		go func(c net.Conn) {
			defer c.Close()
			msg := fmt.Sprintf("Hello from %s!\n", name)
			c.Write([]byte(msg))
		}(conn)
	}
}

// 2. The Load Balancer State
type LoadBalancer struct {
	backends []string
	current  int
	mu       sync.Mutex
}

// NextBackend implements a Thread-Safe Round-Robin algorithm.
func (lb *LoadBalancer) NextBackend() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	backend := lb.backends[lb.current]
	lb.current = (lb.current + 1) % len(lb.backends)
	return backend
}

// 3. The Proxy Logic
func handleConnection(clientConn net.Conn, lb *LoadBalancer) {
	defer clientConn.Close()
	
	targetBackend := lb.NextBackend()
	fmt.Printf("Routing request from %s to %s\n", clientConn.RemoteAddr(), targetBackend)
	
	// Connect to the chosen backend
	backendConn, err := net.Dial("tcp", targetBackend)
	if err != nil {
		log.Printf("Failed to connect to backend %s: %v\n", targetBackend, err)
		return
	}
	defer backendConn.Close()
	
	// We need to stream data in BOTH directions concurrently.
	var wg sync.WaitGroup
	wg.Add(2)
	
	// Copy from Client -> Backend
	go func() {
		defer wg.Done()
		io.Copy(backendConn, clientConn)
	}()
	
	// Copy from Backend -> Client
	go func() {
		defer wg.Done()
		io.Copy(clientConn, backendConn)
	}()
	
	// Wait for both streams to close
	wg.Wait()
}

func main() {
	fmt.Println("--- Project B: TCP Load Balancer ---")
	
	// 1. Start 3 mock backends in the background
	go startMockBackend(":8001", "Backend-A")
	go startMockBackend(":8002", "Backend-B")
	go startMockBackend(":8003", "Backend-C")
	
	time.Sleep(100 * time.Millisecond) // Give them time to start
	
	// 2. Initialize the Load Balancer
	lb := &LoadBalancer{
		backends: []string{"localhost:8001", "localhost:8002", "localhost:8003"},
	}
	
	// 3. Start the Load Balancer Listener
	lbPort := ":8080"
	ln, err := net.Listen("tcp", lbPort)
	if err != nil {
		log.Fatalf("Failed to start Load Balancer: %v", err)
	}
	
	log.Printf("Load Balancer listening on %s (Round-Robin)...\n", lbPort)
	
	// 4. Accept incoming connections
	// (We will run this in a goroutine so we can demonstrate a test client)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println("Accept error:", err)
				continue
			}
			go handleConnection(conn, lb)
		}
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	// 5. Test it! Make 4 requests to the Load Balancer.
	// You should see them hit A, B, C, then A again!
	for i := 0; i < 4; i++ {
		conn, _ := net.Dial("tcp", "localhost:8080")
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Printf("Received response: %s", string(buf[:n]))
		conn.Close()
		time.Sleep(10 * time.Millisecond)
	}
}
