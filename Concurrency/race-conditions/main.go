package main

import (
	"fmt"
	"sync"
)

// ************* Sample with struct
type counter struct {
	mu    sync.Mutex
	count int
}

// ====== Commenting a mutex here will case a race condition
func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// fmt.Printf("CPU num: %d\n", runtime.NumCPU())
	// runtime.GOMAXPROCS(
	// 	4,
	// ) // Restrict Go to only use 4 CPU cores, even though you have 11
	// fmt.Printf("CPU num: %d\n", runtime.NumCPU()) // This will still print 11 (hardware physics)
	var wg sync.WaitGroup
	counter := &counter{}
	numGoroutines := 10

	// wg.Add(numGoroutines)
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++ // Uncomment this and comment the above line too see the difference
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())
}
