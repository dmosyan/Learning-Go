package main

type Set struct {
	elements map[any]bool
}

func NewSet() Set {
	return Set{elements: make(map[any]bool)}
}

func (s *Set) Add(el any) {
	s.elements[el] = false
}

func (s *Set) Remove(el any) {
	delete(s.elements, el)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}
func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() (list []any) {
	for k := range s.elements {
		list = append(list, k)
	}
	return
}
