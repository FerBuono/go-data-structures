# Hash Table Implementation

This project implements a ***hash table*** data structure in **Go**.

### Definition
```
A hash table is a data structure that implements an associative array abstract data type, a structure that can map keys to values.
```

## Implementation Details

- **Hash Table Structure**: The hash table is implemented as an array of elements where each element can be empty, occupied, or deleted.
- **Dynamic Array**: The underlying array dynamically resizes to accommodate more elements or to optimize memory usage.
- **Operations**:
  - **Save**: Adds an element to the hash table or updates an existing element.
  - **Contains**: Checks if a key is in the hash table.
  - **Get**: Retrieves the value associated with a key.
  - **Delete**: Removes an element from the hash table.
  - **Size**: Returns the number of elements in the hash table.
  - **Iterate**: Iterates over all elements in the hash table.
  - **Iterator**: Returns an iterator for the hash table.

## Decision Making

- **Efficiency**: The implementation ensures that most operations have constant time complexity, `O(1)`, on average, which is efficient for dictionary operations.
- **Flexibility**: The implementation supports generic types for keys and values, making it adaptable for different data types.

## Usage

To use this ***hash table*** implementation, you can import the package from the repository and create an instance of it.

### Example

Here's a simple example of how to use the hash table:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/hash"
)

func main() {
    dict := hash.NewHash[string, int]()

    // Save elements
    dict.Save("key1", 1)
    dict.Save("key2", 2)
    dict.Save("key3", 3)

    // Get elements
    fmt.Println("key1:", dict.Get("key1"))
    fmt.Println("key2:", dict.Get("key2"))
    fmt.Println("key3:", dict.Get("key3"))

    // Delete elements
    fmt.Println("Deleted key1:", dict.Delete("key1"))
    fmt.Println("Deleted key2:", dict.Delete("key2"))
    fmt.Println("Deleted key3:", dict.Delete("key3"))
}
```
## Running Tests
To run the tests for this ***hash table*** implementation, navigate to the root directory and run the following command:
```sh
go test ./hash
```