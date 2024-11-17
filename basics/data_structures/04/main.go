package main

import "fmt"

func main() {
	myQueue := Queue{
		elements: []any{1, 3, 4, 5, "something", 1.25},
	}

	myQueue.Enqueue("new")
	myQueue.Dequeue()

	fmt.Println(myQueue)
}
