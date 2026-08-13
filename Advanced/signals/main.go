package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Process ID: ", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify the channel on interreupt or terminate signal
	// Ctrl + C sends `SIGINT` (Signal Interrupt).
	// kill -s `SIGINT` <pid> (by default) sends `SIGINT` (Signal Terminate).
	// kill -s SIGTERM <pid> (by default) sends `SIGTERM` (Signal Terminate).
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		fmt.Println("We received signal: ", sig)
		done <- true
	}()

	// Coordinate goroutines using signals
	// This will just loop forever because in main there is for loop. You can test `SIGSTOP` and `SIGCONT` because of this
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal")
				return
			default:
				fmt.Println("We are working")
				time.Sleep(time.Second)
			}
		}
	}()

	// First Example
	// go func() {
	// 	// sig := <-sigs
	// 	// switch sig { // Another sample is just switch with no for loop
	// 	for sig := range sigs {
	// 		switch sig {
	// 		case syscall.SIGINT:
	// 			fmt.Println("\n[Handler] Caught SIGINT (Ctrl + C pressed)")
	// 			// os.Exit(0)
	// 		case syscall.SIGTERM:
	// 			// os.Exit(0)
	// 			fmt.Println("\n[Handler] Caught SIGTERM ('kill' command received)")
	// 		case syscall.SIGHUP:
	// 			fmt.Println(
	// 				"\n[Handler] Caught SIGHUP ('kill' command received)",
	// 			) // This will not do nothing cause `SIGHUP` is not define in notify
	// 		case syscall.SIGUSR1:
	// 			fmt.Println("\n[Handler] Received SIGUSR1 (User defined Signal 1)")
	// 			fmt.Println("User define function is executed")
	// 			continue
	// 		default:
	// 			fmt.Printf("\n[Handler] Caught unexpected signal: %v\n", sig)
	// 		}
	// 		os.Exit(0) // Instead of putting this per case, you can put it here
	// 	}
	//
	// 	// // Evaluate the signal received from the 'sigs' channel to determine the shutdown trigger using if statement
	// 	// if sig == syscall.SIGINT {
	// 	// 	fmt.Println("User manually stopped the program with Ctrl+C")
	// 	// } else if sig == syscall.SIGTERM {
	// 	// 	fmt.Println("Orchestrator/OS asked the program to terminate")
	// 	// }
	//
	// 	// fmt.Println("Received Signal: ", sig) //Use if switch
	// 	fmt.Println("Graceful exit")
	// 	// os.Exit(0) // User if switch
	// }()

	// Simulate some work
	fmt.Println("Working.......")
	for {
		time.Sleep(time.Second)
	}
}

// Sample kill command
// kill -s SIGUSR1 76362
// kill -9 <pid> - Use this if stuck and can't kill the process
