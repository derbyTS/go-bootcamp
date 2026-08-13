package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// LeakyBucketShaper manages queueing and smooth processing rates per route/service
type LeakyBucketShaper struct {
	queue    chan chan struct{} // Channel of request signal channels
	leakRate time.Duration
	stop     chan struct{}
	wg       sync.WaitGroup
}

func NewLeakyBucketShaper(capacity int, leakRate time.Duration) *LeakyBucketShaper {
	lb := &LeakyBucketShaper{
		queue:    make(chan chan struct{}, capacity), // Capacity = max waiting requests in queue
		leakRate: leakRate,
		stop:     make(chan struct{}),
	}

	lb.wg.Add(1)
	go lb.startLeaking()

	return lb
}

// startLeaking drains 1 queued HTTP request every 'leakRate'
func (lb *LeakyBucketShaper) startLeaking() {
	defer lb.wg.Done()
	ticker := time.NewTicker(lb.leakRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case reqSignal := <-lb.queue:
				reqSignal <- struct{}{} // Notify the waiting HTTP request to proceed
			default:
				// Queue is empty, nothing to leak
			}
		case <-lb.stop:
			return
		}
	}
}

func (lb *LeakyBucketShaper) Stop() {
	close(lb.stop)
	lb.wg.Wait()
}

// ----------------------------------------------------------------------------
// REAL-WORLD HTTP MIDDLEWARE
// ----------------------------------------------------------------------------

func LeakyBucketMiddleware(shaper *LeakyBucketShaper, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Each incoming request gets its own signal channel to wait on
		readySignal := make(chan struct{})

		select {
		case shaper.queue <- readySignal:
			// Successfully queued! Wait until the background leaky ticker signals us
			<-readySignal
			next(w, r)
		default:
			// Queue is at max capacity! Reject immediately without waiting
			w.Header().Set("Retry-After", "2")
			http.Error(
				w,
				"429 Too Many Requests - Traffic queue full, slow down!",
				http.StatusTooManyRequests,
			)
		}
	}
}

func main() {
	// Queue holds max 5 waiting requests.
	// Outflow (leak rate) is strictly throttled to 1 request every 500ms.
	shaper := NewLeakyBucketShaper(5, 500*time.Millisecond)
	defer shaper.Stop()

	// Target API Endpoint (e.g., generating a heavy PDF or writing to a DB)
	heavyJobHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "200 OK: Processed at %s\n", time.Now().Format("15:04:05.000"))
	}

	http.HandleFunc("/api/heavy-task", LeakyBucketMiddleware(shaper, heavyJobHandler))

	fmt.Println("Traffic-shaped server running on :8080...")
	http.ListenAndServe(":8080", nil)
}

// ******* Normal Run *******
// for i in {1..10}; do
//   curl -i http://localhost:8080/api/heavy-task
//   echo ""
// done
// ******* Overwhelm Run *******
// for i in {1..15}; do
//   curl -i http://localhost:8080/api/heavy-task &
// done
// wait
