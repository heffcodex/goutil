package uslice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromMap(t *testing.T) {
	t.Parallel()

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	keys := FromMap(m, UseKeys[string, int])
	assert.ElementsMatch(t, []string{"one", "two", "three"}, keys)

	values := FromMap(m, UseValues[string, int])
	assert.ElementsMatch(t, []int{1, 2, 3}, values)

	combined := FromMap(m, func(k string, v int) string { return k + ":" + strconv.Itoa(v) })
	assert.ElementsMatch(t, []string{"one:1", "two:2", "three:3"}, combined)
}
