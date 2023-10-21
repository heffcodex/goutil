package umap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromArray(t *testing.T) {
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
