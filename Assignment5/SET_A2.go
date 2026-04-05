package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function executed by each goroutine
func printNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)

		// Random delay between 0 to 250 ms
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	// Start 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumbers(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines finished execution")
}
