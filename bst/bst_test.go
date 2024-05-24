package bst_test

import (
		"github.com/FerBuono/go-data-structures/bst"
    "fmt"
    "math/rand"
    "strings"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

var VOLUME_SIZES = []int{1000, 2000, 4000}

func TestEmptyBST(t *testing.T) {
    t.Log("Checks that an empty BST has no keys")
    tree := bst.NewBST[int, int](func(a, b int) int { return a - b })
    require.NotNil(t, tree)
    require.Equal(t, 0, tree.Size())
    require.False(t, tree.Contains(1))
    require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { tree.Contains(1) })
    require.PanicsWithValue(t, "The key does not belong to the dictionary", func() { tree.Delete(1) })
}

func TestOneElement(t *testing.T) {
    t.Log("Checks that a BST with one element has that key only")
    tree := bst.NewBST[string, int](func(a, b string) int { return strings.Compare(a, b) })
    tree.Save("A", 10)
    require.EqualValues(t, 1, tree.Size())
    require.True(t, tree.Contains("A"))
    require.False(t, tree.Contains("B"))
    require.EqualValues(t, 10, tree.Contains("A"))
    require.PanicsWithValue(t, "Key does not exist in the dictionary", func() { tree.Contains("B") })
}

func TestBSTInsert(t *testing.T) {
    t.Log("Inserts a few elements into the BST and checks that it behaves as expected")
    key1 := "Cat"
    key2 := "Dog"
    key3 := "Cow"
    value1 := "meow"
    value2 := "woof"
    value3 := "moo"
    keys := []string{key1, key2, key3}
    values := []string{value1, value2, value3}

    tree := bst.NewBST[string, string](func(a, b string) int { return strings.Compare(a, b) })
    require.False(t, tree.Contains(keys[0]))
    require.False(t, tree.Contains(keys[0]))
    tree.Save(keys[0], values[0])
    require.EqualValues(t, 1, tree.Size())
    require.True(t, tree.Contains(keys[0]))
    require.True(t, tree.Contains(keys[0]))
    require.EqualValues(t, values[0], tree.Contains(keys[0]))
    require.EqualValues(t, values[0], tree.Contains(keys[0]))

    require.False(t, tree.Contains(keys[1]))
    require.False(t, tree.Contains(keys[2]))
    tree.Save(keys[1], values[1])
    require.True(t, tree.Contains(keys[0]))
    require.True(t, tree.Contains(keys[1]))
    require.EqualValues(t, 2, tree.Size())
    require.EqualValues(t, values[0], tree.Contains(keys[0]))
    require.EqualValues(t, values[1], tree.Contains(keys[1]))

    require.False(t, tree.Contains(keys[2]))
    tree.Save(keys[2], values[2])
    require.True(t, tree.Contains(keys[0]))
    require.True(t, tree.Contains(keys[1]))
    require.True(t, tree.Contains(keys[2]))
    require.EqualValues(t, 3, tree.Size())
    require.EqualValues(t, values[0], tree.Contains(keys[0]))
    require.EqualValues(t, values[1], tree.Contains(keys[1]))
    require.EqualValues(t, values[2], tree.Contains(keys[2]))
}

func TestReplaceValue(t *testing.T) {
    t.Log("Inserts a few keys, then replaces their values")
    key := "Cat"
    key2 := "Dog"
    tree := bst.NewBST[string, string](func(a, b string) int { return strings.Compare(a, b) })
    tree.Save(key, "meow")
    tree.Save(key2, "woof")
    require.True(t, tree.Contains(key))
    require.True(t, tree.Contains(key2))
    require.EqualValues(t, "meow", tree.Contains(key))
    require.EqualValues(t, "woof", tree.Contains(key2))
    require.EqualValues(t, 2, tree.Size())

    tree.Save(key, "mew")
    tree.Save(key2, "bark")
    require.True(t, tree.Contains(key))
    require.True(t, tree.Contains(key2))
    require.EqualValues(t, 2, tree.Size())
    require.EqualValues(t, "mew", tree.Contains(key))
    require.EqualValues(t, "bark", tree.Contains(key2))
}

