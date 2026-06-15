package main

import "fmt"

func doub(n int, ch chan int) {
	ch <- n * 2
}

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	ch := make(chan int)

	for _, n := range numbers {
		go doub(n, ch)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-ch
	}

	fmt.Println("sum", sum)

}
