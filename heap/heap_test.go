package heap_test

import (
	"github.com/FerBuono/go-data-structures/heap"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEmptyHeap(t *testing.T) {
    heap := heap.NewHeap(func(a, b int) int { return a - b })

    require.True(t, heap.IsEmpty())
    require.Panics(t, func() { heap.Peek() })
    require.Panics(t, func() { heap.Dequeue() })
    require.Zero(t, heap.Size())
    require.NotNil(t, heap)
}

func TestHeapFunctionality(t *testing.T) {
    heap := heap.NewHeap(func(a, b int) int { return a - b })

    heap.Enqueue(10)
    heap.Enqueue(5)
    heap.Enqueue(20)

    require.False(t, heap.IsEmpty())
    require.Equal(t, 3, heap.Size())
    require.Equal(t, 20, heap.Peek())
    require.Equal(t, 20, heap.Dequeue())

    require.Equal(t, 10, heap.Dequeue())

    require.Equal(t, 5, heap.Peek())
    require.Equal(t, 1, heap.Size())
    require.Equal(t, 5, heap.Dequeue())
    require.True(t, heap.IsEmpty())
    require.Panics(t, func() { heap.Peek() })
    require.Panics(t, func() { heap.Dequeue() })

    require.Zero(t, heap.Size())
    heap.Enqueue(1)
    require.Equal(t, 1, heap.Size())
    require.Equal(t, 1, heap.Peek())
    require.Equal(t, 1, heap.Dequeue())
}

func TestHeapFromArray(t *testing.T) {
    arr := make([]int, 11)

    for i := 0; i <= 10; i++ {
        arr[i] = i
    }
    heap := heap.NewHeapFromArray(arr, func(a, b int) int { return a - b })

    require.Equal(t, 11, heap.Size())
    for i := 10; i >= 0; i-- {
        require.False(t, heap.IsEmpty())
        require.Equal(t, i, heap.Peek())
        require.Equal(t, i, heap.Dequeue())
    }

    require.True(t, heap.IsEmpty())
    require.Zero(t, heap.Size())
    require.Panics(t, func() { heap.Peek() })
    require.Panics(t, func() { heap.Dequeue() })
}

func TestEmptyHeapFromArray(t *testing.T) {
    arr := []int{}
    heap := heap.NewHeapFromArray(arr, func(a, b int) int { return a - b })
    require.Equal(t, 0, heap.Size())
    heap.Enqueue(1)
    heap.Enqueue(3)
    heap.Enqueue(56)
    heap.Enqueue(2)
    heap.Enqueue(100)
    heap.Enqueue(10)
    require.Equal(t, 100, heap.Dequeue())
    require.Equal(t, 56, heap.Dequeue())
    require.Equal(t, 10, heap.Dequeue())
    require.Equal(t, 3, heap.Dequeue())
    require.Equal(t, 2, heap.Dequeue())
    require.Equal(t, 1, heap.Dequeue())
    require.Equal(t, 0, heap.Size())
    require.True(t, heap.IsEmpty())
}

func TestHeapSort(t *testing.T) {
	sortedMerge := make([]int, 20)
	sortedHeap := make([]int, 20)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
			value := rand.Intn(50)
			sortedHeap[i] = value
			sortedMerge[i] = value
	}

	mergeSort(sortedMerge)
	heap.HeapSort(sortedHeap, func(a, b int) int { return b - a })

	for i := 0; i < 20; i++ {
			require.Equal(t, sortedHeap[i], sortedMerge[i])
	}
}


