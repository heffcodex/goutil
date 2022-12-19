package uarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	t.Run("value", func(t *testing.T) {
		idx := FindIndex(arr, Value(3))
		require.Equal(t, 2, idx)

		idx = FindIndex(arr, Value(6))
		require.Equal(t, IdxNotFound, idx)

		t.Run("not", func(t *testing.T) {
			idx = FindIndex(arr, Not(Value(3)))
			require.Equal(t, 0, idx)

			idx = FindIndex(arr, Not(Value(1)))
			require.Equal(t, 1, idx)
		})
	})

	t.Run("any", func(t *testing.T) {
		idx := FindIndex(arr, Any(
			func(arr []int, i int) bool { return arr[i] == 0 },
			func(arr []int, i int) bool { return arr[i] == 3 },
		))
		require.Equal(t, 2, idx)

		idx = FindIndex(arr, Any(
			func(arr []int, i int) bool { return arr[i] == 0 },
			func(arr []int, i int) bool { return arr[i] == 6 },
		))
		require.Equal(t, IdxNotFound, idx)

		t.Run("not", func(t *testing.T) {
			idx = FindIndex(arr, Any(
				func(arr []int, i int) bool { return arr[i] == 0 },
				Not(func(arr []int, i int) bool { return arr[i] == 3 }),
			))
			require.Equal(t, 0, idx)

			idx = FindIndex(arr, Any(
				func(arr []int, i int) bool { return arr[i] == 0 },
				Not(func(arr []int, i int) bool { return arr[i] == 1 }),
			))
			require.Equal(t, 1, idx)
		})
	})

	t.Run("any value", func(t *testing.T) {
		idx := FindIndex(arr, AnyValue(0, 3))
		require.Equal(t, 2, idx)

		idx = FindIndex(arr, AnyValue(0, 6))
		require.Equal(t, IdxNotFound, idx)
	})
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
