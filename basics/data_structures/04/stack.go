package main

import "errors"

type Stack struct {
	elements []any
}

// Add a new element
func (s *Stack) Push(el any) {
	s.elements = append(s.elements, el)
}

// Remove the last element(FIFO)
func (s *Stack) Pop() (el any, err error) {

	if s.IsEmpty() {
		err = errors.New("stack is empty")
		return
	}

	el = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

// Return the last element only
func (s *Stack) Peek() (el any, err error) {

	if s.IsEmpty() {
		err = errors.New("stack is empty")
		return
	}

	el = s.elements[len(s.elements)-1]
	return
}

// Check if the stack size is 0
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Get the size of the stack (number of elements)
func (s *Stack) Size() int {
	return len(s.elements)
}
