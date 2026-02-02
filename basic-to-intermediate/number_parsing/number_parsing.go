package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	numberString := "123"
	num, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(numberString))
	fmt.Println(reflect.TypeOf(num))

	var v any = "123"
	if s, ok := v.(string); ok {
		fmt.Println(s)
	}
	base2 := "1010101"

	num1, err := strconv.ParseInt(base2, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	hex := "ff"

	hexValue, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hexValue)
	fmt.Printf("%T\n", hexValue)

	fmt.Println(num1)
	num2, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num2)

	floatPi := "3.14"
	piValue, err := strconv.ParseFloat(floatPi, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%.2f\n", piValue)
	fmt.Printf("%T\n", piValue)

	piString := strconv.FormatFloat(piValue, 'f', -1, 64)
	fmt.Println(piString)
	fmt.Printf("%T\n", piString)
}
