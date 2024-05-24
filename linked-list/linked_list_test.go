package linked_list_test

import (
	"github.com/FerBuono/go-data-structures/linked-list"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyList(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()

	require.PanicsWithValue(t, "The list is empty", func() { list.DeleteFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeLast() })
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Length())
}

func TestInsertElementsAtStart(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	list.InsertFirst(1)
	require.Equal(t, 1, list.SeeFirst())
	require.Equal(t, 1, list.SeeLast())
	list.InsertFirst(10)
	require.Equal(t, 10, list.SeeFirst())
	require.NotEqual(t, 10, list.SeeLast())
	list.InsertFirst(20)
	list.InsertFirst(30)
	require.Equal(t, 30, list.SeeFirst())
	require.Equal(t, 1, list.SeeLast())
}

func TestInsertElementsAtEnd(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	list.InsertLast(1)
	require.Equal(t, 1, list.SeeFirst())
	require.Equal(t, 1, list.SeeLast())
	list.InsertLast(10)
	require.Equal(t, 10, list.SeeLast())
	require.NotEqual(t, 10, list.SeeFirst())
	list.InsertLast(20)
	list.InsertLast(30)
	require.Equal(t, 30, list.SeeLast())
	require.Equal(t, 1, list.SeeFirst())
}

func TestInsertDeleteElements(t *testing.T) {
	list := linked_list.CreateLinkedList[string]()

	list.InsertFirst("first")
	require.Equal(t, "first", list.SeeFirst())
	require.Equal(t, "first", list.SeeLast())
	require.Equal(t, 1, list.Length())
	require.Equal(t, "first", list.DeleteFirst())
	require.Equal(t, 0, list.Length())
	require.True(t, list.IsEmpty())
	require.PanicsWithValue(t, "The list is empty", func() { list.DeleteFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeLast() })

	list.InsertLast("last")
	require.Equal(t, "last", list.SeeFirst())
	require.Equal(t, "last", list.SeeLast())
	require.Equal(t, 1, list.Length())
	require.Equal(t, "last", list.DeleteFirst())
	require.Equal(t, 0, list.Length())
	require.True(t, list.IsEmpty())
	require.PanicsWithValue(t, "The list is empty", func() { list.DeleteFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeLast() })

	list.InsertLast("first")
	list.InsertLast("last")
	require.Equal(t, "first", list.SeeFirst())
	require.Equal(t, "last", list.SeeLast())
	require.Equal(t, 2, list.Length())
	require.False(t, list.IsEmpty())
	require.Equal(t, "first", list.DeleteFirst())
	require.Equal(t, "last", list.DeleteFirst())
	require.Equal(t, 0, list.Length())
	require.True(t, list.IsEmpty())
	require.PanicsWithValue(t, "The list is empty", func() { list.DeleteFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeLast() })

	list.InsertFirst("four")
	list.InsertFirst("three")
	list.InsertFirst("two")
	list.InsertFirst("one")
	require.Equal(t, "one", list.SeeFirst())
	require.Equal(t, "four", list.SeeLast())
	require.Equal(t, 4, list.Length())
	require.False(t, list.IsEmpty())
	require.Equal(t, "one", list.DeleteFirst())
	require.Equal(t, "two", list.DeleteFirst())
	require.Equal(t, "three", list.DeleteFirst())
	require.Equal(t, "four", list.DeleteFirst())
	require.Equal(t, 0, list.Length())
	require.True(t, list.IsEmpty())
	require.PanicsWithValue(t, "The list is empty", func() { list.DeleteFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeLast() })
}

func TestVolume(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()

	for i := 0; i <= 1000; i++ {
		list.InsertLast(i)
		require.Equal(t, 0, list.SeeFirst())
		require.Equal(t, i, list.SeeLast())
	}
	require.Greater(t, list.Length(), 1000)
	require.False(t, list.IsEmpty())

	for j := 0; j <= 1000; j++ {
		require.Equal(t, j, list.SeeFirst())
		require.Equal(t, 1000, list.SeeLast())
		require.Equal(t, j, list.DeleteFirst())
	}
	require.Equal(t, 0, list.Length())
	require.True(t, list.IsEmpty())
	require.PanicsWithValue(t, "The list is empty", func() { list.DeleteFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeFirst() })
	require.PanicsWithValue(t, "The list is empty", func() { list.SeeLast() })
}

