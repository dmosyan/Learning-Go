package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()

	msg = s
}

func main() {
	msg = "Hello, there!"

	wg.Add(2)

	go updateMessage("Hello, world!")
	go updateMessage("Hello, universe!")

	wg.Wait()

	fmt.Println(msg)
}

// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()

// 	m.Lock()
// 	msg = s
// 	m.Unlock()
// }

// func main() {
// 	msg = "Hello, there!"

// 	var mutex sync.Mutex

// 	wg.Add(2)

// 	go updateMessage("Hello, world!", &mutex)
// 	go updateMessage("Hello, universe!", &mutex)

// 	wg.Wait()

// 	fmt.Println(msg)
// }
