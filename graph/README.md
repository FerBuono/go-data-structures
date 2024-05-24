# Graph Implementation
This project implements a ***graph*** data structure in **Go**. It also includes a variety of algorithms to test the implementation effectively.

### Definition
```
A graph is a collection of nodes (vertices) and edges that connect pairs of nodes. Graphs can be directed or undirected, weighted or unweighted.
```

## Implementation Details
- **Graph Structure**: The graph is implemented using an adjacency list where each vertex maps to a dictionary of adjacent vertices and their edge weights.
- **Directed and Undirected Graphs**: The implementation supports both directed and undirected graphs.
- **Dynamic Vertices and Edges**: Vertices and edges can be added or removed dynamically.
- **Operations**:
  - **Add Vertex**: Adds a new vertex to the graph.
  - **Remove Vertex**: Removes a vertex and all its associated edges from the graph.
  - **Add Edge**: Adds a new edge between two vertices with a specified weight.
  - **Remove Edge**: Removes the edge between two vertices.
  - **Get Weight**: Retrieves the weight of the edge between two vertices.
  - **Contains Vertex**: Checks if a vertex exists in the graph.
  - **Get Vertices**: Returns a list of all vertices in the graph.
  - **Get Adjacent** Vertices: Returns a list of vertices adjacent to a given vertex.
  - **Contains Edge**: Checks if an edge exists between two vertices.
  - **Random Vertex**: Returns a random vertex from the graph.

## Decision Making
- **Efficiency**: 
  - **Add Vertex**: O(1) - Adding a vertex involves inserting a key in the adjacency list dictionary.
  - **Remove Vertex**: O(V + E) - Removing a vertex requires removing it from the adjacency list and updating all edges.
  - **Add Edge**: O(1) - Adding an edge involves inserting an entry in the adjacency list.
  - **Remove Edge**: O(1) - Removing an edge involves removing an entry from the adjacency list.
  - **Get Weight**: O(1) - Retrieving the weight of an edge involves looking up the edge in the adjacency list.
  - **Contains Vertex**: O(1) - Checking for a vertex involves a dictionary lookup.
  - **Get Vertices**: O(V) - Retrieving all vertices involves iterating through the dictionary keys.
  - **Get Adjacent Vertices**: O(1) - Retrieving adjacent vertices involves a dictionary lookup.
  - **Contains Edge**: O(1) - Checking for an edge involves a dictionary lookup.
  - **Random Vertex**: O(1) - Retrieving a random vertex involves accessing a random element in the dictionary.
- **Flexibility**: The implementation supports custom data types for vertices and weights, making it adaptable for different use cases.

## Graph Algorithms
- **BFS (Breadth-First Search)**: Traverses the graph level by level from a starting vertex.
- **DFS (Depth-First Search)**: Traverses the graph depth-wise from a starting vertex.
- **Is Bipartite**: Checks if the graph can be colored with two colors such that no two adjacent vertices share the same color.
- **Topological Sort**: Produces a linear ordering of vertices for directed acyclic graphs.
- **Shortest Path (Unweighted)**: Finds the shortest path from a source vertex to all other vertices.
- **Shortest Path (Dijkstra)**: Finds the shortest path in a weighted graph using Dijkstra's algorithm.
- **Centrality**: Computes the centrality of each vertex in the graph.
- **Min Inversions**: Computes the minimum number of edge reversals needed to make a directed path from one vertex to another.
- **MST (Minimum Spanning Tree) - Prim's Algorithm**: Constructs a minimum spanning tree using Prim's algorithm.
- **MST (Minimum Spanning Tree) - Kruskal's Algorithm**: Constructs a minimum spanning tree using Kruskal's algorithm.
- **Articulation Points**: Finds the articulation points (cut vertices) in the graph.

## Usage

To use this ***graph*** implementation, you can import the package from the repository and create an instance of it.

### Example

Here's a simple example of how to use the graph:

```go
package main

import (
    "fmt"
    "github.com/FerBuono/go-data-structures/graph"
)

func main() {
    vertices := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    g := graph.NewGraph[string](false, vertices)

    // Add edges
    g.AddEdge("A", "B", 1)
    g.AddEdge("A", "C", 2)
    g.AddEdge("B", "D", 3)
    g.AddEdge("C", "E", 4)
    g.AddEdge("D", "F", 5)
    g.AddEdge("E", "G", 6)
    g.AddEdge("F", "H", 7)

    // Perform BFS
    fmt.Println("Performing BFS")
    graph.BFS(g)

    // Perform DFS
    fmt.Println("Performing DFS")
    graph.DFS(g)

    // Check if the graph is bipartite
    fmt.Println("Is the graph bipartite?", graph.IsBipartite(g))

    // Perform topological sort
    fmt.Println("Topological sort:", graph.TopologicalSort(g))

    // Find shortest paths
    fmt.Println("Shortest paths from A:", graph.ShortestPath("A", g))

    // Perform Dijkstra's algorithm
    fmt.Println("Dijkstra's shortest paths from A:", graph.ShortestPathDijkstra("A", g))

    // Calculate centrality
    fmt.Println("Centrality of vertices:", graph.Centrality(g))

    // Find minimum inversions
    fmt.Println("Minimum inversions from A to H:", graph.MinInversions(g, "A", "H"))

    // Construct MST using Prim's algorithm
    fmt.Println("MST using Prim's algorithm:", graph.MSTPrim(g))

    // Construct MST using Kruskal's algorithm
    fmt.Println("MST using Kruskal's algorithm:", graph.MSTKruskal(g))

    // Find articulation points
    fmt.Println("Articulation points:", graph.ArticulationPoints(g))
}

```
## Running Tests
To run the tests for this ***graph*** implementation, navigate to the root directory and run the following command:
```sh
go test ./graph
```