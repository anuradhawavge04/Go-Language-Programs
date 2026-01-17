package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	if num < 10 {
		fmt.Printf("Entered digit is a single digit:%d", num)
	} else {
		fmt.Printf("Entered digit is not a single digit:%d", num)
	}
}
