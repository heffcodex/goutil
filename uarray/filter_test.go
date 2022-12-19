package uarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3, 4, 5, 5}
	res := make([]int, 0, len(arr))

	res = Filter(arr, All[int]())
	require.Equal(t, arr, res)

	res = Filter(arr, None[int]())
	require.Empty(t, res)

	res = Filter(arr, Unique[int]())
	require.Equal(t, []int{1, 2, 3, 4, 5}, res)

	res = Filter(arr, Intersection([]int{1, 4}))
	require.Equal(t, []int{1, 1, 4}, res)

	res = Filter(arr, Unique[int](), Intersection([]int{1, 4}))
	require.Equal(t, []int{1, 4}, res)

	res = Filter(arr, func(item int) bool { return item%2 == 0 })
	require.Equal(t, []int{2, 4}, res)

	res = Filter(
		arr,
		func(item int) bool { return item%2 == 0 },
		func(item int) bool { return item > 2 },
	)
	require.Equal(t, []int{4}, res)
}

func TestLen(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3, 4, 5, 5}

	require.Equal(t, len(arr), Len(arr, All[int]()))
	require.Zero(t, Len(arr, None[int]()))
	require.Equal(t, 5, Len(arr, Unique[int]()))
	require.Equal(t, 3, Len(arr, Intersection([]int{1, 4})))
	require.Equal(t, 2, Len(arr, Unique[int](), Intersection([]int{1, 4})))
	require.Equal(t, 2, Len(arr, func(item int) bool { return item%2 == 0 }))
	require.Equal(t, 1, Len(
		arr,
		func(item int) bool { return item%2 == 0 },
		func(item int) bool { return item > 2 },
	))
}
