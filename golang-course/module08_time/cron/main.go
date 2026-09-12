package main

import (
	"fmt"
	"time"
)

// ─────── SCHEDULER PROJECT ───────
// This project demonstrates a simple in-memory cron-like scheduler.
// It can run tasks periodically using time.Ticker.

// Job represents a task to be executed.
type Job struct {
	Name     string
	Interval time.Duration
	Action   func()
}

// Scheduler manages and runs jobs.
type Scheduler struct {
	jobs []Job
}

// AddJob adds a new job to the scheduler.
func (s *Scheduler) AddJob(name string, interval time.Duration, action func()) {
	s.jobs = append(s.jobs, Job{Name: name, Interval: interval, Action: action})
}

// Start runs all jobs in the background.
// (We are using goroutines here slightly ahead of Module 09,
// but it's the only way to run multiple tickers concurrently!)
func (s *Scheduler) Start() {
	fmt.Println("Starting Scheduler...")
	for _, job := range s.jobs {
		// We capture the job variable for the goroutine
		j := job
		go func() {
			ticker := time.NewTicker(j.Interval)
			defer ticker.Stop() // Good practice
			
			for t := range ticker.C {
				fmt.Printf("[%s] Running job %q\n", t.Format("15:04:05"), j.Name)
				j.Action()
			}
		}()
	}
}

func main() {
	s := &Scheduler{}

	s.AddJob("Ping Database", 1*time.Second, func() {
		// Simulate DB ping
	})

	s.AddJob("Cleanup Temp Files", 3*time.Second, func() {
		// Simulate cleanup
	})
	
	s.AddJob("Send Metric Report", 5*time.Second, func() {
		// Simulate reporting
	})

	s.Start()

	// Keep the main goroutine alive so the background jobs can run.
	// We will learn better ways to do this in Module 10 (sync.WaitGroup).
	fmt.Println("Scheduler running. Press Ctrl+C to exit.")
	time.Sleep(10 * time.Second) // Let it run for 10 seconds for the demo
	fmt.Println("Shutting down.")
}
