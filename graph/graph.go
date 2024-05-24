package graph

import (
    "math/rand"
    "github.com/FerBuono/go-data-structures/hash"
)

type graph[T comparable] struct {
    dicc     hash.Dictionary[T, hash.Dictionary[T, int]]
    directed bool
}

type graphIterator[T comparable] struct {
    iterDicc hash.DictionaryIterator[T, hash.Dictionary[T, int]]
}

func NewGraph[T comparable](directed bool, vertices []T) Graph[T] {
    g := new(graph[T])
    g.dicc = hash.NewHash[T, hash.Dictionary[T, int]]()
    for _, vertex := range vertices {
        g.dicc.Save(vertex, hash.NewHash[T, int]())
    }
    g.directed = directed
    return g
}

func (g *graph[T]) AddVertex(v T) {
    g.dicc.Save(v, hash.NewHash[T, int]())
}

func (g *graph[T]) RemoveVertex(v T) {
    if !g.Contains(v) {
        panic("The vertex does not belong to the graph")
    }
    g.dicc.Delete(v)
    for iter := g.dicc.Iterator(); iter.HasNext(); {
        _, adjDict := iter.Current()
        if adjDict.Contains(v) {
            adjDict.Delete(v)
        }
        iter.Next()
    }
}

func (g *graph[T]) AddEdge(v1, v2 T, weight int) {
    if !g.Contains(v1) || !g.Contains(v2) {
        panic("A vertex does not belong to the graph")
    }
    adjDict := g.dicc.Get(v1)
    adjDict.Save(v2, weight)
    if !g.directed {
        adjDict = g.dicc.Get(v2)
        adjDict.Save(v1, weight)
    }
}

func (g *graph[T]) RemoveEdge(v1, v2 T) {
    if !g.Contains(v1) || !g.Contains(v2) {
        panic("A vertex does not belong to the graph")
    }
    adjDict := g.dicc.Get(v1)
    adjDict.Delete(v2)
    if !g.directed {
        adjDict = g.dicc.Get(v2)
        adjDict.Delete(v1)
    }
}

func (g *graph[T]) Weight(v1, v2 T) int {
    if !g.Contains(v1) || !g.Contains(v2) {
        panic("A vertex does not belong to the graph")
    }
    return g.dicc.Get(v1).Get(v2)
}

func (g *graph[T]) Contains(v T) bool {
    return g.dicc.Contains(v)
}

func (g *graph[T]) GetVertices() []T {
    vertices := []T{}
    for iter := g.dicc.Iterator(); iter.HasNext(); {
        vertices = append(vertices, iter.Next())
    }
    return vertices
}

func (g *graph[T]) Adjacent(v T) []T {
    adj := []T{}
    adjDict := g.dicc.Get(v)
    for iter := adjDict.Iterator(); iter.HasNext(); {
        adj = append(adj, iter.Next())
    }
    return adj
}

func (g *graph[T]) RandomVertex() T {
    vertices := g.GetVertices()
    return vertices[rand.Intn(len(vertices))]
}

func (g *graph[T]) ContainsEdge(v1, v2 T) bool {
    if !g.Contains(v1) || !g.Contains(v2) {
        return false
    }
    return g.dicc.Get(v1).Contains(v2)
}

func (g *graph[T]) Iterator() GraphIterator[T] {
    iter := new(graphIterator[T])
    iter.iterDicc = g.dicc.Iterator()
    return iter
}

func (i *graphIterator[T]) Current() T {
    vertex, _ := i.iterDicc.Current()
    return vertex
}

func (i *graphIterator[T]) HasNext() bool {
    return i.iterDicc.HasNext()
}

func (i *graphIterator[T]) Next() T {
    return i.iterDicc.Next()
}