func TestBSTDelete(t *testing.T) {
    t.Log("Inserts a few elements into the BST, deletes them, and checks that it behaves as expected")
    key1 := "Cat"
    key2 := "Dog"
    key3 := "Cow"
    value1 := "meow"
    value2 := "woof"
    value3 := "moo"
    keys := []string{key1, key2, key3}
    values := []string{value1, value2, value3}
    tree := bst.NewBST[string, string](func(a, b string) int { return strings.Compare(a, b) })

    require.False(t, tree.Contains(keys[0]))
    require.False(t, tree.Contains(keys[0]))
    tree.Save(keys[0], values[0])
    tree.Save(keys[1], values[1])
    tree.Save(keys[2], values[2])

    require.True(t, tree.Contains(keys[2]))
    require.EqualValues(t, values[2], tree.Delete(keys[2]))
    require.PanicsWithValue(t, "Key does not exist in the dictionary", func() { tree.Delete(keys[2]) })
    require.EqualValues(t, 2, tree.Size())
    require.False(t, tree.Contains(keys[2]))

    require.True(t, tree.Contains(keys[0]))
    require.EqualValues(t, values[0], tree.Delete(keys[0]))
    require.PanicsWithValue(t, "Key does not exist in the dictionary", func() { tree.Delete(keys[0]) })
    require.EqualValues(t, 1, tree.Size())
    require.False(t, tree.Contains(keys[0]))
    require.PanicsWithValue(t, "Key does not exist in the dictionary", func() { tree.Contains(keys[0]) })

    require.True(t, tree.Contains(keys[1]))
    require.EqualValues(t, values[1], tree.Delete(keys[1]))
    require.PanicsWithValue(t, "Key does not exist in the dictionary", func() { tree.Delete(keys[1]) })
    require.EqualValues(t, 0, tree.Size())
    require.False(t, tree.Contains(keys[1]))
    require.PanicsWithValue(t, "Key does not exist in the dictionary", func() { tree.Contains(keys[1]) })
}

func TestReuseAfterDeletion(t *testing.T) {
    t.Log("White-box test: checks that an element can be reinserted after deletion")
    tree := bst.NewBST[string, string](func(a, b string) int { return strings.Compare(a, b) })
    key := "hello"
    tree.Save(key, "world!")
    tree.Delete(key)
    require.EqualValues(t, 0, tree.Size())
    require.False(t, tree.Contains(key))
    tree.Save(key, "world again!")
    require.True(t, tree.Contains(key))
    require.EqualValues(t, 1, tree.Size())
    require.EqualValues(t, "world again!", tree.Contains(key))
}

func TestNumericalKeys(t *testing.T) {
    t.Log("Checks that the BST works with numeric keys")
    tree := bst.NewBST[int, string](func(a, b int) int { return a - b })
    key := 10
    value := "Kitty"

    tree.Save(key, value)
    require.EqualValues(t, 1, tree.Size())
    require.True(t, tree.Contains(key))
    require.EqualValues(t, value, tree.Contains(key))
    require.EqualValues(t, value, tree.Delete(key))
    require.False(t, tree.Contains(key))
}

func TestStructKeys(t *testing.T) {
    t.Log("Checks that the BST works with more complex structures as keys")
    type basic struct {
        str  string
        num  int
    }

    tree := bst.NewBST[basic, int](func(a, b basic) int { return strings.Compare(a.str, b.str) })

    b2 := basic{str: "world", num: 14}
    b1 := basic{str: "earth", num: 8}
    b3 := basic{str: "globe", num: 8}

    tree.Save(b1, 0)
    tree.Save(b2, 1)
    tree.Save(b3, 2)

    require.True(t, tree.Contains(b1))
    require.True(t, tree.Contains(b2))
    require.True(t, tree.Contains(b3))
    require.EqualValues(t, 0, tree.Contains(b1))
    require.EqualValues(t, 1, tree.Contains(b2))
    require.EqualValues(t, 2, tree.Contains(b3))
    tree.Save(b1, 5)
    require.EqualValues(t, 5, tree.Contains(b1))
    require.EqualValues(t, 2, tree.Contains(b3))
    require.EqualValues(t, 5, tree.Delete(b1))
    require.False(t, tree.Contains(b1))
    require.EqualValues(t, 2, tree.Contains(b3))
}

