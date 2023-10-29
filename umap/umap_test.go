package umap

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestFromArray(t *testing.T) {
	t.Parallel()

	m := FromArray([]int{1, 2, 3, 4, 5}, func(item int) int {
		return item
	})

	assert.Len(t, m, 5)
	assert.Equal(t, 1, m[1])
	assert.Equal(t, 2, m[2])
	assert.Equal(t, 3, m[3])
	assert.Equal(t, 4, m[4])
	assert.Equal(t, 5, m[5])
}

func TestMap(t *testing.T) {
	t.Parallel()

	type (
		mIn  map[int]bool
		mOut map[string]int
	)

	in := mIn{1: true, 2: false, 3: true, 4: false}

	var out mOut = Map(in, func(k int, v bool) (string, int) {
		out := 0
		if v {
			out = 1
		}

		return strconv.Itoa(k), out
	})

	assert.ElementsMatch(t, maps.Keys(out), []string{"1", "2", "3", "4"})
	assert.Equal(t, 1, out["1"])
	assert.Equal(t, 0, out["2"])
	assert.Equal(t, 1, out["3"])
	assert.Equal(t, 0, out["4"])
}

func TestFindAll(t *testing.T) {
	t.Parallel()

	m := map[uint]int{1: -1, 2: -2, 3: -3}

	assert.ElementsMatch(t, FindAll(m, 1, 2, 5), []int{-1, -2})
}
