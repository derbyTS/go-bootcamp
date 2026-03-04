package main

import (
	"fmt"
	"time"
	// "time"
)

// func main() {
// 	//  ================ Blocking on receive only because buffer is empty
// 	//  ================ The magic here is it pause and wait the channel to be filled before it continue
// 	ch := make(chan int, 2)
//
// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println("End")
// }

func main() {
	// ================ Blocking on send only if buffer is full
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine: Removing 1 from buffer...")
		fmt.Println(<-ch)
	}()
	fmt.Println("Blocking Starts")
	ch <- 3
	fmt.Println("Blocking Ends")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("Buffered Channels")
}
