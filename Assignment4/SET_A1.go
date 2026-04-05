package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define Circle struct
type Circle struct {
	radius float64
}

// Implement Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Define Rectangle struct
type Rectangle struct {
	length float64
	width  float64
}

// Implement Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Implement Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Main function
func main() {
	c := Circle{radius: 5}
	r := Rectangle{length: 10, width: 4}

	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())

	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perimeter:", r.Perimeter())
}
