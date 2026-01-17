package main

import "fmt"

func main() {
	var a int = 45
	var p1 *int
	var p2 **int

	p1 = &a
	p2 = &p1
	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *p1= %d\n", *p1)
	fmt.Printf("value available at **p2= %d\n", **p2)
}
