# Linked List Implementation

This project implements a ***linked-list*** data structure in **Go**. 

### Definition
```
A linked list is a linear data structure where elements are not stored at contiguous memory locations. Instead, each element (node) contains a reference (link) to the next element in the sequence.
```

## Implementation Details

- **Node Structure**: Each node in the linked list contains two fields: the data and a pointer to the next node.
- **Head Pointer**: The linked list maintains a reference to the first node (head) of the list.
- **Operations**:
  - **Insertion**: Nodes can be inserted at the beginning, end, or at a specific position in the list.
  - **Deletion**: Nodes can be deleted from the beginning, end, or from a specific position in the list.
  - **Traversal**: The list can be traversed to access each node's data.
  - **Search**: The list can be searched to find a node with specific data.

## Decision Making

- **Simplicity**: The implementation focuses on simplicity and clarity, making it easy to understand and modify.
- **Efficiency**: Common operations like insertion and deletion are designed to be efficient, with time complexities of O(1) for insertions/deletions at the beginning and O(n) for insertions/deletions at the end or specific positions.
- **Flexibility**: The implementation allows for easy extension to include additional operations or optimizations as needed.

## Usage

To use this ***linked-list*** implementation, you can import the package from the repository and create an instance of it.

## Example

Here's a simple example of how to use the linked list:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/linked_list"
)

func main() {
    list := linked_list.CreateLinkedList()

    // Insert elements
    list.InsertLast(1)
    list.InsertLast(2)
    list.InsertLast(3)

    // Traverse the list
    list.Iterate(func(data int) bool {
        fmt.Println(data)
        return true
    })

    // Search for an element
    found := list.Search(2)
    fmt.Println("Element found:", found)

    // Delete an element
    list.DeleteFirst()
    list.Iterate(func(data int) bool {
        fmt.Println(data)
        return true
    })
}
```

## Running Tests
To run the tests for this ***linked-list*** implementation, navigate to the root directory and run the following command:
```sh
go test ./linked-list
```