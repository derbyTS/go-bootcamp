package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker ID: %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("Worker ID: %d finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup
	tasks := []string{"apple", "ball", "cat"}

	wg.Add(
		len(tasks) + 1,
	) // This causes a permanent deadlock. You are telling the WaitGroup to expect 4 completions, but you only spawn 3 goroutines. Because no one is alive to call Done() for that remaining slot, wg.Wait() will hang forever.
	for i, task := range tasks {
		worker := Worker{
			ID:   i + 1,
			Task: task,
		}
		// wg.Add(1) //This is a good practice
		go worker.PerformTask(&wg)
	}

	wg.Wait()
	fmt.Println("Done")
}

// ***************** sample with channels, worker and jobs

// Sample without channel consumer 'tasks'
// func worker(id int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker ID: %d staring.\n", id)
// 	time.Sleep(time.Second)
// 	result <- id * 2
// 	fmt.Printf("Worker ID: %d finished.\n", id)
// }

// Sample with channel consumer 'tasks'
// func worker(id int, tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker ID: %d staring.\n", id)
// 	time.Sleep(time.Second)
// 	for task := range tasks {
// 		result <- task * 2
// 		// wg.Done() //Uncomment
// 	}
// 	fmt.Printf("Worker ID: %d finished.\n", id)
// }
//
// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)
//
// 	// numJobs is 5, so you call wg.Add(5).
// 	// The WaitGroup counter is now 5.
// 	// You spin up 3 workers.
// 	// The tasks are processed, the tasks channel is closed, and the 3 workers finish their for task := range tasks loops.
// 	// As those 3 workers exit, they each hit their defer wg.Done().\
// 	// The counter drops: 5 - 1 - 1 - 1 = 2.
// 	// The workers are all dead and gone. No one else is alive to call Done().
// 	// wg.Wait() sits there forever waiting for that counter to drop from 2 to 0. Go notices your program is completely stuck with no way to wake up, and crashes with a deadlock error.
// 	// To run with this find Uncomment comments
// 	// wg.Add(numJobs)
//
// 	wg.Add(numWorkers)
//
// 	for i := range numWorkers {
// 		go worker(i, tasks, results, &wg)
// 	}
//
// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks)
//
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()
//
// 	for result := range results {
// 		fmt.Printf("Results: %d\n", result)
// 	}
// }

// ***************** basic sample without channels

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1) // Wrong
// 	fmt.Printf("Worker id: %d\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker id: %d is Done\n", id)
// }
//
// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
//
// 	wg.Add(numWorkers)
//
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}
//
// 	wg.Wait()
// }
