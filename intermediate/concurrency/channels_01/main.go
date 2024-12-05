package main

import "fmt"

func main() {
	// create buffered channel to make this work
	ch := make(chan string, 1)

	ch <- "message"

	fmt.Println(<-ch)

	directionalChannels()
}

func directionalChannels() {
	//bidirectional channel
	ch := make(chan string)

	// send-only channel
	go func(ch chan<- string) {
		ch <- "message"
	}(ch)

	// receive-only channel
	go func(ch <-chan string) {
		fmt.Println(<-ch)
	}(ch)
}
