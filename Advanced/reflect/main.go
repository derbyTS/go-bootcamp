package main

import (
	"fmt"
	"reflect"
)

// ======== Working with methods

type Greeter struct{}

func (g Greeter) Greet(name, lname string) string {
	return "Hello " + name + " " + lname
}

func main() {
	g := Greeter{}
	t := reflect.TypeFor[Greeter]() // LSP recommend
	// t := reflect.TypeOf(g)
	v := reflect.ValueOf(g)
	var method reflect.Method

	fmt.Println(g.Greet("Alice", "Doe"))

	fmt.Println("Type: ", t)
	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d : %s\n", i, method.Name)
	}

	m := v.MethodByName(method.Name)

	results := m.Call([]reflect.Value{reflect.ValueOf("Alice"), reflect.ValueOf("Doe")})

	fmt.Printf("Greet result: %s\n", results[0].String())
}

// ======== Working with structs and fields

// type Person struct {
// 	Name string
// 	Age  int
// }
//
// func main() {
// 	p := Person{
// 		Name: "Alice",
// 		Age:  30,
// 	}
// 	v := reflect.ValueOf(p)
//
// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}
//
// 	v1 := reflect.ValueOf(&p).Elem()
//
// 	nameField := v1.FieldByName("Name")
//
// 	// Remember if the struct is `unexported` or started in lowerCase, it cannot set
// 	if nameField.CanSet() {
// 		nameField.SetString("Bob")
// 	} else {
// 		fmt.Println(
// 			"cannot set",
// 		)
// 	}
//
// 	fmt.Println("Modified person: ", p)
// }

// func main() {
// 	x := 42
// 	v := reflect.ValueOf(x)
// 	t := v.Type()
//
// 	fmt.Println("Value: ", v)
// 	fmt.Println("Type: ", t)
// 	fmt.Println("Kind: ", t.Kind())
// 	fmt.Println("isInt: ", t.Kind() == reflect.Int)
// 	fmt.Println("isString: ", t.Kind() == reflect.String)
// 	fmt.Println("isZero: ", v.IsZero())
//
// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)
//
// 	fmt.Println("v2 Type: ", v2.Type())
//
// 	fmt.Println("Original Value: ", v1.Int())
//
// 	v1.SetInt(18)
//
// 	fmt.Println("Modified Value: ", v1.Int())
//
// 	// var intf interface{} = "Hello"
// 	var intf any = "Hello"
//
// 	v3 := reflect.ValueOf(intf)
// 	fmt.Println("v3 Type: ", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println(v3)
// 	}
// }
