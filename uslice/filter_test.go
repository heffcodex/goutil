package uslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	var (
		arr = []int{1, 1, 2, 3, 3, 4, 5, 5}
		res []int
	)

	res = Filter(arr, Unique(KeyValue[int]))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)

	res = Filter(arr, Intersection([]int{1, 4}, KeyValue[int]))
	assert.Equal(t, []int{1, 1, 4}, res)

	res = Filter(arr, Unique(KeyValue[int]), Intersection([]int{1, 4}, KeyValue[int]))
	assert.Equal(t, []int{1, 4}, res)

	res = Filter(arr, func(item int) bool { return item%2 == 0 })
	assert.Equal(t, []int{2, 4}, res)

	res = Filter(
		arr,
		func(item int) bool { return item%2 == 0 },
		func(item int) bool { return item > 2 },
	)
	assert.Equal(t, []int{4}, res)
}

func TestCount(t *testing.T) {
	t.Parallel()

	arr := []int{1, 1, 2, 3, 3, 4, 5, 5}

	assert.Equal(t, 5, Count(arr, Unique(KeyValue[int])))
	assert.Equal(t, 3, Count(arr, Intersection([]int{1, 4}, KeyValue[int])))
	assert.Equal(t, 2, Count(arr, Unique(KeyValue[int]), Intersection([]int{1, 4}, KeyValue[int])))
	assert.Equal(t, 2, Count(arr, func(item int) bool { return item%2 == 0 }))
	assert.Equal(t, 1, Count(
		arr,
		func(item int) bool { return item%2 == 0 },
		func(item int) bool { return item > 2 },
	))
}
