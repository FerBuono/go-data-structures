package linked_queue_test

import (
    "testing"

    "github.com/stretchr/testify/require"
    "github.com/FerBuono/go-data-structures/linked-queue"
)

func TestQueueIsEmpty(t *testing.T) {
    q := linked_queue.NewLinkedQueue[int]()

    require.True(t, q.IsEmpty())
    require.Panics(t, func() { q.Peek() })
    require.Panics(t, func() { q.Dequeue() })
}

func TestEnqueueAndPeek(t *testing.T) {
    q := linked_queue.NewLinkedQueue[int]()
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    require.False(t, q.IsEmpty())
    require.Equal(t, 1, q.Peek())
}

func TestEnqueueAndDequeue(t *testing.T) {
    q := linked_queue.NewLinkedQueue[int]()
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    require.Equal(t, 1, q.Dequeue())
    require.Equal(t, 2, q.Dequeue())
    require.Equal(t, 3, q.Dequeue())

    require.True(t, q.IsEmpty())
    require.Panics(t, func() { q.Dequeue() })
}

func TestMixedOperations(t *testing.T) {
    q := linked_queue.NewLinkedQueue[int]()
    q.Enqueue(1)
    require.Equal(t, 1, q.Dequeue())
    q.Enqueue(2)
    q.Enqueue(3)
    require.Equal(t, 2, q.Dequeue())
    q.Enqueue(4)
    require.Equal(t, 3, q.Dequeue())
    require.Equal(t, 4, q.Dequeue())

    require.True(t, q.IsEmpty())
    require.Panics(t, func() { q.Dequeue() })
}
