package advanced

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	for range rateLimit {
		rl.tokens <- struct{}{}
	}

	go rl.startRefill()

	return rl

}

// refill the tokens as they are getting used
func (rl *RateLimiter) startRefill() {
	//ticker will create a new token after a certain time
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request Denied")
		}
		//for 300 millisecond o/p is different
		time.Sleep(200 * time.Millisecond)
	}
}

//1 0 ms  First Request Allowed   5 tokens left
//2 200 ms Second Request Allowed 4
//3 400 ms Third Request Allowed  3
//4 600 ms Fourth Request Allowed 2
//5 800 ms Fifth Request Allowed  1
//6 1000 ms Sixth Request Allowed (not 0)  1 tokens left, the startRefill function executes and adds one more token at this point
//7 1200 ms Seventh  Request Allowed
//8 1400 ms Eighth  Request Allowed
//9 1600 ms Ninth Request Allowed
