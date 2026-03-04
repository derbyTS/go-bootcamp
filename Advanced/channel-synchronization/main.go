package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	data := make(chan string)
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	go func() {
		defer close(
			data,
		) // According to LLM: "It is a best practice to use defer close(data) inside the goroutine." "The sender should be the one to close the channel."
		scanner := bufio.NewScanner(file)

		// Read the file line by line
		for scanner.Scan() {
			data <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Something error")
		}
		// close(data)
	}()

	for value := range data {
		fmt.Printf("%s Received value: %s\n", time.Now().Format("15:04:05"), value)
	}
}

// === Synchronization of multiple Goroutine
// func main() {
// 	numGoroutines := 13
// 	ch := make(chan int, numGoroutines)
//
// 	for i := range numGoroutines {
// 		time.Sleep(time.Second)
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working.....\n", id)
// 			ch <- id
// 		}(i)
// 	}
// 	for range numGoroutines { // You can always make it less but not more than numGoroutines
// 		fmt.Printf("Channel received %d.....\n", <-ch)
// 	}
//
// 	fmt.Println("All Goroutine are finished")
// }

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 9 // Blocking until the value is received
// 		time.Sleep(3 * time.Second)
// 		fmt.Println("value sent")
// 	}()
//
// 	value := <-ch // Blocking until the value is received
// 	fmt.Println(value)
// 	fmt.Println("End")
// }

// func main() {
// 	done := make(chan ResponseStruct)
// 	// done := make(chan struct{}) //Blank struct
// 	go func() {
// 		fmt.Println("Working ......")
// 		time.Sleep(2 * time.Second)
// 		// done <- struct{}{} //Passing blank struct
// 		done <- ResponseStruct{
// 			isOk: true,
// 		}
// 	}()
// 	fmt.Println(<-done)
// 	fmt.Println("Finished ......")
// }

type ResponseStruct struct {
	isOk bool
}
