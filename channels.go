package main

import (
	"fmt"
	"time"
)

func worker(done chan<- bool) {

	fmt.Println("worker")
	done <- true

}

func main() {
	done := make(chan<- bool)

	go worker(done)
	time.Sleep(time.Second)

}
