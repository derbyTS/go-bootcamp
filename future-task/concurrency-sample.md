# Go Technical Interview Study Guide: Rate-Limited API Aggregator

In technical interviews for Go roles, interviewers rarely ask you to write a basic `go func()` loop. Instead, they test whether you know how to build **resilient, rate-limited, and production-safe concurrent systems**.

Here is a real-world scenario you will almost certainly encounter during live coding assessments or take-home projects.

---

## The Scenario: "The Rate-Limited API Aggregator"

> **Interview Prompt:**  
> _"We have 100 customer IDs. We need to fetch their billing details from a third-party payment gateway API. However, the gateway has two strict limits:_
>
> 1. _It will block our IP if we make more than **5 concurrent requests**._
> 2. _Individual request timeouts: If any single API call takes longer than **1 second**, cancel that request, record an error, and continue with the rest._
>
> _Write a Go program that processes all 100 IDs as fast as possible within these constraints, collecting all results and errors safely."_

---

## Why Interviewers Pick This Scenario

This single task tests whether you understand the **4 core pillars of production Go concurrency**:

- **Worker Pools:** Bounding concurrency so you don't blow up CPU, RAM, or external API limits.
- **Context & Timeouts (`context.Context`):** Preventing hanging goroutines when network calls stall.
- **Channel Management:** Safely passing jobs in and results out without deadlocks or goroutine leaks.
- **Error Handling:** Aggregating failures without crashing the whole process.

---

## Idiomatic Go Solution

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents input data
type Job struct {
	UserID int
}

// Result holds the outcome of a job execution
type Result struct {
	UserID int
	Data   string
	Err    error
}

// Simulated third-party API call
func fetchBillingData(ctx context.Context, userID int) (string, error) {
	// Simulate unpredictable network latency (between 200ms and 1500ms)
	latency := time.Duration(200+rand.Intn(1300)) * time.Millisecond

	select {
	case <-time.After(latency):
		// API call succeeded within time
		return fmt.Sprintf("BillingInfo_User_%d", userID), nil

	case <-ctx.Done():
		// Context timed out or was cancelled before API finished!
		return "", ctx.Err()
	}
}

// Worker function consuming jobs and emitting results
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Enforce a strict 1-second timeout per individual request
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

		data, err := fetchBillingData(ctx, job.UserID)
		cancel() // Clean up context resources immediately after call

		results <- Result{
			UserID: job.UserID,
			Data:   data,
			Err:    err,
		}
	}
}

func main() {
	totalUsers := 20
	maxConcurrency := 5 // Hard limit: max 5 requests at once

	jobs := make(chan Job, totalUsers)
	results := make(chan Result, totalUsers)

	var wg sync.WaitGroup

	// 1. Start bounded worker pool (exactly 5 workers)
	for w := 1; w <= maxConcurrency; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Feed jobs into queue
	for i := 1; i <= totalUsers; i++ {
		jobs <- Job{UserID: i}
	}
	close(jobs) // Signal workers no more work is coming

	// 3. Close results channel once all workers finish in background
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Process results as they come in
	successCount := 0
	failureCount := 0

	for res := range results {
		if res.Err != nil {
			if errors.Is(res.Err, context.DeadlineExceeded) {
				fmt.Printf("[TIMEOUT] User %d exceeded 1s limit\n", res.UserID)
			} else {
				fmt.Printf("[ERROR]   User %d failed: %v\n", res.UserID, res.Err)
			}
			failureCount++
			continue
		}

		fmt.Printf("[SUCCESS] User %d -> %s\n", res.UserID, res.Data)
		successCount++
	}

	fmt.Printf("\n--- Summary: Successes: %d | Failures/Timeouts: %d ---\n", successCount, failureCount)
}
```

## What to Mention in the Interview (The "Senior Dev" Moves)

If you explain **why** you wrote the code this way, you instantly stand out:

- **Why `close(jobs)` before `wg.Wait()`?**  
  If you call `wg.Wait()` on the main thread before closing `jobs`, `main()` blocks forever, and workers never see the closed channel—causing a **deadlock**.

- **Why defer `cancel()` inside the worker?**  
  Calling `cancel()` releases timer resources associated with `context.WithTimeout`. Forgetting this causes memory/timer leaks inside the Go runtime.

- **Why use a channel for results instead of a shared Slice/Map?**  
  If multiple goroutines write to a shared slice without a `sync.Mutex`, your app suffers from **data races** (`go run -race`). Channels pass data ownership safely.

- **How do you guarantee exactly 5 concurrent requests?**  
  Instead of spawning 100 goroutines, you spawn **only 5 worker goroutines**. They pull from the channel as they become free, guaranteeing the downstream API never sees more than 5 concurrent connections.
