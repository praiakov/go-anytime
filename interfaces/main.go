package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	widht, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.widht * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2*r.widht + 2*r.height
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func Calculate(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rectangle{widht: 5, height: 8}
	c := circle{radius: 6}

	Calculate(r)
	Calculate(c)
}
