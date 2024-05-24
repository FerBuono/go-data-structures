# Heap Implementation

This project implements a ***heap*** data structure in **Go**.

### Definition
```
A heap is a special type of queue in which each element is associated with a priority and is served according to its priority (highest priority first).
```

## Implementation Details

- **Heap Structure**: The heap is implemented as an array where the parent node is at index `(i - 1) / 2`, and the child nodes are at indices `2*i + 1` and `2*i + 2`.
- **Dynamic Array**: The underlying array dynamically resizes to accommodate more elements or to optimize memory usage.
- **Operations**:
  - **Enqueue**: Adds an element to the heap while maintaining the heap property.
  - **Dequeue**: Removes and returns the element with the highest priority.
  - **Peek**: Returns the element with the highest priority without removing it.
  - **IsEmpty**: Checks if the heap is empty.
  - **Size**: Returns the number of elements in the heap.

## Decision Making

- **Efficiency**: The implementation ensures that insertion and deletion operations have logarithmic time complexity, `O(log n)`, which is efficient for priority queue operations.
- **Flexibility**: The implementation allows custom comparison functions to define the priority, making it adaptable for different types and criteria.

## Usage

To use this ***heap*** implementation, you can import the package from the repository and create an instance of it.

### Example

Here's a simple example of how to use the heap:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/heap"
)

func main() {
    pq := heap.NewHeap(func(a, b int) int {
        return a - b // Max-Heap
    })

    // Enqueue elements
    pq.Enqueue(10)
    pq.Enqueue(5)
    pq.Enqueue(20)

    // Peek at the highest priority element
    fmt.Println("Top element:", pq.Peek())

    // Dequeue elements
    fmt.Println("Dequeued:", pq.Dequeue())
    fmt.Println("Dequeued:", pq.Dequeue())
    fmt.Println("Dequeued:", pq.Dequeue())
}
```

## Running Tests
To run the tests for this ***heap*** implementation, navigate to the root directory and run the following command:
```sh
go test ./heap
```