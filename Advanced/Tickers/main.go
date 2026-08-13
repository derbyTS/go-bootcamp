package main

import (
	"fmt"
	"time"
)

// Sample multiple ticker by LLM
func main() {
	// 1. Initialize multiple tickers with different intervals
	heartbeatTicker := time.NewTicker(2 * time.Second)
	metricsTicker := time.NewTicker(5 * time.Second)
	cleanupTicker := time.NewTicker(10 * time.Second)

	// Ensure they are cleaned up when the function exits
	defer heartbeatTicker.Stop()
	defer metricsTicker.Stop()
	defer cleanupTicker.Stop()

	// 2. A 'done' channel to gracefully shut down the loop
	done := make(chan bool)

	// In a real app, you might trigger 'done' via a SIGTERM signal
	go func() {
		time.Sleep(25 * time.Second)
		done <- true
	}()

	fmt.Println("Monitoring service started...")

	// 3. The multiplexing loop
	for {
		select {
		case <-done:
			fmt.Println("Shutting down monitor...")
			return

		case t := <-heartbeatTicker.C:
			fmt.Printf("[%v] Heartbeat: Service is healthy\n", t.Format("15:04:05"))

		case t := <-metricsTicker.C:
			fmt.Printf("[%v] Metrics: Logged 42 requests in the last 5s\n", t.Format("15:04:05"))

		case t := <-cleanupTicker.C:
			fmt.Printf("[%v] Cleanup: Removed 3 stale temp files\n", t.Format("15:04:05"))
		}
	}
}

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	stop := time.After(5 * time.Second)
//
// 	defer ticker.Stop()
//
// 	for {
// 		select {
// 		case tick := <-ticker.C:
// 			fmt.Println("Tick at: ", tick)
// 		case <-stop:
// 			fmt.Println("Stop ticker")
// 			return
// 		}
// 	}
// }

// ***** scheduling logging, periodi task and polling for updates
// func periodicTask() {
// 	fmt.Println("Performing period task at: ", time.Now())
// }
//
// func main() {
// 	ticker := time.NewTicker(time.Second)
//
// 	defer ticker.Stop()
//
// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
//
// 	defer ticker.Stop()
//
// 	i := 1
// 	for range ticker.C {
// 		i *= 2
// 		fmt.Println(i)
// 		if i > 1000 {
// 			return
// 		}
// 	}
// }
