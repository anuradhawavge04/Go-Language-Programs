package main

import "fmt"

// Function to calculate sum of squares
func squareSum(n int, ch chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n = n / 10
	}
	ch <- sum
}

// Function to calculate sum of cubes
func cubeSum(n int, ch chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit * digit
		n = n / 10
	}
	ch <- sum
}

func main() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)

	// Channels
	sqChan := make(chan int)
	cuChan := make(chan int)

	// Goroutines
	go squareSum(number, sqChan)
	go cubeSum(number, cuChan)

	// Receive values
	squares := <-sqChan
	cubes := <-cuChan

	// Output
	fmt.Println("Sum of squares =", squares)
	fmt.Println("Sum of cubes =", cubes)
	fmt.Println("Final sum =", squares+cubes)
}
