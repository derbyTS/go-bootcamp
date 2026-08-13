package main

import (
	"fmt"
	"time"
)

// type Ratelimiter struct {
// 	tokens     chan struct{}
// 	refillTime time.Duration
// }
//
// func NewRateLimiter(rateLimit int, refillTime time.Duration) *Ratelimiter {
// 	r1 := &Ratelimiter{
// 		tokens:     make(chan struct{}, rateLimit),
// 		refillTime: refillTime,
// 	}
//
// 	for range rateLimit {
// 		r1.tokens <- struct{}{}
// 	}
//
// 	go r1.startRefill()
//
// 	return r1
// }
//
// func (rl *Ratelimiter) allow() bool {
// 	select {
// 	case <-rl.tokens:
// 		return true
// 	default:
// 		return false
// 	}
// }
//
// func (rl *Ratelimiter) startRefill() {
// 	ticker := time.NewTicker(rl.refillTime)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			select {
// 			case rl.tokens <- struct{}{}:
// 			default:
// 			}
// 		}
// 	}
// }
//
// func main() {
// 	rateLimiter := NewRateLimiter(5, time.Second)
//
// 	for i := range 20 {
// 		if rateLimiter.allow() {
// 			fmt.Printf("Request %d allowed\n", i)
// 		} else {
// 			fmt.Printf("Request %d denied\n", i)
// 		}
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

// ***************** Gemini Recommendation *****************

type Ratelimiter struct {
	tokens chan struct{}
	stop   chan struct{} // Signal channel to shut down the refill goroutine
}

// NewRateLimiter creates a bucket of size 'capacity' that adds 1 token every 'fillInterval'
func NewRateLimiter(capacity int, fillInterval time.Duration) *Ratelimiter {
	rl := &Ratelimiter{
		tokens: make(chan struct{}, capacity),
		stop:   make(chan struct{}),
	}

	// Fill bucket initially
	for range capacity {
		rl.tokens <- struct{}{}
	}

	go rl.startRefill(fillInterval)

	return rl
}

func (rl *Ratelimiter) Allow() (bool, int) {
	select {
	case <-rl.tokens:
		// len(rl.tokens) now holds the count AFTER taking 1 token
		return true, len(rl.tokens)
	default:
		// Bucket is empty, 0 tokens left
		return false, 0
	}
}

func (rl *Ratelimiter) startRefill(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default: // Bucket full, drop token
			}
		case <-rl.stop: // Graceful cleanup
			return
		}
	}
}

// Stop shuts down the refill goroutine cleanly
func (rl *Ratelimiter) Stop() {
	close(rl.stop)
}

func main() {
	// Capacity of 5, adds 1 token every 200ms (5 tokens/sec rate)
	rateLimiter := NewRateLimiter(5, 200*time.Millisecond)
	defer rateLimiter.Stop()

	for i := range 20 {
		allowed, remaining := rateLimiter.Allow()
		if allowed {
			fmt.Printf("Request %d allowed | Tokens left: %d\n", i, remaining)
		} else {
			fmt.Printf("Request %d denied  | Tokens left: %d\n", i, remaining)
		}
		// time.Sleep(200 * time.Millisecond)
		time.Sleep(100 * time.Millisecond) // <-- Changed from 200ms to 100ms will show bucket drain
	}
}
