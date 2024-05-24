package linked_list

type nodeList[T any] struct {
	data T
	next *nodeList[T]
}

type linkedList[T any] struct {
	first *nodeList[T]
	last  *nodeList[T]
	length int
}

type listIterator[T any] struct {
	list     *linkedList[T]
	current  *nodeList[T]
	previous *nodeList[T]
}

// List Primitives

func (l *linkedList[T]) IsEmpty() bool {
	return l.length == 0 && l.first == nil && l.last == nil
}

func (l *linkedList[T]) InsertFirst(data T) {
	newNode := l.createNode(data)

	if l.IsEmpty() {
		l.last = newNode
	} else {
		newNode.next = l.first
	}
	l.first = newNode

	l.length++
}

func (l *linkedList[T]) InsertLast(data T) {
	newNode := l.createNode(data)

	if l.IsEmpty() {
		l.first = newNode
	} else {
		l.last.next = newNode
	}

	l.last = newNode
	l.length++
}

func (l *linkedList[T]) DeleteFirst() T {
	if l.IsEmpty() {
		panic("The list is empty")
	}

	first := l.first.data

	if l.first.next == nil {
		l.last = nil
	}
	l.first = l.first.next

	l.length--
	return first
}

func (l *linkedList[T]) SeeFirst() T {
	if l.IsEmpty() {
		panic("The list is empty")
	}

	return l.first.data
}

func (l *linkedList[T]) SeeLast() T {
	if l.IsEmpty() {
		panic("The list is empty")
	}

	return l.last.data
}

func (l *linkedList[T]) Length() int {
	return l.length
}

func (l *linkedList[T]) Iterator() ListIterator[T] {
	iter := new(listIterator[T])
	iter.list = l
	iter.current = l.first
	iter.previous = nil
	return iter
}

func (l *linkedList[T]) Iterate(visit func(T) bool) {
	current := l.first
	for current != nil && visit(current.data) {
		current = current.next
	}
}

// ListIterator Primitives

func (i *listIterator[T]) SeeCurrent() T {
	if !i.HasNext() {
		panic("The iterator has finished iterating")
	}
	return i.current.data
}

func (i *listIterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *listIterator[T]) Next() T {
	if !i.HasNext() {
		panic("The iterator has finished iterating")
	}

	current := i.current.data
	i.previous = i.current
	i.current = i.current.next

	return current
}

func (i *listIterator[T]) Insert(data T) {
	newNode := i.list.createNode(data)
	newNode.next = i.current

	if i.previous == nil {
		i.list.first = newNode
		if i.current == nil {
			i.list.last = newNode
		}
	} else {
		i.previous.next = newNode
		if i.current == nil {
			i.list.last = newNode
		}
	}
	i.current = newNode

	i.list.length++
}

func (i *listIterator[T]) Delete() T {
	if !i.HasNext() {
		panic("The iterator has finished iterating")
	}

	data := i.current.data

	if i.previous == nil {
		i.list.first = i.current.next
		if i.current.next == nil {
			i.list.last = i.current.next
		}
	} else {
		i.previous.next = i.current.next
		if i.current.next == nil {
			i.list.last = i.previous
		}
	}

	i.current = i.current.next

	i.list.length--
	return data
}

func (l *linkedList[T]) createNode(data T) *nodeList[T] {
	newNode := new(nodeList[T])
	newNode.data = data
	return newNode
}

func CreateLinkedList[T any]() List[T] {
	l := new(linkedList[T])
	return l
}
