package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("CPU num: %d\n", runtime.NumCPU())
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()

	fmt.Printf("Final value of counter: %d\n", counter)
}

// ************* Sample with struct
// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }
//
// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }
//
// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }
//
// func main() {
// 	fmt.Printf("CPU num: %d\n", runtime.NumCPU())
// 	runtime.GOMAXPROCS(
// 		4,
// 	) // Restrict Go to only use 4 CPU cores, even though you have 11
// 	fmt.Printf("CPU num: %d\n", runtime.NumCPU()) // This will still print 11 (hardware physics)
// 	var wg sync.WaitGroup
// 	counter := &counter{}
// 	numGoroutines := 10
//
// 	// wg.Add(numGoroutines)
// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 				// counter.count++ //Uncomment this and comment the above line too see the difference
// 			}
// 		}()
// 	}
//
// 	wg.Wait()
// 	fmt.Printf("Final counter value: %d\n", counter.getValue())
// }
