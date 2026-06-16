package main

import "fmt"

func Square(ch chan int, n int) {

	for i := 1; i <= n; i++ {
		ch <- i * i
	}
	close(ch)
}

func main() {

	ch := make(chan int)

	go Square(ch, 8)

	for Sq := range ch {
		fmt.Println("Square is: ", Sq)
	}

}
