package main

import (
	"fmt"
	"time"
)

// ─────── 2. CHANNELS ───────
// "Do not communicate by sharing memory; instead, share memory by communicating."
//
// Channels are typed conduits through which you can send and receive values
// between goroutines. They act as synchronization points, preventing race conditions
// naturally without needing explicit locks (mutexes).

func demonstrateChannels() {
	fmt.Println("\n--- Channels (Unbuffered) ---")
	
	// Unbuffered channel: The sender blocks until the receiver is ready,
	// and the receiver blocks until the sender is ready. They synchronize exactly.
	ch := make(chan string)
	
	go func() {
		fmt.Println("[Goroutine] Doing some work...")
		time.Sleep(20 * time.Millisecond)
		
		// Send a value INTO the channel
		ch <- "Work Complete!"
		fmt.Println("[Goroutine] Value sent.")
	}()
	
	fmt.Println("[Main] Waiting for channel data...")
	// Receive a value FROM the channel (This blocks!)
	msg := <-ch 
	fmt.Printf("[Main] Received: %s\n", msg)
}

// ─────── 3. BUFFERED CHANNELS ───────
// Buffered channels have a capacity.
// Sends only block when the buffer is full.
// Receives block when the buffer is empty.

func demonstrateBufferedChannels() {
	fmt.Println("\n--- Buffered Channels ---")
	
	// Buffered channel with capacity 2
	ch := make(chan int, 2)
	
	// We can send 2 values without a receiver ready!
	ch <- 1
	ch <- 2
	// ch <- 3 // If we did this, we would deadlock because the buffer is full.
	
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

// ─────── 4. CLOSING CHANNELS & RANGE ───────
// Senders can close a channel to indicate that no more values will be sent.
// Receivers can test whether a channel has been closed by assigning a second parameter: `v, ok := <-ch`
// A closed channel never blocks. Receiving from a closed channel returns the zero value immediately.

func demonstrateChannelClose() {
	fmt.Println("\n--- Closing Channels ---")
	
	ch := make(chan int, 3)
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		// ONLY the sender should ever close a channel!
		// Sending to a closed channel causes a panic.
		close(ch)
	}()
	
	// `for range` on a channel receives values until the channel is closed.
	for val := range ch {
		fmt.Println("Received:", val)
	}
	fmt.Println("Channel closed, loop terminated.")
}
