package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))

	pi := 3.14
	d := int(pi)
	fmt.Println(d)

	e := "Hello"
	var f []byte
	f = []byte(e)
	fmt.Println(string(f))

	g := []rune{'a', 'テ', '🩷'}
	fmt.Println(g)
	h := []byte{97}
	fmt.Println(string(h))
	i := string(h[0])
	fmt.Println(i)
}
