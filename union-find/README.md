# Union-Find Implementation

This project implements a ***union-find*** data structure in Go.

### Definition
```
Union-Find is a data structure that keeps track of a set of elements partitioned into disjoint (non-overlapping) subsets.
```

## Implementation Details

- **Union-Find Structure**: The Union-Find structure is implemented using a hash map to store the parent of each element.
- **Operations**:
  - **Find**: Returns the representative of the set containing the element.
  - **Union**: Merges two sets containing the two elements.

## Decision Making

- **Efficiency**: The implementation ensures that both `Find` and `Union` operations have nearly constant time complexity, `O(α(n))`, where `α` is the inverse Ackermann function, which grows very slowly.
- **Flexibility**: The implementation supports generic types, making it adaptable for different data types.

## Usage

To use this ***union-find*** implementation, you can import the package from the repository and create an instance of it.

### Example

Here's a simple example of how to use the Union-Find:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/union-find"
)

func main() {
    vertices := []int{1, 2, 3, 4, 5}
    uf := union_find.NewUnionFind(vertices)

    fmt.Println("Find 1:", uf.Find(1))
    fmt.Println("Find 2:", uf.Find(2))

    uf.Union(1, 2)
    fmt.Println("After Union(1, 2):")
    fmt.Println("Find 1:", uf.Find(1))
    fmt.Println("Find 2:", uf.Find(2))
}
```

## Running Tests
To run the tests for this ***union-find*** implementation, navigate to the root directory and run the following command:
```sh
go test ./union-find
```