package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// (c circle) is a value receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// (c *circle) is a pointer receiver
func (c *circle) area1() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println("area:", z.area())
}

func main() {
	c := circle{5}
	info(c)
	info(&c)
}
