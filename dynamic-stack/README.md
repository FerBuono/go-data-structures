# Dynamic Stack Implementation

This project implements a ***dynamic-stack*** data structure in **Go**.

### Stack Definition
```
A stack is a collection of elements that follows the Last-In-First-Out (LIFO) principle. The last element added to the stack will be the first one to be removed.
```

### Implementation Details

- **Dynamic Array**: The stack uses an array that dynamically resizes to accommodate more elements or to optimize memory usage.
- **Operations**:
  - **Push**: Adds an element to the top of the stack.
  - **Pop**: Removes and returns the top element of the stack.
  - **Top**: Returns the top element without removing it.
  - **IsEmpty**: Checks if the stack is empty.

### Decision Making

- **Efficiency**: The implementation ensures that stack and unstack operations have constant time complexity, O(1), which is efficient for stack operations.

- **Flexibility**: The implementation allows for any data type to be used with the stack, making it adaptable for different use cases.

## Usage

To use this ***dynamic-stack*** implementation, you can import the package from the repository and create an instance of it.

### Example

Here's a simple example of how to use the dynamic stack:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/dynamic-stack"
)

func main() {
    s := dynamic_stack.NewDynamicStack[int]()

    // Push elements
    s.Push(1)
    s.Push(2)
    s.Push(3)

    // Peek at the top element
    fmt.Println("Top element:", s.Top())

    // Pop elements
    fmt.Println("Popped:", s.Pop())
    fmt.Println("Popped:", s.Pop())
    fmt.Println("Popped:", s.Pop())
}
```

## Running Tests
To run the tests for this ***dynamic-stack*** implementation, navigate to the root directory and run the following command:
```sh
go test ./dynamic-stack
```