func TestInsertElementsWithIterator(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()

	for i := 1; i < 10; i++ {
		list.InsertLast(i)
	}

	iter := list.Iterator()

	iter.Insert(0)
	require.Equal(t, 0, list.SeeFirst())
	require.Equal(t, iter.SeeCurrent(), list.SeeFirst())

	for i := 0; i < list.Length()/2; i++ {
		iter.Next()
	}

	current := iter.SeeCurrent()
	iter.Insert(26)
	require.Equal(t, 26, iter.Next())
	require.Equal(t, current, iter.SeeCurrent())

	for iter.HasNext() {
		iter.Next()
	}

	iter.Insert(10)
	require.Equal(t, 10, list.SeeLast())
	require.Equal(t, iter.SeeCurrent(), list.SeeLast())
}

func TestDeleteElementsWithIterator(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()

	iter := list.Iterator()
	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.SeeCurrent() })
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Next() })

	for i := 1; i < 10; i++ {
		list.InsertLast(i)
	}

	iter = list.Iterator()

	first := list.SeeFirst()
	require.Equal(t, first, iter.Delete())
	require.NotEqual(t, first, list.SeeFirst())

	for i := 0; i < list.Length()/2; i++ {
		iter.Next()
	}

	current := iter.SeeCurrent()
	require.Equal(t, current, iter.Delete())
	require.NotEqual(t, current, iter.SeeCurrent())

	iter = list.Iterator()

	for i := 0; i < list.Length()-1; i++ {
		iter.Next()
	}

	last := list.SeeLast()
	require.Equal(t, last, iter.Delete())
	require.NotEqual(t, last, list.SeeLast())
}

func TestIteratorFunctionality(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	iter := list.Iterator()
	for i := 0; i <= 10; i++ {
		iter.Insert(i)
	}
	for i := 10; i >= 0; i-- {
		require.True(t, iter.HasNext())
		require.Equal(t, i, iter.Next())
	}
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.SeeCurrent() })
	require.PanicsWithValue(t, "The iterator has finished iterating", func() { iter.Next() })
	iter2 := list.Iterator()
	for i := 10; i >= 0; i-- {
		require.True(t, iter2.HasNext())
		require.Equal(t, i, iter2.Delete())
	}
	require.Equal(t, 0, list.Length())
	for i := 0; i <= 10; i++ {
		iter2.Insert(i)
		if i == 5 {
			iter2.Insert(30)
		}
	}
	iter3 := list.Iterator()
	for i := 10; i >= 0; i-- {
		if i == 5 {
			require.True(t, iter3.HasNext())
			require.Equal(t, 30, iter3.Delete())
		}
		require.True(t, iter3.HasNext())
		require.Equal(t, i, iter3.Delete())
	}
}

func TestIteratorVolume(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	iter := list.Iterator()
	for i := 0; i <= 2000; i++ {
		iter.Insert(i)
	}
	for i := 2000; i >= 0; i-- {
		require.True(t, iter.HasNext())
		require.Equal(t, i, iter.Next())
	}
	iter2 := list.Iterator()
	for i := 2000; i >= 0; i-- {
		require.True(t, iter2.HasNext())
		require.Equal(t, i, iter2.Delete())
	}
	require.Equal(t, 0, list.Length())
}

func TestInternalIteratorFunctionality(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	for i := 0; i < 10; i++ {
		list.InsertFirst(i)
	}
	iteratedData := 0
	list.Iterate(func(data int) bool {
		iteratedData++
		return true
	})
	require.Equal(t, 10, iteratedData)
}

func TestIteratorEndsAtEndOfList(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	for i := 0; i <= 10; i++ {
		list.InsertLast(i)
	}
	iteratedSum := 0
	list.Iterate(func(data int) bool {
		require.Equal(t, data, iteratedSum)
		iteratedSum++
		return true
	})
}

func TestIteratorEndsWhenUserIndicates(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	sum := 0
	for i := 0; i < 20; i++ {
		list.InsertLast(i)
		sum += i
	}
	iteratedSum := 0
	list.Iterate(func(data int) bool {
		iteratedSum += data
		return iteratedSum < 20
	})
	require.Less(t, iteratedSum, sum)
	iteratedSum = sum
	list.Iterate(func(data int) bool {
		iteratedSum = iteratedSum - (data * 2)
		return iteratedSum > 0
	})
	require.Less(t, iteratedSum, sum)
}

func TestInternalIteratorEmptyList(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	number := 0
	list.Iterate(func(data int) bool {
		number += data
		return true
	})
	require.Zero(t, number)
}

func TestVolumeInternalIterator(t *testing.T) {
	list := linked_list.CreateLinkedList[int]()
	sum := 0
	for i := 0; i < 2000; i++ {
		sum += i
		list.InsertLast(i)
	}
	iteratedSum := 0
	list.Iterate(func(data int) bool {
		iteratedSum += data
		return true
	})
}