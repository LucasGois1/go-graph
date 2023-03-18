package data_structure

type Stack[T any] struct {
	elements []*T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		elements: make([]*T, 0),
	}
}

func (s *Stack[T]) Push(element *T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
