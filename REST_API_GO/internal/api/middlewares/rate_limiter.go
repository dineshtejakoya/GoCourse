package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string]int
	limit    int
	//resetTime is the time duration after which the user can send more requests to the API
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	//start the reset routine
	go rl.resetVisitorCount()
	return rl
}

// resetting the time after given duration
func (rl *rateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		//we are going to empty the visitors value after the sleep duration for accepting new requests
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	fmt.Println("RateLimiter Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RateLimiter Middleware being returned...")
		rl.mu.Lock()
		defer rl.mu.Unlock()

		visitorIP := r.RemoteAddr //You might want to extract the IP in a more sophisticated way
		rl.visitors[visitorIP]++
		fmt.Printf("Visitor Count from %v is %v\n", visitorIP, rl.visitors[visitorIP])

		if rl.visitors[visitorIP] > rl.limit {
			http.Error(w, "Too Many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
		fmt.Println("Rate Limiter ends...")
	})
}
