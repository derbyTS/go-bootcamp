package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	pool := sync.Pool{}

	pool.Put(&person{
		name: "John",
		age:  30,
	})

	// Get the object from the pool
	person1 := pool.Get().(*person)
	fmt.Println("Got person: ", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got person again: ", person2)

	person3 := pool.Get()
	if person3 != nil {
		fmt.Println("Got person again: ", person3)
	} else {
		fmt.Println("Sync pool is empty")
	}

	// Returning object to the pool
	pool.Put(person2)
	pool.Put(person3)

	fmt.Println("Returned person to pool")

	person4 := pool.Get()
	fmt.Println("Got person again: ", person4)

	person5 := pool.Get()
	if person5 != nil {
		fmt.Println("Got person again: ", person5)
	} else {
		fmt.Println("Sync pool is empty")
	}
}

// func main() {
// 	pool := sync.Pool{
// 		New: func() any {
// 			fmt.Println("creating new person")
// 			return &person{}
// 		},
// 	}
//
// 	// Get the object from the pool
// 	person1 := pool.Get().(*person)
// 	person1.name = "John"
// 	person1.age = 30
// 	fmt.Println("Got person: ", person1)
//
// 	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)
//
// 	pool.Put(person1)
// 	fmt.Println("Returned Person to pool")
//
// 	person2 := pool.Get().(*person)
// 	fmt.Println("Got person again: ", person2)
//
// 	person3 := pool.Get().(*person)
// 	fmt.Println("Got person again: ", person3)
//
// 	// Returning object to the pool
// 	pool.Put(person2)
// 	pool.Put(person3)
//
// 	fmt.Println("Returned person to pool")
//
// 	person4 := pool.Get().(*person)
// 	fmt.Println("Got person again: ", person4)
//
// 	person5 := pool.Get().(*person)
// 	fmt.Println("Got person again: ", person5)
// }
