package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

// func NewRatelimiter(limit int, window time.Duration) *RateLimiter {
// 	return &RateLimiter{
// 		limit:  limit,
// 		window: window,
// 	}
// }

// Suggested by llm
func NewRatelimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:     limit,
		window:    window,
		resetTime: time.Now().Add(window), // This is added
	}
}

// Request to LLM to show what is capacity remaining so it add another value to return instead of bool only
func (rl *RateLimiter) Allow() (bool, int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true, rl.limit - rl.count
	}
	return false, 0
}

func main() {
	var wg sync.WaitGroup
	rateLimiter := NewRatelimiter(5, 2*time.Second)

	for i := range 20 {
		wg.Add(1)
		go func() {
		}()
		allowed, remaining := rateLimiter.Allow()
		if allowed {
			fmt.Printf("Request %d: Allowed | Capacity remaining in window: %d\n", i, remaining)
		} else {
			fmt.Printf("Request %d: Denied  | Capacity remaining in window: %d\n", i, remaining)
		}
		// time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}
