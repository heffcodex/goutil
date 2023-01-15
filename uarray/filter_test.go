package uarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3, 4, 5, 5}
	res := make([]int, 0, len(arr))

	res = Filter(arr, Unique[int](IDValue[int]))
	require.Equal(t, []int{1, 2, 3, 4, 5}, res)

	res = Filter(arr, Intersection([]int{1, 4}, IDValue[int]))
	require.Equal(t, []int{1, 1, 4}, res)

	res = Filter(arr, Unique[int](IDValue[int]), Intersection([]int{1, 4}, IDValue[int]))
	require.Equal(t, []int{1, 4}, res)

	res = Filter(arr, func(arr []int, i int) bool { return arr[i]%2 == 0 })
	require.Equal(t, []int{2, 4}, res)

	res = Filter(
		arr,
		func(arr []int, i int) bool { return arr[i]%2 == 0 },
		func(arr []int, i int) bool { return arr[i] > 2 },
	)
	require.Equal(t, []int{4}, res)
}

func TestCount(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3, 4, 5, 5}

	require.Equal(t, 5, Count(arr, Unique[int](IDValue[int])))
	require.Equal(t, 3, Count(arr, Intersection([]int{1, 4}, IDValue[int])))
	require.Equal(t, 2, Count(arr, Unique[int](IDValue[int]), Intersection([]int{1, 4}, IDValue[int])))
	require.Equal(t, 2, Count(arr, func(arr []int, i int) bool { return arr[i]%2 == 0 }))
	require.Equal(t, 1, Count(
		arr,
		func(arr []int, i int) bool { return arr[i]%2 == 0 },
		func(arr []int, i int) bool { return arr[i] > 2 },
	))
}
