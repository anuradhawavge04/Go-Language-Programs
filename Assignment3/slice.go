package main

import "fmt"

func main() {
	number := []int{
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println("Slice elements are:")
	for i, value := range number {
		fmt.Printf("Index %d:%d\n", i, value)
	}
}
