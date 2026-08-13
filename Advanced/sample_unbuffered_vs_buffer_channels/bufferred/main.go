// Buffered Channel: The Webhook Log Ingestion System
// The Scenario
// You run an e-commerce API that receives thousands of webhooks per second during a holiday sale (e.g., Black Friday). Every time a user does something, you need to log that event to a heavy analytics database (like Elasticsearch or BigQuery).
//
// Why Buffered?
// Writing directly to a database takes time (e.g., 50ms). If you used an unbuffered channel here, your API would block and freeze the user's checkout experience while waiting for the database to finish writing the log.
//
// By using a buffered channel, the API can instantly dump the log event into memory (the buffer) and return an immediate HTTP 200 OK to the customer. A background pool of database workers can then drain that buffer at their own steady pace.

package main

import (
	"fmt"
	"time"
)

type LogEvent struct {
	EventID   string
	Timestamp time.Time
}

// dbWorker runs constantly in the background, saving logs to the DB
func dbWorker(logQueue chan LogEvent, workerID int) {
	for log := range logQueue {
		// Simulate a slow database write
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("[Worker %d] Saved log %s to Database.\n", workerID, log.EventID)
	}
}

func main() {
	// Buffer of 100 means we can handle a sudden burst of 100 incoming requests
	// without slowing down the user's API response time.
	logQueue := make(chan LogEvent, 100)

	// Start 3 background workers to process the queue
	for w := 1; w <= 3; w++ {
		go dbWorker(logQueue, w)
	}

	// Simulate incoming rapid-fire API webhooks
	for i := 1; i <= 10; i++ {
		event := LogEvent{EventID: fmt.Sprintf("EVT-%d", i), Timestamp: time.Now()}

		// This push is lightning fast because the channel has a buffer.
		// The user doesn't have to wait for the DB write to finish.
		logQueue <- event
		fmt.Printf("API Gateway: Received %s, sent 200 OK to user.\n", event.EventID)
	}

	// Hang around to let the background workers finish printing for the demo
	time.Sleep(1 * time.Second)
}
