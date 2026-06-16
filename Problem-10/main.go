package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	DSN string
}

var (
	cfg  Config
	once sync.Once
)

func loadConfig() {
	fmt.Println("Loading config...")
	time.Sleep(100 * time.Millisecond)

	cfg = Config{
		DSN: "postgres://localhost/mydb",
	}
}

func GetConfig() Config {
	once.Do(loadConfig)
	return cfg
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			config := GetConfig()
			fmt.Println(config.DSN)
		}()
	}

	wg.Wait()
}
