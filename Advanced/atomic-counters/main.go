package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup
	numGoRoutines := 10
	counter := &AtomicCounter{}

	value := 1
	for range numGoRoutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				value++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final value: %d\n", counter.getValue())
	fmt.Printf("Final value: %d\n", value)
}