func TestHeapWithStrings(t *testing.T) {
	heap := heap.NewHeap(func(a, b string) int { return strings.Compare(b, a) }) // max-heap for strings
	elem1 := "Cat"
	elem2 := "Dog"
	elem3 := "Cow"
	elem4 := "Duck"
	elem5 := "Chicken"
	elem6 := "Horse"
	elem7 := ""

	// Correct expected order
	elements := []string{"", "Cat", "Chicken", "Cow", "Dog", "Duck", "Horse"}

	heap.Enqueue(elem1)
	heap.Enqueue(elem2)
	heap.Enqueue(elem3)
	heap.Enqueue(elem4)
	heap.Enqueue(elem5)
	heap.Enqueue(elem6)
	heap.Enqueue(elem7)

	require.Equal(t, 7, heap.Size())
	require.False(t, heap.IsEmpty())
	require.Equal(t, elements[0], heap.Peek())

	for i := 0; i < 7; i++ {
			require.Equal(t, elements[i], heap.Dequeue())
	}

	require.True(t, heap.IsEmpty())
	require.Zero(t, heap.Size())
	require.Panics(t, func() { heap.Peek() })
	require.Panics(t, func() { heap.Dequeue() })
}


func TestHeapVolumeWithMaxValues(t *testing.T) {
    heap := heap.NewHeap(func(a, b int) int { return a - b })

    heap.Enqueue(10000)
    for i := 0; i < 3000; i++ {
        heap.Enqueue(rand.Intn(5000))
    }

    require.Equal(t, 3001, heap.Size())
    require.Equal(t, 10000, heap.Peek())
    require.False(t, heap.IsEmpty())

    previous := heap.Dequeue()
    for i := 3000; i > 0; i-- {
        require.True(t, heap.Peek() <= previous)
        previous = heap.Dequeue()
    }

    require.Zero(t, heap.Size())
    require.True(t, heap.IsEmpty())
    require.Panics(t, func() { heap.Peek() })
    require.Panics(t, func() { heap.Dequeue() })
}

func TestHeapVolumeWithMinValues(t *testing.T) {
    heap := heap.NewHeap(func(a, b int) int { return b - a })

    heap.Enqueue(0)
    for i := 0; i < 3000; i++ {
        heap.Enqueue(rand.Intn(5000))
    }

    require.Equal(t, 3001, heap.Size())
    require.Equal(t, 0, heap.Peek())
    require.False(t, heap.IsEmpty())

    previous := heap.Dequeue()
    for i := 3000; i > 0; i-- {
        require.True(t, heap.Peek() >= previous)
        previous = heap.Dequeue()
    }

    require.Zero(t, heap.Size())
    require.True(t, heap.IsEmpty())
    require.Panics(t, func() { heap.Peek() })
    require.Panics(t, func() { heap.Dequeue() })
}

func TestHeapFromArrayVolume(t *testing.T) {
    arr := make([]int, 3001)

    for i := 0; i <= 3000; i++ {
        arr[i] = rand.Intn(5000)
    }
    arr[2000] = 10000

    heap := heap.NewHeapFromArray(arr, func(a, b int) int { return a - b })

    require.Equal(t, 3001, heap.Size())
    require.Equal(t, 10000, heap.Peek())

    previous := heap.Dequeue()
    for i := 0; i < 3000; i++ {
        require.False(t, heap.IsEmpty())
        require.True(t, previous >= heap.Peek())
        previous = heap.Dequeue()
    }

    require.True(t, heap.IsEmpty())
    require.Zero(t, heap.Size())
    require.Panics(t, func() { heap.Peek() })
    require.Panics(t, func() { heap.Dequeue() })
}

func TestHeapSortVolume(t *testing.T) {
	sortedMerge := make([]int, 3001)
	sortedHeap := make([]int, 3001)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3000; i++ {
			value := rand.Intn(50)
			sortedHeap[i] = value
			sortedMerge[i] = value
	}

	mergeSort(sortedMerge)
	heap.HeapSort(sortedHeap, func(a, b int) int { return b - a })

	for i := 0; i < 3000; i++ {
			require.Equal(t, sortedHeap[i], sortedMerge[i])
	}
}


// Auxiliary function

func mergeSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    mid := len(arr) / 2
    return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
    length := len(left) + len(right)
    arr := make([]int, length)
    i, j := 0, 0
    for k := 0; k < length; k++ {
        if i > len(left)-1 && j <= len(right)-1 {
            arr[k] = right[j]
            j++
        } else if j > len(right)-1 && i <= len(left)-1 {
            arr[k] = left[i]
            i++
        } else if left[i] < right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
    }
    return arr
}
