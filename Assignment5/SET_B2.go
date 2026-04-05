package main

import "fmt"

// Goroutine to generate Fibonacci series
func fibonacci(n int, ch chan int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}

	close(ch) // close channel after writing
}

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	ch := make(chan int)

	// Start goroutine
	go fibonacci(n, ch)

	// Read from channel
	fmt.Println("Fibonacci Series:")
	for num := range ch {
		fmt.Println(num)
	}
}