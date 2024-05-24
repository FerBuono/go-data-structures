# Linked Queue Implementation

This project implements a ***linked-queue*** data structure in **Go**.

### Definition
```
A queue is a linear data structure that follows the First-In-First-Out (FIFO) principle. The first element added to the queue will be the first one to be removed.
```

## Implementation Details

- **Linked List Structure**: The queue is implemented as a linked list where each node contains a value and a pointer to the next node.
- **Operations**:
  - **IsEmpty**: Checks if the queue is empty.
  - **Peek**: Returns the value of the first element without removing it.
  - **Enqueue**: Adds a new element to the end of the queue.
  - **Dequeue**: Removes and returns the first element of the queue.

## Decision Making

- **Efficiency**: The implementation ensures that enqueue and dequeue operations have constant time complexity, `O(1)`, which is efficient for queue operations.
- **Flexibility**: The implementation supports generic types, making it adaptable for different data types.

## Usage

To use this ***linked-queue*** implementation, you can import the package from the repository and create an instance of it.

### Example

Here's a simple example of how to use the queue:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/linked-queue"
)

func main() {
    q := queue.NewLinkedQueue[int]()

    // Enqueue elements
    q.Enqueue(10)
    q.Enqueue(20)
    q.Enqueue(30)

    // Peek at the first element
    fmt.Println("First element:", q.Peek())

    // Dequeue elements
    fmt.Println("Dequeued:", q.Dequeue())
    fmt.Println("Dequeued:", q.Dequeue())
    fmt.Println("Dequeued:", q.Dequeue())
}
```

## Running Tests
To run the tests for this ***linked-queue*** implementation, navigate to the root directory and run the following command:
```sh
go test ./linked-queue
```