package dynamic_stack

const _INITIAL_CAPACITY = 10
const _INCREASE_FACTOR = 2
const _DECREASE_FACTOR = 2
const _SHRINK_THRESHOLD = 4

// Definition of the dynamic stack struct

type dynamicStack[T any] struct {
	data     []T
	count    int
}

func (s *dynamicStack[T]) IsEmpty() bool {
	return s.count == 0
}

func (s *dynamicStack[T]) Top() T {
	if s.IsEmpty() {
		panic("The stack is empty")
	}
	return s.data[s.count-1]
}

func (s *dynamicStack[T]) Push(element T) {
	if s.count == cap(s.data) {
		s.resize(cap(s.data) * _INCREASE_FACTOR)
	}

	s.data[s.count] = element
	s.count++
}

func (s *dynamicStack[T]) Pop() T {
	if s.IsEmpty() {
		panic("The stack is empty")
	}

	if s.count <= cap(s.data)/_SHRINK_THRESHOLD && cap(s.data) > _INITIAL_CAPACITY {
		s.resize(cap(s.data) / _DECREASE_FACTOR)
	}

	element := s.data[s.count-1]
	s.count--

	return element
}

func (s *dynamicStack[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	copy(newData, s.data)
	s.data = newData
}

func NewDynamicStack[T any]() Stack[T] {
	s := new(dynamicStack[T])
	s.data = make([]T, _INITIAL_CAPACITY)
	return s
}
