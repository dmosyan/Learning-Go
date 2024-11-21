package main

import "fmt"

func main() {
	myQueue := Queue{
		elements: []any{1, 3, 4, 5, "something", 1.25},
	}

	myQueue.Enqueue("new")
	myQueue.Dequeue()

	fmt.Println(myQueue)

	myStack := Stack{
		elements: []any{1, 2, 3, 4, 5, "stack", 5.21},
	}

	myStack.Peek()
	myStack.Pop()

	fmt.Println(myStack)
}
