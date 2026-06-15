package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	delay := 100 + rand.Intn(401)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Println("downloaded", filename)
}

func main() {

	files := []string{
		"file_1.txt",
		"file_2.txt",
		"file_3.txt",
		"file_4.txt",
		"file_5.txt",
	}

	var wg sync.WaitGroup

	wg.Add(len(files))

	for _, file := range files {
		go download(file, &wg)

	}

	wg.Wait()

	fmt.Println("All downloads complete")

}
