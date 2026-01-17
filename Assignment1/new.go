package main

import "fmt"

func main() {
	var p *int
	p = new(int)

	fmt.Println("Value Stored:", *p)
	fmt.Println("Address:", p)

	*p = 25
	fmt.Println("Updated value:", *p)
}
