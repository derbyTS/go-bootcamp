package main

import "fmt"

func main() {
	ch := make(chan int)
	produceData(ch)
	consumeData(ch)
}

func produceData(ch chan<- int) {
	go func(ch chan<- int) {
		defer close(ch)
		for i := range 5 {
			ch <- i
		}
	}(ch)
}

func consumeData(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
