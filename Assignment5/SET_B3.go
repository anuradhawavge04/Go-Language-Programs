package main

import "fmt"

// Function to send data to channel
func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // close the channel
}

func main() {
	// Create channel
	ch := make(chan int)

	// Start goroutine
	go sendData(ch)

	// Read using for range (automatically stops when channel is closed)
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
