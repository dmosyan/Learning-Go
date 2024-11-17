package main

import "errors"

type Queue struct {
	elements []any
}

// Add a new element
func (q *Queue) Enqueue(el any) {
	q.elements = append(q.elements, el)
}

// Remove the first element
func (q *Queue) Dequeue() (el any, err error) {
	if q.IsEmtpy() {
		err = errors.New("empty queue")
	}

	el = q.elements[0]
	q.elements = q.elements[1:]
	return
}

// Check if the queue size is 0
func (q *Queue) IsEmtpy() bool {
	return q.Size() == 0
}

// Get the length of the queue
func (q *Queue) Size() int {
	return len(q.elements)
}
