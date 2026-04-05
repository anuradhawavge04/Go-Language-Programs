package main

import "fmt"

func main() {
	// Declare interface
	var i interface{}

	// Assign value to interface
	i = "Hello Go"

	// Type assertion
	value, ok := i.(string)

	if ok {
		fmt.Println("Value:", value)
		fmt.Println("Type assertion successful")
	} else {
		fmt.Println("Type assertion failed")
	}
}
