package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var err error
	fmt.Println("Before")
	go sayHello()
	fmt.Println("After")

	go printNumbers()
	go printLetters()

	go func() {
		err = doWork()
	}()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("doWork Success")
	}

	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("doWork Success Outside")
	}
	t := runtime.GOMAXPROCS(0)
	fmt.Println(t)
	fmt.Println(runtime.NumGoroutine())
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello Goroutine")
}

func printNumbers() {
	for i := range 5 {
		fmt.Println("number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("letter: ", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Error here\n")
}
