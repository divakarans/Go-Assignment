package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range jobs {
		fmt.Printf("Worker %d processed order %d\n", id, order)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan int, 12)

	var wg sync.WaitGroup

	// Start 4 workers
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send 12 orders
	for order := 1; order <= 12; order++ {
		jobs <- order
	}

	// No more jobs
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All orders processed")
}
