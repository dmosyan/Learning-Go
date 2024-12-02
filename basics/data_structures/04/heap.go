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

func (h *Heap[T]) Push(ele T) {
	h.nodes = append(h.nodes, ele)
	i := len(h.nodes) - 1

	for ; h.nodes[i] > h.nodes[parent(i)]; i = parent(i) {
		h.swap(i, parent(i))
	}
}

func (h *Heap[T]) Pop() (ele T) {
	ele = h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.rearrange(0)
	return
}

func (h *Heap[T]) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *Heap[T]) rearrange(i int) {

	largest := i
	left, right, size := leftChild(i), rightChild(i), len(h.nodes)

	if left < size && h.nodes[left] > h.nodes[largest] {
		largest = left
	}
	if right < size && h.nodes[right] > h.nodes[largest] {
		largest = right
	}
	if largest != i {
		h.swap(i, largest)
		h.rearrange(largest)
	}
}

func (h *Heap[T]) Size() int {
	return len(h.nodes)
}
func (h *Heap[T]) IsEmpty() bool {
	return h.Size() == 0
}
