package main

import "fmt"

// Goroutine to handle even numbers
func evenWorker(evenCh chan int, done chan bool) {
	for num := range evenCh {
		fmt.Println("Even:", num)
	}
	done <- true
}

// Goroutine to handle odd numbers
func oddWorker(oddCh chan int, done chan bool) {
	for num := range oddCh {
		fmt.Println("Odd:", num)
	}
	done <- true
}

func main() {
	// Slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenCh := make(chan int)
	oddCh := make(chan int)
	done := make(chan bool)

	// Start goroutines
	go evenWorker(evenCh, done)
	go oddWorker(oddCh, done)

	// Send numbers to respective channels
	for _, num := range numbers {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	// Close channels
	close(evenCh)
	close(oddCh)

	// Wait for both goroutines
	<-done
	<-done
}