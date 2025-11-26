package main

import (
	"fmt"
	"retry-interfaces/interfaces/mdas"
)

func main() {
	// Can use as Adder
	var a mdas.Adder = mdas.MdasWriter{}
	fmt.Println("Sum:", a.Add(3, 2))

	// Or as Differ
	var d mdas.Differ = mdas.MdasWriter{}
	fmt.Println("Diff:", d.Dif(3, 2))

	// Or directly
	m := mdas.MdasWriter{Offset: 1}
	fmt.Println("Sum with offset:", m.Add(3, 2))  // 6
	fmt.Println("Diff with offset:", m.Dif(3, 2)) // 2
}
