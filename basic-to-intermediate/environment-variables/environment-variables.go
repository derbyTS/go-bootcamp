package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env vairable: ", user)
	fmt.Println("Home env vairable: ", home)

	err := os.Setenv("FRUIT", "APPLE") // It set the env variable during run time only
	if err != nil {
		fmt.Println(err)
	}
	// err = os.Unsetenv("P9K_SSH")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	osEnv := os.Environ()

	fruit := os.Getenv("FRUIT")
	fmt.Println("Fruit environment: ", fruit)

	for i, v := range osEnv {
		s := strings.SplitN(v, "=", 2)
		fmt.Printf("%d. %s: %s\n", i, s[0], s[1])
	}
}
