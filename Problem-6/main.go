package main

import "fmt"

func logMessage(ch chan string) {
	ch <- "INFO: server started"
	ch <- "INFO: request received"
	ch <- "WARN: high memory"
	ch <- "ERROR: db timeout"
	ch <- "INFO: request done"
	close(ch)
}

func main() {
	ch := make(chan string, 5)

	logMessage(ch)

	for msg := range ch {
		fmt.Println(msg)
	}

}
