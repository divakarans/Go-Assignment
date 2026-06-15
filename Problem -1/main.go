package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {

		defer wg.Done()
		fmt.Println("Hello, Diva!")

	}()

	go func() {

		defer wg.Done()
		fmt.Println("Hello, luffy!")

	}()

	go func() {

		defer wg.Done()
		fmt.Println("Hello, zoro!")

	}()

	go func() {

		defer wg.Done()
		fmt.Println("Hello, harshitha!")

	}()

	go func() {

		defer wg.Done()
		fmt.Println("Hello, Fahad!")

	}()

	fmt.Println("main functions waiting")

	wg.Wait()

	fmt.Println("main func finished")

}
