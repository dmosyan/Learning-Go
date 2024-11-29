package main

import "golang.org/x/exp/constraints"

type Element[T constraints.Ordered] struct {
	value T
	next  *Element[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *Element[T]
	size int
}
