package union_find

import (
	"github.com/FerBuono/go-data-structures/hash"
)

type unionFind[T comparable] struct {
	groups hash.Dictionary[T, T]
}

func NewUnionFind[T comparable](vertices []T) UnionFind[T] {
	u := new(unionFind[T])
	u.groups = hash.NewHash[T, T]()
	for _, vertex := range vertices {
		u.groups.Save(vertex, vertex)
	}
	return u
}

func (u *unionFind[T]) Find(vertex T) T {
	if u.groups.Get(vertex) == vertex {
		return vertex
	}

	realGroup := u.Find(u.groups.Get(vertex))
	u.groups.Save(vertex, realGroup)
	return realGroup
}

func (u *unionFind[T]) Union(v1, v2 T) {
	newGroup := u.Find(v1)
	other := u.Find(v2)
	u.groups.Save(other, newGroup)
}
