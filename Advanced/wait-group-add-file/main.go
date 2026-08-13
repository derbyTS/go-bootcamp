package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// The worker function loops until the jobs channel is closed
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range jobs {
		// Process your line or DB row here
		fmt.Printf("Worker %d processed: %s\n", id, line)
	}
}

func main() {
	file, _ := os.Open("huge_file.txt")
	defer file.Close()

	jobs := make(chan string, 100) // Buffered channel to stream lines
	var wg sync.WaitGroup

	// 1. Spin up a fixed number of parallel workers (e.g., 4)
	numWorkers := 4
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // We only add 1 per worker!
		go worker(w, jobs, &wg)
	}

	// 2. Stream the file line by line without knowing the total count
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text() // Blocks if the channel buffer is full
	}

	// 3. Close the channel to signal workers that no more work is coming
	close(jobs)

	// 4. Wait for the workers to drain the channel and finish
	wg.Wait()
	fmt.Println("All done processing huge file!")
}
