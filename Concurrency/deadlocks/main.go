package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)

		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")

		mu1.Unlock()
		mu2.Unlock()
	}()

	go func() {
		mu1.Lock() // Swap this to cause a deadlock
		fmt.Println("Goroutine 2 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock() // Swap this to cause a deadlock

		fmt.Println("Goroutine 2 locked mu2")

		mu2.Unlock()
		mu1.Unlock()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
	// select {} // Wait forver to see the goroutines to finish
}
