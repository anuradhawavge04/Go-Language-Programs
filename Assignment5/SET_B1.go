package main

import "fmt"

func main() {
	// Create buffered channel with capacity 5
	ch := make(chan int, 5)

	// Insert values
	ch <- 10
	ch <- 20
	ch <- 30

	// Check capacity and length
	fmt.Println("Channel Capacity:", cap(ch))
	fmt.Println("Channel Length:", len(ch))

	// Read one value
	fmt.Println("Removed:", <-ch)

	// Check length after removing one element
	fmt.Println("Channel Length after removing one value:", len(ch))

	// Read remaining values
	fmt.Println("Removed:", <-ch)
	fmt.Println("Removed:", <-ch)

	// Final length
	fmt.Println("Final Channel Length:", len(ch))
}