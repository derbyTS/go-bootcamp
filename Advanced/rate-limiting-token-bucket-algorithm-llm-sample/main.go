package main

import (
	"fmt"
	"net/http"
	"time"
)

type Ratelimiter struct {
	tokens chan struct{}
	stop   chan struct{}
}

func NewRateLimiter(capacity int, fillInterval time.Duration) *Ratelimiter {
	rl := &Ratelimiter{
		tokens: make(chan struct{}, capacity),
		stop:   make(chan struct{}),
	}

	for range capacity {
		rl.tokens <- struct{}{}
	}

	go rl.startRefill(fillInterval)

	return rl
}

func (rl *Ratelimiter) Allow() (bool, int) {
	select {
	case <-rl.tokens:
		return true, len(rl.tokens)
	default:
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
			default:
			}
		case <-rl.stop:
			return
		}
	}
}

func (rl *Ratelimiter) Stop() {
	close(rl.stop)
}

// ----------------------------------------------------------------------------
// REAL-WORLD HTTP MIDDLEWARE
// ----------------------------------------------------------------------------

func RateLimitMiddleware(limiter *Ratelimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowed, remaining := limiter.Allow()

		// Include rate limit info in the standard HTTP response headers
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))

		if !allowed {
			// 429 Too Many Requests is the standard HTTP status for rate limiting
			http.Error(w, "429 Too Many Requests - Slow down!", http.StatusTooManyRequests)
			return
		}

		// Token consumed successfully! Proceed to the actual API handler
		next(w, r)
	}
}

func main() {
	// Allow a burst of 5 requests max, replenishing 1 token every 200ms
	limiter := NewRateLimiter(5, 200*time.Millisecond)
	defer limiter.Stop()

	// Target API Endpoint (e.g., getting user profile or processing an order)
	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		// Simulate doing real work (database lookup, external API call, etc.)
		time.Sleep(50 * time.Millisecond)
		fmt.Fprintln(w, "Request processed successfully!")
	}

	// Protect the API endpoint with our rate limiter middleware
	http.HandleFunc("/api/data", RateLimitMiddleware(limiter, apiHandler))

	fmt.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}

/*
> go run main.go
>
for i in {1..10}; do
  curl -i http://localhost:8080/api/data
  echo ""
done

*/
