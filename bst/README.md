# Binary Search Tree (BST) Implementation

This project implements a ***binary search tree*** (BST) data structure in Go.

### BST Definition
```
A binary search tree is a hierarchical data structure in which each node has at most two children referred to as the left child and the right child. For each node, the left subtree contains only nodes with keys less than the node's key, and the right subtree contains only nodes with keys greater than the node's key.
```

### Implementation Details

- **Tree Structure**: The tree is composed of nodes, each containing a key and a value, along with references to the left and right children.
- **Operations**:
  - **Insert**: Adds a key-value pair to the tree.
  - **Delete**: Removes a key-value pair from the tree.
  - **Search**: Retrieves the value associated with a given key.
  - **Traversal**: Allows iteration over the elements in the tree.
  - **Range Iteration**: Iterates over a specific range of keys.
- **Iterators**:
  - **Standard Iterator**: Iterates over all elements in the tree.
  - **Range Iterator**: Iterates over elements within a specified range of keys.

### Decision Making

- **Efficiency**:
  - **Insert**: O(log n) on average, O(n) in the worst case.
  - **Delete**: O(log n) on average, O(n) in the worst case.
  - **Search**: O(log n) on average, O(n) in the worst case.
  - **Traversal**: O(n) - Visiting each node once.
- **Flexibility**: The implementation supports generic types for both keys and values, making it adaptable for various use cases.

## Usage

To use this ***binary search tree*** implementation, you can import the package from the repository and create an instance of it.


### Example

Here's a simple example of how to use the binary search tree:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/bst"
)

func main() {
    tree := bst.NewBST(func(a, b int) int {
        return a - b
    })

    // Insert elements
    tree.Insert(10, "ten")
    tree.Insert(5, "five")
    tree.Insert(20, "twenty")

    // Search for an element
    fmt.Println("Value for key 10:", tree.Search(10))

    // Delete an element
    tree.Delete(10)

    // Traverse the tree
    tree.Iterate(func(key int, value string) bool {
        fmt.Printf("Key: %d, Value: %s\n", key, value)
        return true
    })
}
```

## Running Tests
To run the tests for this ***binary search tree*** implementation, navigate to the root directory and run the following command:
```sh
go test ./bst
```