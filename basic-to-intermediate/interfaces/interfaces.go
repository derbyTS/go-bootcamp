package main

import (
	"fmt"
	"math"
)

func main() {
	r := rectangle{
		width:  10,
		length: 5,
	}

	c := circle{
		radius: 12,
	}

	measure(&r)
	measure(&c)

	fmt.Println("rectangle ....................")
	fmt.Println(getGeometry("rectangle", 10, 10))
	fmt.Println(getGeometry("rectangle", 10, 10).area())
	fmt.Println(getGeometry("rectangle", 10, 10).perimeter())
	fmt.Println("circle ....................")
	fmt.Println(getGeometry("circle", 10, 0))
	fmt.Println(getGeometry("rectangle", 10, 0).area())
	fmt.Println(getGeometry("rectangle", 10, 0).perimeter())
	cir := getGeometry("rectangle", 10, 0)
	fmt.Println(cir)
	fmt.Println(cir.area())
	fmt.Println(cir.perimeter())
}

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	length float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return (r.length + r.width) * 2
}

func (c circle) area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func getGeometry(kind string, a, b float64) geometry {
	switch kind {
	case "rectangle":
		return rectangle{width: a, length: b}
	case "circle":
		return rectangle{width: a, length: b}
	default:
		return nil
	}
}
