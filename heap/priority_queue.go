package heap

type PriorityQueue[T comparable] interface {

    // IsEmpty returns true if the queue is empty, false otherwise.
    IsEmpty() bool

    // Enqueue adds an element to the heap.
    Enqueue(T)

    // Peek returns the element with the highest priority. If empty, it panics with the message "The queue is empty".
    Peek() T

    // Dequeue removes the element with the highest priority and returns it. If empty, it panics with the message "The queue is empty".
    Dequeue() T

    // Size returns the number of elements in the priority queue.
    Size() int
}