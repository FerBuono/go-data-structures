package hash_test

import (
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/FerBuono/go-data-structures/hash"
)

func TestEmptyHash(t *testing.T) {
    dicc := hash.NewHash[string, int]()

    require.Equal(t, 0, dicc.Size())
    require.Panics(t, func() { dicc.Get("key1") })
    require.Panics(t, func() { dicc.Delete("key1") })
}

func TestSaveAndGet(t *testing.T) {
    dicc := hash.NewHash[string, int]()
    dicc.Save("key1", 1)
    dicc.Save("key2", 2)
    dicc.Save("key3", 3)

    require.Equal(t, 1, dicc.Get("key1"))
    require.Equal(t, 2, dicc.Get("key2"))
    require.Equal(t, 3, dicc.Get("key3"))
}

func TestContains(t *testing.T) {
    dicc := hash.NewHash[string, int]()
    dicc.Save("key1", 1)
    dicc.Save("key2", 2)

    require.True(t, dicc.Contains("key1"))
    require.True(t, dicc.Contains("key2"))
    require.False(t, dicc.Contains("key3"))
}

func TestDelete(t *testing.T) {
    dicc := hash.NewHash[string, int]()
    dicc.Save("key1", 1)
    dicc.Save("key2", 2)

    require.Equal(t, 1, dicc.Delete("key1"))
    require.Panics(t, func() { dicc.Get("key1") })
    require.Equal(t, 2, dicc.Delete("key2"))
    require.Panics(t, func() { dicc.Get("key2") })
}

func TestResize(t *testing.T) {
    dicc := hash.NewHash[int, int]()
    for i := 0; i < 1000; i++ {
        dicc.Save(i, i)
    }
    require.Equal(t, 1000, dicc.Size())

    for i := 0; i < 1000; i++ {
        require.Equal(t, i, dicc.Get(i))
    }
}

func TestIterate(t *testing.T) {
    dicc := hash.NewHash[string, int]()
    dicc.Save("key1", 1)
    dicc.Save("key2", 2)
    dicc.Save("key3", 3)

    sum := 0
    dicc.Iterate(func(key string, value int) bool {
        sum += value
        return true
    })
    require.Equal(t, 6, sum)
}

func TestIterator(t *testing.T) {
    dicc := hash.NewHash[string, int]()
    dicc.Save("key1", 1)
    dicc.Save("key2", 2)
    dicc.Save("key3", 3)

    iter := dicc.Iterator()
    keys := make(map[string]bool)
    for iter.HasNext() {
        key, _ := iter.Current()
        keys[key] = true
        iter.Next()
    }
    require.True(t, keys["key1"])
    require.True(t, keys["key2"])
    require.True(t, keys["key3"])
}
