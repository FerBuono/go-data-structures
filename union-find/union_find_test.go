package union_find_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/FerBuono/go-data-structures/union-find"
)

func TestUnionFind(t *testing.T) {
	vertices := []int{1, 2, 3, 4, 5}
	uf := union_find.NewUnionFind(vertices)

	require.Equal(t, 1, uf.Find(1))
	require.Equal(t, 2, uf.Find(2))
	require.Equal(t, 3, uf.Find(3))

	uf.Union(1, 2)
	require.Equal(t, uf.Find(1), uf.Find(2))

	uf.Union(3, 4)
	require.Equal(t, uf.Find(3), uf.Find(4))

	uf.Union(1, 3)
	require.Equal(t, uf.Find(1), uf.Find(4))
	require.Equal(t, uf.Find(2), uf.Find(3))
}
