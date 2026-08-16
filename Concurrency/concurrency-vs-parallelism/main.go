package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Worker accepts a channel of slice chunks (chan []int)
func worker(id int, jobs <-chan []int, wg *sync.WaitGroup) {
	defer wg.Done()

	processed := 0

	// Each 'batch' is a slice of 1,000 numbers
	for batch := range jobs {
		processed += len(batch)
	}

	fmt.Printf("Worker ID %d finished and dynamically processed %d items\n", id, processed)
}

func main() {
	numWorkers := 4
	runtime.GOMAXPROCS(numWorkers)

	totalJobs := 100_000_000
	batchSize := 1_000

	// Buffer holds 100 BATCHES (not 100 single ints)
	jobs := make(chan []int, 100)
	var wg sync.WaitGroup

	startTime := time.Now()

	// 1. Launch worker goroutines
	for i := range numWorkers {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// 2. Feed jobs in BATCHES of 1,000
	batch := make([]int, 0, batchSize)

	for j := range totalJobs {
		batch = append(batch, j)

		// When the batch hits 1,000 items, push the batch into the channel
		if len(batch) == batchSize {
			jobs <- batch
			batch = make([]int, 0, batchSize) // Allocate a fresh batch
		}
	}

	// Push any remaining items if totalJobs isn't a multiple of batchSize
	if len(batch) > 0 {
		jobs <- batch
	}

	// 3. Close channel and wait
	close(jobs)
	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Elapsed time: %v | Workers: %d | Batch Size: %d\n", elapsed, numWorkers, batchSize)
}

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
//
// 	processed := 0
//
// 	// Workers dynamically pull jobs as fast as they can handle them.
// 	// The loop stops automatically when the jobs channel is closed.
// 	for range jobs {
// 		processed++
// 	}
//
// 	fmt.Printf("Worker ID %d finished and dynamically processed %d items\n", id, processed)
// }
//
// func main() {
// 	gomaxprocsNum := 4
// 	runtime.GOMAXPROCS(gomaxprocsNum)
//
// 	totalJobs := 10_000_000
// 	numWorkers := 4
//
// 	// Create a buffered channel to hold jobs
// 	jobs := make(chan int, 10_000_000)
// 	var wg sync.WaitGroup
//
// 	start := time.Now()
// 	// 1. Launch 4 worker goroutines waiting for work
// 	for i := range numWorkers {
// 		wg.Add(1)
// 		go worker(i, jobs, &wg)
// 	}
//
// 	// 2. Feed jobs into the queue
// 	for j := range totalJobs {
// 		jobs <- j
// 	}
//
// 	// 3. Close channel to notify workers "no more work coming"
// 	close(jobs)
//
// 	// 4. Wait for all workers to finish
// 	wg.Wait()
//
// 	elapsed := time.Since(start)
// 	fmt.Printf("Elapsed time: %v, worker: %d, GOMAXPROCS: %d\n", elapsed, numWorkers, gomaxprocsNum)
// }

// ========== Sample in udemy show how each worker has same speed. It didn't divide the task into 4
// func heavyTask(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Task id %d starting\n", id)
//
// 	for range 100_000_000 {
// 	}
//
// 	fmt.Println(time.Now())
// 	fmt.Printf("Task id %d finished\n", id)
// }
//
// func main() {
// 	start := time.Now()
// 	runtime.GOMAXPROCS(4)
//
// 	numWorkers := 4
//
// 	var wg sync.WaitGroup
//
// 	for i := range numWorkers {
// 		wg.Add(1)
// 		go heavyTask(i, &wg)
// 	}
// 	wg.Wait()
//
// 	elapsed := time.Since(start)
//
// 	defer fmt.Println(elapsed)
// }

// func printNumbers() {
// 	for i := range 5 {
// 		fmt.Println(time.Now())
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
//
// func printLetters() {
// 	for _, letter := range "ABCDE" {
// 		fmt.Println(time.Now())
// 		fmt.Println(string(letter))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
//
// func main() {
// 	go printNumbers()
// 	go printLetters()
//
// 	time.Sleep(3 * time.Second)
// }
