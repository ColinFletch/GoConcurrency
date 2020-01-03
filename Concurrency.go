package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // Channel
	go count("sheep", c)
	for msg := range c { // blocking call
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i += 2 {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // close channel when finished sending
}
