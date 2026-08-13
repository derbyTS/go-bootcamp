package main

import (
	"fmt"
	"time"
)

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan TicketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- TicketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}

	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Thank you %d\n", <-ticketResults)
	}
}

type TicketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func ticketProcessor(requests <-chan TicketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf(
			"Processing %d tickets of personID %d with total cost of %d\n",
			req.numTickets,
			req.personID,
			req.cost,
		)
		time.Sleep(time.Second)
		results <- req.personID
	}
}

//*************** Basic Worker Pool
// func main() {
// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)
//
// 	// Create worker
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}
//
// 	for i := range numJobs {
// 		tasks <- i
// 	}
//
// 	close(tasks)
//
// 	// Collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result: ", result)
// 	}
// }
//
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }
