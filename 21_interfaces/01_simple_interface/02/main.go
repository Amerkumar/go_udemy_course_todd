package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Printf("%T \n", z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	var k Shape
	fmt.Printf("%T \n", k)
	k = Square{100}
	fmt.Printf("%T \n", k)
	info(s)
	info(c)
}
