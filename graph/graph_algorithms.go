package graph

import (
    "sort"
    "github.com/FerBuono/go-data-structures/linked-queue"
    "github.com/FerBuono/go-data-structures/hash"
    "github.com/FerBuono/go-data-structures/heap"
    "github.com/FerBuono/go-data-structures/union-find"
)

func BFS[T comparable](g Graph[T]) {
    visited := hash.NewHash[T, *T]()
    parent := hash.NewHash[T, *T]()
    for _, vertex := range g.GetVertices() {
        if !visited.Contains(vertex) {
            visited.Save(vertex, nil)
            bfs(g, vertex, parent, visited)
        }
    }
}

func bfs[T comparable](g Graph[T], startVertex T, parent hash.Dictionary[T, *T], visited hash.Dictionary[T, *T]) {
    q := linked_queue.NewLinkedQueue[T]()
    q.Enqueue(startVertex)
    for !q.IsEmpty() {
        vertex := q.Dequeue()
        for _, adjacent := range g.Adjacent(vertex) {
            if !visited.Contains(adjacent) {
                visited.Save(adjacent, nil)
                parent.Save(adjacent, &vertex)
                q.Enqueue(adjacent)
            }
        }
    }
}

func DFS[T comparable](g Graph[T]) {
    visited := hash.NewHash[T, *T]()
    parent := hash.NewHash[T, *T]()
    for _, vertex := range g.GetVertices() {
        if !visited.Contains(vertex) {
            visited.Save(vertex, nil)
            dfs(g, vertex, parent, visited)
        }
    }
}

func dfs[T comparable](g Graph[T], startVertex T, parent hash.Dictionary[T, *T], visited hash.Dictionary[T, *T]) {
    parent.Save(startVertex, nil)
    for _, adjacent := range g.Adjacent(startVertex) {
        if !visited.Contains(adjacent) {
            visited.Save(adjacent, nil)
            parent.Save(adjacent, &startVertex)
            dfs(g, adjacent, parent, visited)
        }
    }
}

func IsBipartite[T comparable](g Graph[T]) bool {
    colors := hash.NewHash[T, int]()
    for _, vertex := range g.GetVertices() {
        if !colors.Contains(vertex) {
            if !isBipartite(g, vertex, colors) {
                return false
            }
        }
    }
    return true
}

func isBipartite[T comparable](g Graph[T], vertex T, colors hash.Dictionary[T, int]) bool {
    q := linked_queue.NewLinkedQueue[T]()
    q.Enqueue(vertex)
    colors.Save(vertex, 0)
    for !q.IsEmpty() {
        v := q.Dequeue()
        for _, adjacent := range g.Adjacent(v) {
            if colors.Contains(adjacent) {
                if colors.Get(adjacent) == colors.Get(v) {
                    return false
                }
            } else {
                colors.Save(adjacent, colors.Get(v)-1)
                q.Enqueue(adjacent)
            }
        }
    }
    return true
}

func TopologicalSort[T comparable](g Graph[T]) []T {
    inDegrees := hash.NewHash[T, int]()
    for _, vertex := range g.GetVertices() {
        inDegrees.Save(vertex, 0)
        for _, adjacent := range g.Adjacent(vertex) {
            if !inDegrees.Contains(adjacent) {
                inDegrees.Save(adjacent, 1)
            } else {
                inDegrees.Save(adjacent, inDegrees.Get(adjacent)+1)
            }
        }
    }

    q := linked_queue.NewLinkedQueue[T]()
    for _, vertex := range g.GetVertices() {
        if inDegrees.Get(vertex) == 0 {
            q.Enqueue(vertex)
        }
    }

    output := []T{}

    for !q.IsEmpty() {
        v := q.Dequeue()
        output = append(output, v)
        for _, adjacent := range g.Adjacent(v) {
            inDegrees.Save(adjacent, inDegrees.Get(adjacent)-1)
            if inDegrees.Get(adjacent) == 0 {
                q.Enqueue(adjacent)
            }
        }
    }
    return output
}

func ShortestPath[T comparable](source T, g Graph[T]) (hash.Dictionary[T, T], hash.Dictionary[T, int]) {
    var NONE T
    distance := hash.NewHash[T, int]()
    parent := hash.NewHash[T, T]()
    visited := hash.NewHash[T, bool]()

    for _, vertex := range g.GetVertices() {
        distance.Save(vertex, int(^uint(0)>>1))
    }

    distance.Save(source, 0)
    parent.Save(source, NONE)
    visited.Save(source, true)

    q := linked_queue.NewLinkedQueue[T]()
    q.Enqueue(source)

    for !q.IsEmpty() {
        v := q.Dequeue()
        for _, adjacent := range g.Adjacent(v) {
            if !visited.Contains(adjacent) {
                distance.Save(adjacent, distance.Get(v)+1)
                parent.Save(adjacent, v)
                visited.Save(adjacent, true)
                q.Enqueue(adjacent)
            }
        }
    }
    return parent, distance
}

type dist[T comparable] struct {
    vertex T
    weight int
}

func ShortestPathDijkstra[T comparable](source T, g Graph[T]) (hash.Dictionary[T, T], hash.Dictionary[T, int]) {
    var NONE T
    distance := hash.NewHash[T, int]()
    parent := hash.NewHash[T, T]()

    for _, vertex := range g.GetVertices() {
        distance.Save(vertex, int(^uint(0)>>1))
    }

    distance.Save(source, 0)
    parent.Save(source, NONE)

    h := heap.NewHeap(func(a, b dist[T]) int { return b.weight - a.weight })
    h.Enqueue(dist[T]{source, 0})

    for !h.IsEmpty() {
        v := h.Dequeue().vertex
        for _, adjacent := range g.Adjacent(v) {
            if distance.Get(v)+g.Weight(v, adjacent) < distance.Get(adjacent) {
                distance.Save(adjacent, distance.Get(v)+g.Weight(v, adjacent))
                parent.Save(adjacent, v)
                h.Enqueue(dist[T]{adjacent, distance.Get(adjacent)})
            }
        }
    }

    result := []dist[T]{}
    for iter := distance.Iterator(); iter.HasNext(); {
        v, p := iter.Current()
        result = append(result, dist[T]{v, p})
        iter.Next()
    }
    return parent, distance
}

