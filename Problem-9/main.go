package main

import (
	"fmt"
	"time"
)

func ticker(done chan struct{}, out chan int) {
	count := 1

	for {
		select {
		case <-done:
			return

		case <-time.After(200 * time.Millisecond):
			out <- count
			count++
		}
	}
}

func main() {
	done := make(chan struct{})
	out := make(chan int)

	go ticker(done, out)

	for i := 0; i < 5; i++ {
		fmt.Println(<-out)
	}

	close(done)

	fmt.Println("ticker stopped")
}
