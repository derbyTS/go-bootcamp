package main

import (
	"fmt"

	"example.com/project/client"
	"example.com/project/service"
)

func main() {
	c := client.RealClient{}
	s := service.Service{C: c}

	fmt.Println(client.Sample(3))

	fmt.Println(s.Process())
}
