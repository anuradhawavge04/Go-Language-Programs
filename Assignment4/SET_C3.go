package main

import "fmt"

// Base interfaces
type Shape interface {
	Area() float64
}

type Color interface {
	ShowColor()
}

// Embedded interface
type ColoredShape interface {
	Shape
	Color
}

// Struct implementing both interfaces
type Rectangle struct {
	length float64
	width  float64
	color  string
}

// Implement Area()
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Implement ShowColor()
func (r Rectangle) ShowColor() {
	fmt.Println("Color:", r.color)
}

// Function using embedded interface
func display(cs ColoredShape) {
	fmt.Println("Area:", cs.Area())
	cs.ShowColor()
}

func main() {
	r := Rectangle{length: 10, width: 5, color: "Red"}

	display(r)
}
