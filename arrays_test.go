package goutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrArrayIntersect(t *testing.T) {
	a := []string{"A", "a", "b"}
	b := []string{"a"}
	c := []string{"c"}

	require.ElementsMatch(t, []string{"a"}, StrArrayIntersect(a, b))
	require.ElementsMatch(t, []string{}, StrArrayIntersect(a, c))
}