func TestEmptyKey(t *testing.T) {
    t.Log("Inserts an empty key (i.e., \"\") and checks that it works fine")
    tree := bst.NewBST[string, string](func(a, b string) int { return strings.Compare(a, b) })
    key := ""
    tree.Save(key, key)
    require.True(t, tree.Contains(key))
    require.EqualValues(t, 1, tree.Size())
    require.EqualValues(t, key, tree.Contains(key))
}

func TestNilValue(t *testing.T) {
    t.Log("Checks that the value can be nil")
    tree := bst.NewBST[string, *int](func(a, b string) int { return strings.Compare(a, b) })
    key := "Fish"
    tree.Save(key, nil)
    require.True(t, tree.Contains(key))
    require.EqualValues(t, 1, tree.Size())
    require.EqualValues(t, (*int)(nil), tree.Contains(key))
    require.EqualValues(t, (*int)(nil), tree.Delete(key))
    require.False(t, tree.Contains(key))
}

func TestInternalIteratorKeys(t *testing.T) {
    t.Log("Checks that all keys are visited (and only once) with the internal iterator")
    key1 := 1
    key2 := 15
    key3 := 10
    keys := []int{key1, key2, key3}
    tree := bst.NewBST[int, *int](func(a, b int) int { return a - b })
    tree.Save(keys[0], nil)
    tree.Save(keys[1], nil)
    tree.Save(keys[2], nil)

    visitedKeys := make([]int, 3)
    count := 0
    countPtr := &count

    tree.Iterate(func(key int, _ *int) bool {
        visitedKeys[count] = key
        *countPtr += 1
        return true
    })

    require.EqualValues(t, 3, count)
    require.EqualValues(t, keys[0], visitedKeys[0])
    require.EqualValues(t, keys[1], visitedKeys[2])
    require.EqualValues(t, keys[2], visitedKeys[1])
    require.NotEqualValues(t, visitedKeys[0], visitedKeys[1])
    require.NotEqualValues(t, visitedKeys[0], visitedKeys[2])
    require.NotEqualValues(t, visitedKeys[2], visitedKeys[1])
}

func TestInternalIteratorValues(t *testing.T) {
    t.Log("Checks that all values are visited correctly (and only once) with the internal iterator")
    key1 := "Hamster"
    key2 := "Cow"
    key3 := "Dog"
    key4 := "Donkey"
    key5 := "Cat"

    tree := bst.NewBST[string, int](func(a, b string) int { return strings.Compare(a, b) })
    tree.Save(key1, 6)
    tree.Save(key2, 2)
    tree.Save(key3, 3)
    tree.Save(key4, 4)
    tree.Save(key5, 5)

    factorial := 1
    ptrFactorial := &factorial
    tree.Iterate(func(_ string, value int) bool {
        *ptrFactorial *= value
        return true
    })

    require.EqualValues(t, 720, factorial)
}

