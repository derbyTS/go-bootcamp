package main

import (
	"fmt"
	"sync"
	"time"
)

// type LeakyBucket struct {
// 	capacity int
// 	leakRate time.Duration
// 	tokens   int
// 	lastLeak time.Time
// 	mu       sync.Mutex
// }
//
// func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
// 	return &LeakyBucket{
// 		capacity: capacity,
// 		leakRate: leakRate,
// 		tokens:   capacity,
// 		lastLeak: time.Now(),
// 	}
// }
//
// func (lb *LeakyBucket) Allow() bool {
// 	lb.mu.Lock()
// 	defer lb.mu.Unlock()
//
// 	now := time.Now()
// 	elapsedTime := now.Sub(lb.lastLeak)
// 	tokensToAdd := int(elapsedTime / lb.leakRate)
// 	lb.tokens += tokensToAdd
//
// 	if lb.tokens > lb.capacity {
// 		lb.tokens = lb.capacity
// 	}
//
// 	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
//
// 	if lb.tokens > 0 {
// 		lb.tokens--
// 		return true
// 	}
//
// 	return false
// }
//
// func main() {
// 	leakyBucketInstance := NewLeakyBucket(5, 500*time.Millisecond)
//
// 	for range 10 {
// 		if leakyBucketInstance.Allow() {
// 			fmt.Println("Request Accepted")
// 		} else {
// 			fmt.Println("Request Rejected")
// 		}
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

// LLM said this
// Oh boy, this Udemy code is actually a Token Bucket algorithm disguised under the name LeakyBucket!
//
// Your instructor wrote code that uses tokens, refills them over elapsed time up to a capacity, and decrements them when a request comes in. That is literally the textbook definition of a Token Bucket algorithm.
//
// Here is what is fundamentally wrong with this Udemy example:
//
// Why this Udemy sample is completely incorrect
//
// 1. It uses Tokens instead of a Queue/Water Level
//
// 		Real Leaky Bucket: Requests enter a queue/bucket (like water), and leak out at a constant, fixed rate.
//
// 		This Code: Has a pool of tokens that gets refilled over time. That's Token Bucket logic.
//
// 2. It ALLOWS Bursts (Which Leaky Bucket is explicitly designed to prevent!)
//
// 		In this Udemy code, if you start with capacity: 5, all 5 requests will be allowed instantly in a single burst.
//
// 		A real Leaky Bucket never allows bursts. Even if 5 requests arrive at the exact same microsecond, it forces them into a queue and processes them one-by-one at a strictly enforced interval (e.g., 1 request every 500ms).
//
// 3. It's the Token Bucket implementation backwards
//
// 		The instructor starts with tokens = capacity, adds tokens lazily when time passes, and decrements tokens on requests. That is identical to the lazy-refill Token Bucket algorithm we discussed at the very beginning of our conversation.
//
// The Verdict
// Your instructor has completely confused Token Bucket and Leaky Bucket.
//
// They wrote a Token Bucket algorithm, named the struct LeakyBucket, and called the refill mechanism leakRate. You were right to be confused earlier—your course is teaching Token Bucket twice under two different names!

// ********************** Suggest by llm **********************