func Centrality[T comparable](g Graph[T]) []dist[T] {
    cent := hash.NewHash[T, int]()
    for _, vertex := range g.GetVertices() {
        cent.Save(vertex, 0)
    }
    for _, v := range g.GetVertices() {
        parent, _ := ShortestPath(v, g)
        for _, w := range g.GetVertices() {
            if v == w {
                continue
            }
            if !parent.Contains(w) {
                continue
            }
            current := parent.Get(w)
            for current != v {
                cent.Save(current, cent.Get(current)+1)
                current = parent.Get(current)
            }
        }
    }
    result := []dist[T]{}
    for iter := cent.Iterator(); iter.HasNext(); {
        v, p := iter.Current()
        result = append(result, dist[T]{v, p / 2})
        iter.Next()
    }
    return result
}

func MinInversions[T comparable](g Graph[T], s, t T) int {
    weightedGraph := NewGraph(true, []T{})
    for _, vertex := range g.GetVertices() {
        if !weightedGraph.Contains(vertex) {
            weightedGraph.AddVertex(vertex)
        }
        for _, adjacent := range g.Adjacent(vertex) {
            if !weightedGraph.Contains(adjacent) {
                weightedGraph.AddVertex(adjacent)
            }
            weightedGraph.AddEdge(vertex, adjacent, 0)
            if !g.ContainsEdge(adjacent, vertex) {
                weightedGraph.AddEdge(adjacent, vertex, 1)
            }
        }
    }
    _, path := ShortestPathDijkstra(s, weightedGraph)
    return path.Get(t)
}

type edge[T comparable] struct {
    source T
    target T
    weight int
}

func MSTPrim[T comparable](g Graph[T]) Graph[T] {
    source := g.RandomVertex()
    visited := hash.NewHash[T, bool]()
    visited.Save(source, true)

    h := heap.NewHeap(func(a, b edge[T]) int { return b.weight - a.weight })
    for _, adjacent := range g.Adjacent(source) {
        h.Enqueue(edge[T]{source, adjacent, g.Weight(source, adjacent)})
    }

    mst := NewGraph[T](false, g.GetVertices())
    for _, vertex := range g.GetVertices() {
        mst.AddVertex(vertex)
    }

    for !h.IsEmpty() {
        e := h.Dequeue()
        if visited.Contains(e.target) {
            continue
        }
        mst.AddEdge(e.source, e.target, e.weight)
        visited.Save(e.target, true)
        for _, adjacent := range g.Adjacent(e.target) {
            if !visited.Contains(adjacent) {
                h.Enqueue(edge[T]{e.target, adjacent, g.Weight(e.target, adjacent)})
            }
        }
    }
    return mst
}

func GetEdges[T comparable](g Graph[T]) []edge[T] {
    edges := []edge[T]{}
    visited := hash.NewHash[T, bool]()
    for _, vertex := range g.GetVertices() {
        for _, adjacent := range g.Adjacent(vertex) {
            if !visited.Contains(adjacent) {
                edges = append(edges, edge[T]{vertex, adjacent, g.Weight(vertex, adjacent)})
            }
        }
        visited.Save(vertex, true)
    }
    return edges
}

func MSTKruskal[T comparable](g Graph[T]) Graph[T] {
    sets := union_find.NewUnionFind(g.GetVertices())
    edges := GetEdges(g)
    sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
    mst := NewGraph[T](false, g.GetVertices())
    for _, e := range edges {
        if sets.Find(e.source) == sets.Find(e.target) {
            continue
        }
        mst.AddEdge(e.source, e.target, e.weight)
        sets.Union(e.source, e.target)
    }
    return mst
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func dfsArticulationPoints[T comparable](g Graph[T], v T, visited hash.Dictionary[T, bool], parent hash.Dictionary[T, T], order hash.Dictionary[T, int], low hash.Dictionary[T, int], points hash.Dictionary[T, T], isRoot bool) {
    children := 0
    low.Save(v, order.Get(v))
    for _, w := range g.Adjacent(v) {
        if !visited.Contains(w) {
            children++
            order.Save(w, order.Get(v)+1)
            parent.Save(w, v)
            visited.Save(w, true)
            dfsArticulationPoints(g, w, visited, parent, order, low, points, false)

            if low.Get(w) >= order.Get(v) && !isRoot {
                points.Save(v, v)
            }

            low.Save(v, min(low.Get(v), low.Get(w)))
        } else if parent.Get(v) != w {
            low.Save(v, min(low.Get(v), order.Get(w)))
        }
    }

    if isRoot && children > 1 {
        points.Save(v, v)
    }
}

func ArticulationPoints[T comparable](g Graph[T]) hash.Dictionary[T, T] {
    var NONE T
    source := g.RandomVertex()
    visited := hash.NewHash[T, bool]()
    parent := hash.NewHash[T, T]()
    order := hash.NewHash[T, int]()
    low := hash.NewHash[T, int]()
    articulationPoints := hash.NewHash[T, T]()

    visited.Save(source, true)
    parent.Save(source, NONE)
    order.Save(source, 0)

    dfsArticulationPoints(g, source, visited, parent, order, low, articulationPoints, true)

    return articulationPoints
}
