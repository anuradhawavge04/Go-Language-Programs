package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function
func worker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done()

	// Simulate work
	fmt.Println("Worker", id, "is working...")
	time.Sleep(time.Duration(id) * time.Second)

	fmt.Println("Worker", id, "finished work and waiting at checkpoint")

	// Arrive at checkpoint
	barrier.Done() // signal completion
	barrier.Wait() // wait for others

	// After checkpoint
	fmt.Println("Worker", id, "is assembling parts after checkpoint")
}

func main() {
	var wg sync.WaitGroup
	var barrier sync.WaitGroup

	numWorkers := 4

	// Barrier count = number of workers
	barrier.Add(numWorkers)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &barrier)
	}

	// Wait for all workers to finish completely
	wg.Wait()

	fmt.Println("All workers completed final assembly")
}
