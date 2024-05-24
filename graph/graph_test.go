package graph_test

import (
    "testing"
    "fmt"
    "github.com/stretchr/testify/require"
    "github.com/FerBuono/go-data-structures/graph"
)

func TestGraph(t *testing.T) {
    vertices := []int{1, 2, 3, 4, 5}
    g := graph.NewGraph[int](false, vertices)

    require.True(t, g.Contains(1))
    require.True(t, g.Contains(2))
    require.False(t, g.Contains(6))

    g.AddEdge(1, 2, 10)
    require.True(t, g.ContainsEdge(1, 2))
    require.True(t, g.ContainsEdge(2, 1))
    require.Equal(t, 10, g.Weight(1, 2))
    require.Equal(t, 10, g.Weight(2, 1))

    g.RemoveEdge(1, 2)
    require.False(t, g.ContainsEdge(1, 2))
    require.False(t, g.ContainsEdge(2, 1))

    g.AddVertex(6)
    require.True(t, g.Contains(6))

    g.RemoveVertex(6)
    require.False(t, g.Contains(6))

    g.AddEdge(1, 3, 5)
    g.AddEdge(1, 4, 15)
    require.Equal(t, []int{3, 4}, g.Adjacent(1))
}

func TestDirectedGraph(t *testing.T) {
    vertices := []int{1, 2, 3, 4, 5}
    g := graph.NewGraph[int](true, vertices)

    g.AddEdge(1, 2, 10)
    require.True(t, g.ContainsEdge(1, 2))
    require.False(t, g.ContainsEdge(2, 1))

    require.Equal(t, 10, g.Weight(1, 2))
    require.Panics(t, func() { g.Weight(2, 1) })

    g.RemoveEdge(1, 2)
    require.False(t, g.ContainsEdge(1, 2))
}

func TestRandomVertex(t *testing.T) {
    vertices := []int{1, 2, 3, 4, 5}
    g := graph.NewGraph[int](false, vertices)
    vertex := g.RandomVertex()
    require.Contains(t, vertices, vertex)
}

func TestGraphOperations(t *testing.T) {
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

    require.Equal(t, 1, g.Weight("A", "B"))
    require.Equal(t, 2, g.Weight("A", "C"))
    require.Equal(t, 3, g.Weight("B", "D"))
    require.Equal(t, 4, g.Weight("C", "E"))
    require.Equal(t, 5, g.Weight("D", "F"))
    require.Equal(t, 6, g.Weight("E", "G"))
    require.Equal(t, 7, g.Weight("F", "H"))

    require.True(t, g.Contains("A"))
    require.False(t, g.Contains("Z"))

    require.True(t, g.ContainsEdge("A", "B"))
    require.False(t, g.ContainsEdge("A", "Z"))

    vertices = g.GetVertices()
    require.Contains(t, vertices, "A")
    require.Contains(t, vertices, "B")
    require.Contains(t, vertices, "C")
    require.Contains(t, vertices, "D")
    require.Contains(t, vertices, "E")
    require.Contains(t, vertices, "F")
    require.Contains(t, vertices, "G")
    require.Contains(t, vertices, "H")

    adjacent := g.Adjacent("A")
    require.Contains(t, adjacent, "B")
    require.Contains(t, adjacent, "C")

    g.RemoveEdge("A", "B")
    require.False(t, g.ContainsEdge("A", "B"))

    g.RemoveVertex("A")
    require.False(t, g.Contains("A"))
}

func TestGraphAlgorithms(t *testing.T) {
    vertices := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    g := graph.NewGraph[string](false, vertices)
    g.AddEdge("A", "B", 1)
    g.AddEdge("A", "C", 2)
    g.AddEdge("B", "D", 3)
    g.AddEdge("C", "E", 4)
    g.AddEdge("D", "F", 5)
    g.AddEdge("E", "G", 6)
    g.AddEdge("F", "H", 7)

    fmt.Println("Testing BFS")
    graph.BFS(g)

    fmt.Println("Testing DFS")
    graph.DFS(g)

    fmt.Println("Testing IsBipartite")
    require.True(t, graph.IsBipartite(g))

    fmt.Println("Testing TopologicalSort")
    sorted := graph.TopologicalSort(g)
    fmt.Println("Topologically sorted vertices:", sorted)

    fmt.Println("Testing ShortestPath")
    parent, distance := graph.ShortestPath("A", g)
    fmt.Println("Parent map:", parent)
    fmt.Println("Distance map:", distance)

    fmt.Println("Testing ShortestPathDijkstra")
    parent, distance = graph.ShortestPathDijkstra("A", g)
    fmt.Println("Parent map:", parent)
    fmt.Println("Distance map:", distance)

    fmt.Println("Testing Centrality")
    centrality := graph.Centrality(g)
    fmt.Println("Centrality:", centrality)

    fmt.Println("Testing MinInversions")
    inversions := graph.MinInversions(g, "A", "H")
    fmt.Println("Minimum inversions:", inversions)

    fmt.Println("Testing MSTPrim")
    mst := graph.MSTPrim(g)
    fmt.Println("MST (Prim):", mst)

    fmt.Println("Testing MSTKruskal")
    mst = graph.MSTKruskal(g)
    fmt.Println("MST (Kruskal):", mst)

    fmt.Println("Testing ArticulationPoints")
    articulationPoints := graph.ArticulationPoints(g)
    fmt.Println("Articulation points:", articulationPoints)
}
