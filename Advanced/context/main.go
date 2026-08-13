package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	rootCtx := context.Background()
	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second) //Uncomment this with defer cancel() and context.WithValue
	ctx, cancel := context.WithCancel(rootCtx) // Uncomment this with go func

	// defer cancel()

	go func() {
		// Simulate "Heavy Work" by doing something computationally expensive
		sum := 0
		for i := 0; i <= 1000000; i++ {
			sum += i
			// In a real heavy task, you might check ctx here too
		}
		ctx = context.WithValue(ctx, "requestID", strconv.Itoa(sum))
		fmt.Println("Heavy task finished calculation, now cancelling...")
		cancel()
	}()

	// ctx = context.WithValue(ctx, "requestID", "123dfaskodk412")

	go doWork(ctx)

	time.Sleep(4 * time.Second)

	requestID := ctx.Value("requestID")

	if requestID != nil {
		fmt.Println("Request ID: ", requestID)
	} else {
		fmt.Println("No Request ID found")
	}

	logWithContext(ctx, "This is a test")
}

func logWithContext(ctx context.Context, message string) {
	requestID := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v", requestID, message)
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled: ", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// Cancel sample
// func main() {
// 	ctx := context.TODO()
// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result from context TODO: ", result)
//
// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancelling(cancel)
//
// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result from timeout context: ", result)
//
// 	time.Sleep(2 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout: ", result)
// }
//
// func cancelling(c context.CancelFunc) {
// 	fmt.Println("Will cancel now")
// 	c()
// }
//
// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

//TODO  and Background sample
// func main() {
// 	todoContext := context.TODO()
// 	contextBkg := context.Background()
//
// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))
//
// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }
