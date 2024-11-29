package main

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Element[T constraints.Ordered] struct {
	value T
	next  *Element[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *Element[T]
	size int
}

func (l *LinkedList[T]) Add(el *Element[T]) {

	if l.head == nil {
		l.head = el
	} else {
		el.next = l.head
		l.head = el
	}

	l.size++
}

func (l *LinkedList[T]) Insert(el *Element[T], marker T) error {

	for current := l.head; current.next != nil; current = current.next {
		if current.value == marker {
			el.next = current.next
			current.next = el
			l.size++
			return nil
		}
	}

	return errors.New("element not found")

}

func (l *LinkedList[T]) Delete(el *Element[T]) error {

	current := l.head
	prev := l.head

	for current != nil {

		if current.value == el.value {
			if current == l.head {
				l.head = current.next
			} else {
				prev.next = current.next
			}

			l.size--
			return nil
		}

		prev = current
		current = current.next
	}

	return errors.New("element not found")
}
