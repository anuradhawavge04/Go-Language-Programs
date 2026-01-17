package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	fmt.Printf("Value of a:%d\n", a)
	fmt.Printf("Value of b:%d\n", b)

	result := sum(a, b)
	fmt.Println("Sum of a and b is:", result)

}

func sum(num1, num2 int) int {

	result := num1 + num2
	return result
}
