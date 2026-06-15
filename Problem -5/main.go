package main

import (
	"fmt"
	"sync"
)

func cnt(count *int, mu *sync.Mutex, wg *sync.WaitGroup) {

	defer wg.Done()
	mu.Lock()
	*count++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go cnt(&count, &mu, &wg)
	}

	wg.Wait()

	fmt.Println("Total count : ", count)

}
