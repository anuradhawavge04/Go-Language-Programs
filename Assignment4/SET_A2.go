package main

import "fmt"

// Method to print multiplication table
func printTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Call method
	printTable(number)
}