// type LeakyBucket struct {
// 	queue    chan func() // The bucket holding queued requests
// 	leakRate time.Duration
// 	stop     chan struct{}
// 	wg       sync.WaitGroup
// }
//
// func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
// 	lb := &LeakyBucket{
// 		queue:    make(chan func(), capacity), // Capacity = max queue depth
// 		leakRate: leakRate,
// 		stop:     make(chan struct{}),
// 	}
//
// 	// Start the background routine that "leaks" (processes) 1 request every leakRate
// 	lb.wg.Add(1)
// 	go lb.startLeaking()
//
// 	return lb
// }
//
// // Submit tries to add a request to the bucket.
// // Returns true if queued, false if the bucket is full (overflow).
// func (lb *LeakyBucket) Submit(task func()) bool {
// 	select {
// 	case lb.queue <- task:
// 		return true // Request successfully queued
// 	default:
// 		return false // Bucket is full! Overflow / dropped
// 	}
// }
//
// func (lb *LeakyBucket) startLeaking() {
// 	defer lb.wg.Done()
// 	ticker := time.NewTicker(lb.leakRate)
// 	defer ticker.Stop()
//
// 	for {
// 		select {
// 		case <-ticker.C:
// 			// Time to leak 1 request from the queue
// 			select {
// 			case task := <-lb.queue:
// 				task() // Execute the request at a steady rate
// 			default:
// 				// Bucket is empty, nothing to leak right now
// 			}
// 		case <-lb.stop:
// 			return
// 		}
// 	}
// }
//
// func (lb *LeakyBucket) Stop() {
// 	close(lb.stop)
// 	lb.wg.Wait()
// }
//
// func main() {
// 	// Bucket capacity = 5 queued requests max
// 	// Leak rate = 1 request processed every 500ms
// 	lb := NewLeakyBucket(5, 500*time.Millisecond)
//
// 	// Fire 10 rapid requests (every 100ms)
// 	for i := range 10 {
// 		reqID := i
// 		submitted := lb.Submit(func() {
// 			fmt.Printf("Request %d EXECUTED at %s\n", reqID, time.Now().Format("15:04:05.000"))
// 		})
//
// 		if submitted {
// 			fmt.Printf("Request %d QUEUED   at %s\n", reqID, time.Now().Format("15:04:05.000"))
// 		} else {
// 			fmt.Printf(
// 				"Request %d DROPPED  at %s (Bucket Full!)\n",
// 				reqID,
// 				time.Now().Format("15:04:05.000"),
// 			)
// 		}
//
// 		time.Sleep(100 * time.Millisecond)
// 	}
//
// 	// Wait long enough for queued requests to finish leaking
// 	time.Sleep(3 * time.Second)
// 	lb.Stop()
// }

// *********** Another Sample with almost similar with Udemy
// LeakyBucket implements a rate limiter using the Leaky Bucket algorithm.
type LeakyBucket struct {
	mu         sync.Mutex
	capacity   float64   // Maximum burst capacity of the bucket
	leakRate   float64   // How many requests leak out per second
	waterLevel float64   // Current water level (buffered requests)
	lastLeak   time.Time // Last time the bucket leaked water
}

// NewLeakyBucket creates a new bucket with a defined capacity and leak rate per second.
func NewLeakyBucket(capacity float64, leakRate float64) *LeakyBucket {
	return &LeakyBucket{
		capacity:   capacity,
		leakRate:   leakRate,
		waterLevel: 0,
		lastLeak:   time.Now(),
	}
}

// Allow checks if a request can be processed. Returns true if allowed, false if dropped.
func (b *LeakyBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastLeak).Seconds()
	b.lastLeak = now

	// Calculate how much water leaked since the last request
	leakedAmount := elapsed * b.leakRate
	b.waterLevel -= leakedAmount
	if b.waterLevel < 0 {
		b.waterLevel = 0
	}

	// Check if the bucket can hold one more unit of water (the new request)
	if b.waterLevel+1 <= b.capacity {
		b.waterLevel += 1
		return true // Request accepted
	}

	return false // Bucket overflow, request dropped
}

func main() {
	// Create a bucket with a capacity of 3 requests, leaking at 2 requests per second
	bucket := NewLeakyBucket(3, 2)

	// Simulate an immediate burst of 5 requests
	fmt.Println("--- Simulating immediate burst of 5 requests ---")
	for i := 1; i <= 5; i++ {
		allowed := bucket.Allow()
		fmt.Printf(
			"Request %d: Allowed = %v (Current Water Level: %.2f)\n",
			i,
			allowed,
			bucket.waterLevel,
		)
	}

	// Wait for 1 second to allow the bucket to leak/empty out
	fmt.Println("\n--- Waiting 1 second for the bucket to leak ---")
	time.Sleep(1 * time.Second)

	// Send more requests after the leak window
	fmt.Println("\n--- Sending 3 more requests after waiting ---")
	for i := 6; i <= 8; i++ {
		allowed := bucket.Allow()
		fmt.Printf(
			"Request %d: Allowed = %v (Current Water Level: %.2f)\n",
			i,
			allowed,
			bucket.waterLevel,
		)
	}
}
