package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // wait for 1 goroutine

	go func() {
		count("sheep")
		wg.Done() // wg --
	}()

	wg.Wait() //Block until 0
}

func count(thing string) {
	for i := 1; i <= 5; i += 2 {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 250)
	}
}
