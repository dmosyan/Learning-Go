package main

import "golang.org/x/exp/constraints"

type Heap[T constraints.Ordered] struct {
	nodes []T
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return (2 * i) + 1
}

func rightChild(i int) int {
	return (2 * i) + 2
}
