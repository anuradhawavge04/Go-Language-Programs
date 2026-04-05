package main

import "fmt"

// Function that uses type switch
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Type is int, Value:", v)
	case string:
		fmt.Println("Type is string, Value:", v)
	case float64:
		fmt.Println("Type is float64, Value:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType(10)
	checkType("Hello Go")
	checkType(3.14)
	checkType(true)
}
