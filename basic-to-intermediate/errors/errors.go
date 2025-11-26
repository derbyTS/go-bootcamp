package main

import "fmt"

func main() {
	quotient, err := quo(2, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)

	pTest := personTest{
		isJokeOk: true,
	}

	fmt.Println(pTest.isTheJokeValid())

	fmt.Println(getPeopleOpinion().isTheJokeValid())
}

type myError struct {
	message string
}

func (e myError) Error() string {
	return fmt.Sprintf("Custom error: %s", e.message)
}

func quo(x, y int) (int, error) {
	if y == 0 {
		// return -1, &myError{"Cannot divide to zero"}
		return -1, &myError{message: "Cannot divide to zero"}
	}

	return x / y, nil
}

// The implementation of the error above is like this

type jokeInterface interface {
	isTheJokeValid() bool
}

type personTest struct {
	isJokeOk bool
}

func (p personTest) isTheJokeValid() bool {
	return p.isJokeOk
}

// So you have an interface that works above now implement the interface by creating a function with the same signature and make jokeInterface as return type

type peopleListining struct {
	isJokeOk bool
}

func (p peopleListining) isTheJokeValid() bool {
	return p.isJokeOk
}

func getPeopleOpinion() jokeInterface {
	return &peopleListining{true}
}
