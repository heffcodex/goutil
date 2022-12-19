package uvalue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassIf(t *testing.T) {
	require.Equal(t, 1, PassIf(1, func() bool { return true }))
	require.Equal(t, 0, PassIf(1, func() bool { return false }))
}
