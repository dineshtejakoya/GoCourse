package advanced

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of person %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		//simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	//start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	//send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}

	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}
}
