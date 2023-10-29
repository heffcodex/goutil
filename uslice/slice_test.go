package uslice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testSlice []int

func testSlices() []testSlice {
	return []testSlice{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	res := Map(func(item int) int { return item * 2 }, testSlices()...)
	require.Equal(t, []int{2, 4, 6, 8, 10, 12, 14, 16, 18}, res)
}

func TestMerge(t *testing.T) {
	t.Parallel()

	res := Merge(testSlices()...)
	require.Equal(t, testSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}, res)
}

func TestLen(t *testing.T) {
	t.Parallel()

	res := Len(testSlices()...)
	require.Equal(t, 9, res)
}