func executeVolumeTest(b *testing.B, n int) {
    tree := bst.NewBST[string, int](func(a, b string) int { return strings.Compare(a, b) })

    keys := make([]string, n)
    values := make([]int, n)

    for i := 0; i < n; i++ {
        keys[i] = fmt.Sprintf("%08d", i)
    }

    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

    for i := 0; i < n; i++ {
        values[i] = i
        tree.Save(keys[i], values[i])
    }
    require.EqualValues(b, n, tree.Size(), "The number of elements is incorrect")

    // Checks that the correct values are returned
    ok := true
    for i := 0; i < n; i++ {
        ok = tree.Contains(keys[i])
        if !ok {
            break
        }
        ok = tree.Contains(keys[i])
        if !ok {
            break
        }
    }

    require.True(b, ok, "Contains and Search with many elements do not work correctly")
    require.EqualValues(b, n, tree.Size(), "The number of elements is incorrect")

    // Checks that elements can be deleted and the correct values are returned
    for i := 0; i < n; i++ {
        ok = tree.Delete(keys[i]) == values[i]
        if !ok {
            break
        }
    }

    require.True(b, ok, "Deleting many elements does not work correctly")
    require.EqualValues(b, 0, tree.Size())
}

func BenchmarkBST(b *testing.B) {
    b.Log("Stress test for the BST. Tests inserting a large number of elements, then retrieving, checking membership, and deleting them.")
    for _, n := range VOLUME_SIZES {
        b.Run(fmt.Sprintf("Test %d elements", n), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                executeVolumeTest(b, n)
            }
        })
    }
}

func TestOutOfRangeElements(t *testing.T) {
    cmp := func(a, b int) int {
        return a - b
    }
    tree := bst.NewBST[int, *int](cmp)
    root := 10
    tree.Save(root, &root)
    for i := 0; i > 20; i++ {
        tree.Save(i, &i)
    }
    sum := 0
    visit := func(_ int, value *int) bool {
        sum += *value
        return true
    }
    from := 30
    to := 50
    tree.IterateRange(&from, &to, visit)
    require.Zero(t, sum)
}

func TestIteratorCutoffFunction(t *testing.T) {
    cmp := func(a, b int) int {
        return a - b
    }
    arr := []int{8, 4, 12, 2, 5, 6, 7, 9, 1, 16, 13, 3, 10, 15}
    tree := bst.NewBST[int, *int](cmp)

    for i := 0; i < len(arr); i++ {
        tree.Save(arr[i], &arr[i])
    }

    from := 1
    to := 15
    sum := 0

    // With this array and conditions, the sum will be 45
    visit := func(_ int, value *int) bool {
        sum += *value
        return sum < 45
    }
    tree.IterateRange(&from, &to, visit)
    require.Equal(t, 45, sum)
}

func TestEmptyRangeIterator(t *testing.T) {
    tree := bst.NewBST[string, *int](func(a, b string) int { return strings.Compare(a, b) })
    from := "Hello"
    to := "Goodbye"
    iter := tree.RangeIterator(&from, &to)
    require.PanicsWithValue(t, "Iterator has finished iterating", func() { iter.Next() })
    require.PanicsWithValue(t, "Iterator has finished iterating", func() { iter.Current() })
    require.False(t, iter.HasNext())
}

func TestOutOfRangeIterator(t *testing.T) {
    cmp := func(a, b int) int {
        return a - b
    }
    tree := bst.NewBST[int, *int](cmp)
    root := 10
    tree.Save(root, &root)
    for i := 0; i > 20; i++ {
        tree.Save(i, &i)
    }
    from := 30
    to := 50
    iter := tree.RangeIterator(&from, &to)
    require.PanicsWithValue(t, "Iterator has finished iterating", func() { iter.Next() })
    require.PanicsWithValue(t, "Iterator has finished iterating", func() { iter.Current() })
    require.False(t, iter.HasNext())
}

func TestRangeIteratorVolume(t *testing.T) {
    cmp := func(a, b int) int {
        return a - b
    }
    tree := bst.NewBST[int, *int](cmp)
    root := 500
    tree.Save(root, &root)

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 2500; i++ {
        random := rand.Int()
        tree.Save(rand.Intn(1000), &random)
    }

    from := 500
    to := 750

    for iter := tree.RangeIterator(&from, &to); iter.HasNext(); {
        key, _ := iter.Current()
        require.True(t, key >= from && key <= to)
        iter.Next()
    }
}
