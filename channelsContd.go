package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(8 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(9 * time.Second)
		c2 <- "two"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("exitng wait timeout over ")
	}

}
