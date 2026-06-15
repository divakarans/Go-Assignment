package main

import (
	"fmt"
	"time"
)

func slow(ch chan string) {

	time.Sleep(time.Second * 3)
	ch <- "job done"

}

func main() {
	ch := make(chan string)

	go slow(ch)

	select {
	case value := <-ch:
		fmt.Println("result:", value)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
