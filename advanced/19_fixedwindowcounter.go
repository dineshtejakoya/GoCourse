package advanced

import (
	"fmt"
	"sync"
	"time"
)

/*
Divides time into equal sized windows
it allows a fixed number of request per window, if the request count exceeds the limit per window subsequent rquest are denied until the window reset
simple and easy to implement can lead to burst of traffic at the edge of windows
*/

// track number of requests,count of requests(limit),duration of rate limiting window and the time that window should reset
type RateLimiter struct {
	//we want to protect our data when we are modifying it, locking and unlocking data when we are modifying
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	//5,2*time.second
	//3,1*time.second
	rateLimiter := NewRateLimiter(5, 2*time.Second)

	for range 10 {
		if rateLimiter.Allow() {
			fmt.Println("Request Allowed.")
		} else {
			fmt.Println("Request Denied.")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
