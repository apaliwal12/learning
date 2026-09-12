package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ─────── WORKER POOL PROJECT ───────
// A Worker Pool limits the number of concurrent goroutines running at once.
// If you have 10,000 tasks, you don't necessarily want 10,000 goroutines hitting
// a database simultaneously. A worker pool of 10 goroutines can chew through
// the 10,000 tasks methodically.

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
}

// worker reads jobs from the `jobs` channel, processes them, and sends the
// result to the `results` channel.
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j.ID)
		
		// Simulate expensive work
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		
		output := j.Value * 2
		results <- Result{JobID: j.ID, Output: output}
	}
	fmt.Printf("Worker %d shutting down (job channel closed)\n", id)
}

func main() {
	fmt.Println("--- Worker Pool Demo ---")
	
	const numJobs = 10
	const numWorkers = 3
	
	// Create buffered channels
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	// 1. Start the workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}
	
	// 2. Send jobs to the pool
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j * 10}
	}
	
	// 3. Close the jobs channel
	// This signals the workers that no more jobs are coming.
	// They will finish their current job and then exit their `for range` loop.
	close(jobs)
	
	// 4. Collect results
	// We know exactly how many jobs we sent, so we can loop that many times.
	// (In Module 10, we'll learn how to use sync.WaitGroup to do this more robustly).
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("Result collected: Job %d -> %d\n", res.JobID, res.Output)
	}
	
	fmt.Println("All jobs completed!")
}
