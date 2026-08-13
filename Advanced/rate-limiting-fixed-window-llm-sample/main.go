package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// clientWindow tracks the request count and reset boundary for a single client IP
type clientWindow struct {
	count     int
	resetTime time.Time
}

// FixedWindowRateLimiter manages counters for all connected clients
type FixedWindowRateLimiter struct {
	mu         sync.Mutex
	clients    map[string]*clientWindow
	limit      int
	windowSize time.Duration
}

func NewFixedWindowRateLimiter(limit int, windowSize time.Duration) *FixedWindowRateLimiter {
	limiter := &FixedWindowRateLimiter{
		clients:    make(map[string]*clientWindow),
		limit:      limit,
		windowSize: windowSize,
	}

	// Periodically clean up stale client entries to prevent memory leaks
	go limiter.cleanupLoop(5 * time.Minute)

	return limiter
}

func (rl *FixedWindowRateLimiter) Allow(ip string) (bool, int, time.Duration) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	client, exists := rl.clients[ip]

	// If new client or window expired, reset window for this IP
	if !exists || now.After(client.resetTime) {
		rl.clients[ip] = &clientWindow{
			count:     1,
			resetTime: now.Add(rl.windowSize),
		}
		return true, rl.limit - 1, rl.windowSize
	}

	// If limit reached, return denied + time remaining in current window
	if client.count >= rl.limit {
		timeToReset := client.resetTime.Sub(now)
		return false, 0, timeToReset
	}

	client.count++
	return true, rl.limit - client.count, client.resetTime.Sub(now)
}

// Routine cleanup prevents the map from growing endlessly with old IPs
func (rl *FixedWindowRateLimiter) cleanupLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		rl.mu.Lock()
		now := time.Now()
		for ip, client := range rl.clients {
			if now.After(client.resetTime) {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// ----------------------------------------------------------------------------
// REAL-WORLD HTTP MIDDLEWARE & HEADERS
// ----------------------------------------------------------------------------

func FixedWindowMiddleware(
	limiter *FixedWindowRateLimiter,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract client IP address
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		allowed, remaining, resetDuration := limiter.Allow(ip)

		// Set standard RFC 6585 and IETF RateLimit HTTP Headers
		w.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.limit))
		w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(resetDuration).Unix()))

		if !allowed {
			// Tell the client exactly how many seconds to wait before trying again
			retryAfterSeconds := int(resetDuration.Seconds()) + 1
			w.Header().Set("Retry-After", fmt.Sprintf("%d", retryAfterSeconds))

			http.Error(
				w,
				fmt.Sprintf("429 Too Many Requests. Try again in %d seconds.", retryAfterSeconds),
				http.StatusTooManyRequests,
			)
			return
		}

		next(w, r)
	}
}

func main() {
	// Limit: 5 requests per 10-second fixed window per IP address
	limiter := NewFixedWindowRateLimiter(5, 10*time.Second)

	// Simulated authentication endpoint
	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "200 OK: Authentication attempt processed.")
	}

	http.HandleFunc("/api/login", FixedWindowMiddleware(limiter, loginHandler))

	fmt.Println("Auth server running on :8080...")
	http.ListenAndServe(":8080", nil)
}

/*
>go run main.go
>
for i in {1..20}; do
  curl -i http://localhost:8080/api/login
  echo ""
done
*/
