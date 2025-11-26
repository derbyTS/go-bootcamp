package main

import (
	"fmt"
)

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "Paranaque",
			country: "Ph",
		},
		Phone: Phone{
			phoneNumber: "123",
		},
	}
	p1 := Person{
		firstName: "jane",
		age:       30,
	}

	fmt.Println(p)
	fmt.Println(p1)

	user := struct {
		userName string
		password string
	}{
		userName: "user123",
		password: "pass",
	}

	fmt.Println(user)

	fmt.Println(p.fullName())
	fmt.Println(p1.fullName())

	fmt.Println(p)
	p.increaseAge()
	fmt.Println(p)
	increaseManual(&p.age)
	p.changeCity("Bacoor")
	fmt.Println(p)
	fmt.Println(p.address)
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	Phone
}

func (p *Person) increaseAge() {
	p.age++
}

func increaseManual(age *int) {
	*age++
}

type Address struct {
	city    string
	country string
}

func (p *Person) changeCity(city string) {
	p.address.city = city
}

type Phone struct {
	phoneNumber string
}
