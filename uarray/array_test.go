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
