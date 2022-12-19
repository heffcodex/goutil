package uarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	idx := FindIndex(arr, Value(3))
	require.Equal(t, 2, idx)

	idx = FindIndex(arr, Value(6))
	require.Equal(t, IdxNotFound, idx)

	idx = FindIndex(arr, AnyValue(0, 3))
	require.Equal(t, 2, idx)

	idx = FindIndex(arr, AnyValue(0, 6))
	require.Equal(t, IdxNotFound, idx)

	idx = FindIndex(arr, Any(
		func(item int) bool { return item == 0 },
		func(item int) bool { return item == 3 },
	))
	require.Equal(t, 2, idx)

	idx = FindIndex(arr, Any(
		func(item int) bool { return item == 0 },
		func(item int) bool { return item == 6 },
	))
	require.Equal(t, IdxNotFound, idx)
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	item, ok := Find(arr, Value(3))
	require.True(t, ok)
	require.Equal(t, 3, item)

	item, ok = Find(arr, Value(6))
	require.False(t, ok)
	require.Equal(t, 0, item)
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	require.True(t, Contains(arr, Value(3)))
	require.False(t, Contains(arr, Value(6)))
}
