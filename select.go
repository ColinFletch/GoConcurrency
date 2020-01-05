package main

import (
	"fmt"
	"time"
)

func statement() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every other second"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select { //pick whichever channel is available, don't let one channel block
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
