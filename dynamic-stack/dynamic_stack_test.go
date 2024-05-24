package dynamic_stack_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/FerBuono/go-data-structures/dynamic-stack"
)

func TestDynamicStack(t *testing.T) {
	s := dynamic_stack.NewDynamicStack[int]()

	// Test initial state
	require.True(t, s.IsEmpty())

	// Test pushing elements
	s.Push(1)
	require.False(t, s.IsEmpty())
	require.Equal(t, 1, s.Top())

	s.Push(2)
	require.Equal(t, 2, s.Top())

	s.Push(3)
	require.Equal(t, 3, s.Top())

	// Test popping elements
	require.Equal(t, 3, s.Pop())
	require.Equal(t, 2, s.Top())

	require.Equal(t, 2, s.Pop())
	require.Equal(t, 1, s.Top())

	require.Equal(t, 1, s.Pop())
	require.True(t, s.IsEmpty())

	// Test popping from an empty stack (should panic)
	require.PanicsWithValue(t, "The stack is empty", func() { s.Pop() })
}

func TestDynamicStackResize(t *testing.T) {
	s := dynamic_stack.NewDynamicStack[int]()

	// Push elements to trigger resize
	for i := 0; i < 20; i++ {
		s.Push(i)
	}

	require.Equal(t, 19, s.Top())

	// Pop elements to trigger shrink
	for i := 19; i >= 0; i-- {
		require.Equal(t, i, s.Pop())
	}

	require.True(t, s.IsEmpty())
}
