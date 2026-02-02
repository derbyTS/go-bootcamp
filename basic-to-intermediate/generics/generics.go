package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b := swap('a', 'b')
	fmt.Printf("%c %c\n", a, b)

	i := Stack[int]{}

	// TypeOf assertion
	var v any = 123
	if s, ok := v.(int); ok {
		fmt.Println(ok)
		fmt.Println("s is ", s)
	}
	var v1 any = "123"
	if s, ok := v1.(int); ok {
		fmt.Println(ok)
		fmt.Println("s is ", s)
	}

	i.push(1)
	i.push(2)
	i.push(3)
	i.push(4)
	i.push(5)
	i.printAll()
	fmt.Println(i.isStackEmpty())
	p, isDeleteSuccess := i.pop()
	fmt.Println(p)
	fmt.Println(isDeleteSuccess)
	p, isDeleteSuccess = i.pop()
	fmt.Println(i.isStackEmpty())
	fmt.Println(p)
	fmt.Println(isDeleteSuccess)

	s := Stack[string]{}
	s.push("hello")
	s.push("world")
	s.printAll()
	sp, _ := s.pop()
	fmt.Println(sp)
	sp, _ = s.pop()
	fmt.Println(sp)
	sp, _ = s.pop()
	fmt.Println(sp == "")
	fmt.Println(len(sp) == 0)
	var test string
	fmt.Println("This is test", len(test))
	fmt.Println(reflect.TypeOf(sp))
	myPrint(a, 'a')
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(e T) {
	s.elements = append(s.elements, e)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s *Stack[T]) isStackEmpty() bool {
	return len(s.elements) <= 0
}

func (s *Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack elements: ")
	for _, element := range s.elements {
		fmt.Print(" ", element)
	}
	fmt.Println("")
}

func myPrint(i ...any) {
	for _, v := range i {
		// TypeOf via switch
		switch v.(type) {
		case int:
			fmt.Println("The type of ", v, " is int")
		case string:
			fmt.Println("The type of ", v, " is string")
		case rune:
			fmt.Println("The type of ", v, " is rune")
		default:
			fmt.Println("The type of ", v, reflect.TypeOf(v))
		}
	}
}
