package union_find

type UnionFind[T comparable] interface {
    // Find returns the representative element of the set containing the given element.
    Find(T) T

    // Union merges the sets containing the two given elements.
    Union(T, T)
}
