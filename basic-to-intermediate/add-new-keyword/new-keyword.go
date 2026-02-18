package main

import "fmt"

type Player struct {
	Name  string
	Level int
}

func main() {
	// 1. Using new
	// This gives us a *Player (pointer) with Name: "" and Level: 0
	p1 := new(Player)
	fmt.Printf("Type: %T, Value: %+v\n", p1, p1)

	// 2. Using a Composite Literal (The preferred way)
	// This also gives us a *Player, but we can set values immediately
	p2 := &Player{
		Name:  "Gopher",
		Level: 10,
	}
	fmt.Printf("Type: %T, Value: %+v\n", p2, p2)

	p3 := Player{
		Name:  "Gopher",
		Level: 10,
	}
	fmt.Printf("Type: %T, Value: %+v\n", p3, p3)
}
