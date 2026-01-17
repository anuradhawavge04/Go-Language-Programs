package main

import "fmt"

func main() {
	var num1, num2 int = 20, 60
	fmt.Println("Number Before Swapping:", "Num1", num1, "Num2:", num2)
	num1 = num1 - num2
	num2 = num2 + num1
	num1 = num2 - num1
	fmt.Println("Number After Swapping", "Num1:", num1, "Num2:", num2)
}
