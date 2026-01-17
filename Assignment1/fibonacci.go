package main

import (
	"fmt"
)

func main() {
	var n int = 5
	a, b := 0, 1
	fmt.Println("Fibonacci Sequence:", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", a)
		a, b = b, a+b
	}
}
