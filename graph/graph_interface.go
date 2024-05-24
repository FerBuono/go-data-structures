package graph

type Graph[T comparable] interface {

    // AddVertex adds a new vertex to the graph.
    AddVertex(T)

    // RemoveVertex removes a vertex from the graph.
    RemoveVertex(T)

    // AddEdge adds an edge between two vertices with a specified weight.
    AddEdge(T, T, int)

    // RemoveEdge removes the edge between two vertices.
    RemoveEdge(T, T)

    // Weight returns the weight of the edge between two vertices.
    Weight(T, T) int

    // Contains checks if a vertex is in the graph.
    Contains(T) bool

    // GetVertices returns a slice of all vertices in the graph.
    GetVertices() []T

    // Adjacent returns a slice of all vertices adjacent to a given vertex.
    Adjacent(T) []T

    // ContainsEdge checks if an edge exists between two vertices.
    ContainsEdge(T, T) bool

    // RandomVertex returns a random vertex from the graph.
    RandomVertex() T

    // Iterator returns an iterator for the graph.
    Iterator() GraphIterator[T]
}

type GraphIterator[T comparable] interface {

    // Current returns the current vertex.
    Current() T

    // HasNext checks if there are more vertices to iterate over.
    HasNext() bool

    // Next moves the iterator to the next vertex and returns it.
    Next() T
}
