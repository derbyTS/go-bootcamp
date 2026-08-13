package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker id: %d\n", id)

	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(id) * time.Second)
		close(ch)
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout happen")
	case <-ch:
		fmt.Printf("Worker id: %d is Done\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(ctx, i, &wg)
	}

	wg.Wait()
}
