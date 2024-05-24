package linked_list

type List[T any] interface {
	// IsEmpty returns true if the list has no elements, false otherwise.
	IsEmpty() bool

	// InsertFirst adds a new element at the beginning of the list.
	InsertFirst(T)

	// InsertLast adds a new element at the end of the list.
	InsertLast(T)

	// DeleteFirst removes the first element of the list. If the list has elements, the first one is removed and its value is returned.
	// If it's empty, it panics with the message "The list is empty".
	DeleteFirst() T

	// SeeFirst returns the value of the first element of the list. If the list has elements, the value of the first one is returned.
	// If it's empty, it panics with the message "The list is empty".
	SeeFirst() T

	// SeeLast returns the value of the last element of the list. If the list has elements, the value of the last one is returned.
	// If it's empty, it panics with the message "The list is empty".
	SeeLast() T

	// Length returns the number of elements in the list. If it's empty, it returns 0.
	Length() int

	// Iterator returns an iterator for the list, which has its own primitives.
	Iterator() ListIterator[T]

	// Iterate receives a function that will be applied to the data in the list in order, until the list ends or the function returns false.
	Iterate(func(T) bool)
}

type ListIterator[T any] interface {
	// SeeCurrent returns the value of the element where the iterator is. If the iterator has already iterated all elements,
	// it panics with the message "The iterator has finished iterating".
	SeeCurrent() T

	// HasNext returns true if there is a next element, false otherwise.
	HasNext() bool

	// Next returns the value of the current element and then moves to the next one.
	// If the iterator has already iterated all elements, it panics with the message "The iterator has finished iterating".
	Next() T

	// Insert adds a new element to the list at the position where the iterator is, moving the original element at that position to the next one.
	Insert(T)

	// Delete removes the element at the current position of the iterator, linking the previous one with the next one.
	// If the iterator has already iterated all elements, it panics with the message "The iterator has finished iterating".
	Delete() T
}
