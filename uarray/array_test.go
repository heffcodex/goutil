package uarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	res := Map(arr, func(item int) int { return item * 2 })
	require.Equal(t, []int{2, 4, 6, 8, 10}, res)
}

func TestMerge(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}

	res := Merge(arr1, arr2, arr3)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
}
