package advanced

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	//leakrate is how often tokens are added to bucket
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	//if we are getting concurrent request there won't be incorrect counting
	mu sync.Mutex
}

// create a function that will create new instance of leakybucket
func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

// allow method associated with leakybucket struct
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	//calculate the time since the last leak
	//now.Sub() it'll subtract the current time that we are going to pass it as the argument
	elapsedTime := now.Sub(lb.lastLeak)

	//1/0.5 =>2 tokens to add
	tokensToAdd := int(elapsedTime / lb.leakRate) //(0.2/0.5) result is 0.4 int value is 0
	//3+4=7 however the capacity is 5
	lb.tokens = lb.tokens + tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	//update the lastleaktime
	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	//lb.lastLeak = lb.lastLeak.Add(elapsedTime)
	/*
		elapsedtime = 1.3
		initial lastleak=0
		tokenstoadd = int(1.3/.5)= 2.6 => 2 tokens
		lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) + lb.leakRate)
		= 0 + (2*0.5) => 1 second
		lb.lastLeak = lb.lastLeak.Add(elapsedTime) //wrong approoach

	*/

	fmt.Printf("Tokens added %d,Tokens subtracted %d,Total Tokens : %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("Last leak time: %v\n", lb.lastLeak)

	//if lb.tokens >0 we need to allow the function and consume the tokens
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	return false

}

func main() {
	leakyBucketInst := NewLeakyBucket(5, 500*time.Millisecond)

	for range 10 {
		if leakyBucketInst.Allow() {
			fmt.Println("Current Time:", time.Now())
			fmt.Println("Request Accepted")
		} else {
			fmt.Println("Current Time:", time.Now())
			fmt.Println("Request Denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
