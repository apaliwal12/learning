package main

import "fmt"

// ─────── 12. BUILDING COLLECTIONS ───────
// Go's standard library doesn't have Sets or Queues built-in.
// We build them using slices and maps.

// Set Implementation (using map[T]struct{})
// We use `struct{}` (an empty struct) as the value because it takes exactly 0 bytes of memory.
// Using a boolean (map[T]bool) takes 1 byte per entry.

func demonstrateSet() {
	fmt.Println("\n--- Sets (map[string]struct{}) ---")
	
	// Create a set
	visited := make(map[string]struct{})
	
	// Add to set
	visited["/home"] = struct{}{}
	visited["/about"] = struct{}{}

	// Check if in set (we only care about 'ok', not the value)
	_, ok := visited["/home"]
	fmt.Printf("/home visited? %v\n", ok)

	_, ok = visited["/contact"]
	fmt.Printf("/contact visited? %v\n", ok)
}

// Queue Implementation (using slice)
// Enqueue: append(q, value)
// Dequeue: q[0], q = q[1:]
// Note: Dequeueing from a slice using `q[1:]` keeps the underlying array in memory.
// For large, long-lived queues, this can cause memory leaks or capacity exhaustion.
// For robust queues, you often need a ring buffer or a linked list (container/list).

func demonstrateQueue() {
	fmt.Println("\n--- Queue (using slice) ---")
	
	var queue []string
	
	// Enqueue
	queue = append(queue, "task1", "task2", "task3")
	fmt.Printf("Queue: %v\n", queue)

	// Dequeue
	if len(queue) > 0 {
		head := queue[0]
		queue = queue[1:] // Reslice to remove head
		fmt.Printf("Dequeued: %s, Remaining: %v\n", head, queue)
	}
}
