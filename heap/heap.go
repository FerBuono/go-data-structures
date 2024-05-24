package heap

const initialCapacity = 10
const increaseFactor = 2
const decreaseFactor = 2
const reduceThreshold = 4


type heap[T comparable] struct {
    data     []T
    count    int
    compare  func(T, T) int
}

func NewHeap[T comparable](compare func(T, T) int) PriorityQueue[T] {
    h := &heap[T]{
        data:    make([]T, initialCapacity),
        compare: compare,
    }
    return h
}

func NewHeapFromArray[T comparable](array []T, compare func(T, T) int) PriorityQueue[T] {
    h := &heap[T]{
        data:    make([]T, max(initialCapacity, len(array))),
        count:   len(array),
        compare: compare,
    }
    copy(h.data, array)
    heapify(h.data, compare)
    return h
}

func HeapSort[T comparable](elements []T, compare func(T, T) int) {
    heapify(elements, compare)
    for i := 0; i < len(elements); i++ {
        swap(&elements[0], &elements[len(elements)-1-i])
        downheap(elements[:len(elements)-1-i], 0, compare, len(elements)-1-i)
    }
}

// PriorityQueue methods

func (h *heap[T]) IsEmpty() bool {
    return h.count == 0
}

func (h *heap[T]) Enqueue(element T) {
    if h.count == cap(h.data) {
        h.resize(cap(h.data) * increaseFactor)
    }
    h.data[h.count] = element
    h.count++
    upheap(h.data, h.count-1, h.compare)
}

func (h *heap[T]) Peek() T {
    if h.IsEmpty() {
        panic("The queue is empty")
    }
    return h.data[0]
}

func (h *heap[T]) Dequeue() T {
    if h.IsEmpty() {
        panic("The queue is empty")
    }
    if h.count <= cap(h.data)/reduceThreshold && cap(h.data) > initialCapacity {
        h.resize(cap(h.data) / decreaseFactor)
    }
    item := h.data[0]
    swap(&h.data[0], &h.data[h.count-1])
    h.count--
    downheap(h.data, 0, h.compare, h.count)
    return item
}

func (h *heap[T]) Size() int {
    return h.count
}

// Auxiliary methods/functions

func upheap[T comparable](data []T, childIndex int, compare func(T, T) int) {
    if childIndex <= 0 {
        return
    }
    parentIndex := (childIndex - 1) / 2
    if compare(data[parentIndex], data[childIndex]) < 0 {
        swap(&data[parentIndex], &data[childIndex])
        upheap(data, parentIndex, compare)
    }
}

func downheap[T comparable](data []T, parentIndex int, compare func(T, T) int, count int) {
    if parentIndex >= count {
        return
    }
    leftChildIndex := 2*parentIndex + 1
    rightChildIndex := 2*parentIndex + 2
    if leftChildIndex >= count && rightChildIndex >= count {
        return
    }
    var replacementIndex int
    if rightChildIndex >= count {
        replacementIndex = findReplacement(data, parentIndex, leftChildIndex, leftChildIndex, compare)
    } else {
        replacementIndex = findReplacement(data, parentIndex, leftChildIndex, rightChildIndex, compare)
    }
    if replacementIndex >= count {
        return
    }
    if replacementIndex != parentIndex {
        swap(&data[parentIndex], &data[replacementIndex])
        downheap(data, replacementIndex, compare, count)
    }
}

func heapify[T comparable](arr []T, compare func(T, T) int) {
    for i := len(arr) - 1; i >= 0; i-- {
        downheap(arr, i, compare, len(arr))
    }
}

func findReplacement[T comparable](data []T, parentIndex, leftChildIndex, rightChildIndex int, compare func(T, T) int) int {
    if compare(data[rightChildIndex], data[parentIndex]) > 0 && compare(data[rightChildIndex], data[leftChildIndex]) >= 0 {
        return rightChildIndex
    } else if compare(data[leftChildIndex], data[parentIndex]) > 0 && compare(data[leftChildIndex], data[rightChildIndex]) > 0 {
        return leftChildIndex
    }
    return parentIndex
}

func (h *heap[T]) resize(newCapacity int) {
    newData := make([]T, newCapacity)
    copy(newData, h.data)
    h.data = newData
}

func swap[T comparable](x, y *T) {
    *x, *y = *y, *x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
