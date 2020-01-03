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

	d := make(chan string, 2) // buffer allows us to store hello and world
	d <- "hello"
	d <- "world"

	msg2 := <-d
	fmt.Println(msg2)

	msg2 = <-d
	fmt.Println(msg2)

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i += 2 {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // close channel when finished sending
}
