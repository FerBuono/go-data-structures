package linked_queue

type nodeQueue[T any] struct {
    value T
    next  *nodeQueue[T]
}

type linkedQueue[T any] struct {
    first *nodeQueue[T]
    last  *nodeQueue[T]
}

func (q *linkedQueue[T]) IsEmpty() bool {
    return q.first == nil && q.last == nil
}

func (q *linkedQueue[T]) Peek() T {
    if q.IsEmpty() {
        panic("The queue is empty")
    }
    return q.first.value
}

func (q *linkedQueue[T]) Enqueue(value T) {
    newNode := q.createNode(value)

    if q.IsEmpty() {
        q.first = newNode
    } else {
        q.last.next = newNode
    }

    q.last = newNode
}

func (q *linkedQueue[T]) Dequeue() T {
    if q.IsEmpty() {
        panic("The queue is empty")
    }

    value := q.first.value

    if q.first.next == nil {
        q.last = nil
    }
    q.first = q.first.next

    return value
}

func (q *linkedQueue[T]) createNode(value T) *nodeQueue[T] {
    newNode := new(nodeQueue[T])
    newNode.value = value
    return newNode
}

func NewLinkedQueue[T any]() Queue[T] {
    q := new(linkedQueue[T])
    return q
}
