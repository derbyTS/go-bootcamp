package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTimer(time.Second)
	t2 := time.NewTimer(time.Second)

	for {
		time.Sleep(2 * time.Second)
		select {
		case <-t1.C:
			fmt.Println("T1 Done")
		case <-t2.C:
			fmt.Println("T2 Done")
		default:
			fmt.Println("End now")
			return
		}
	}
}

// **********Scheduling delayed operation
// func main() {
// 	timer := time.NewTimer(2 * time.Second) // always rember that timer is non-blocking
// 	defer timer.Stop()
//
// 	go func() {
// 		<-timer.C // it didn't block the main thread but the goroutine
// 		fmt.Println("Delayed operation execute")
// 	}()
//
// 	fmt.Println(
// 		"Waiting ....",
// 	) // execute immediately even the goroutine is block by <-timer.c it dind't block the main thread
// 	time.Sleep(3 * time.Second) // this block the main thread
// 	fmt.Println("End of program")
// }

// func longRunning() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }
//
// // Example by LLM
// func main() {
// 	timer := time.NewTimer(5 * time.Second)
// 	done := make(chan bool)
//
// 	go func() {
// 		longRunning()
// 		done <- true
// 	}()
//
// 	select {
// 	case <-timer.C:
// 		fmt.Println("Operation timeout")
// 	case <-done:
// 		// According to LLM Stop the timer immediately to free up memory!
// 		if !timer.Stop() {
// 			<-timer.C // Drain the channel if it already expired
// 		}
// 		fmt.Println("Operation completed")
// 	}
// }

// func main() {
// 	timeout := time.After(2 * time.Second)
// 	done := make(chan bool)
//
// 	go func() {
// 		longRunning()
// 		done <- true
// 	}()
//
// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation timeout")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}
// }

// func main() {
// 	// time.Sleep(time.Second)
// 	fmt.Println("Starting the app")
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for timer.C")
//
// 	stopped := timer.Stop()
//
// 	if stopped {
// 		fmt.Println("Timer Stopped")
// 	} else {
// 		fmt.Println("Timer Continue")
// 	}
// 	timer.Reset(time.Second)
// 	<-timer.C // Blocking in nature
// 	fmt.Println("Timer Expired")
// }
