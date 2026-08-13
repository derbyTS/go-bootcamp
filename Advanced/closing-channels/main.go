package main

import (
	"fmt"
	// "time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}

func producer(ch chan<- int) {
	defer close(ch)
	for i := range 5 {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int) {
	defer close(out)
	for val := range in {
		out <- val
	}
}

// Range over closed channel
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for i := range 5 {
// 			ch <- i
// 		}
// 	}()
// 	fmt.Println("Before Receiving")
// 	time.Sleep(3 * time.Second)
// 	for value := range ch {
// 		fmt.Println(value)
// 	}
// }

// Receiving from close channel
// func main() {
// 	ch := make(chan int, 5)
//
// 	close(ch)
//
// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel is close")
// 	} else {
// 		fmt.Println(val)
// 	}
// }

// Simple closing channel
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for i := range 5 {
// 			ch <- i
// 		}
// 	}()
//
// 	for value := range ch {
// 		fmt.Println(value)
// 	}
// }
