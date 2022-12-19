package uvalue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type invertZero bool

func (iz invertZero) IsZero() bool { return !bool(iz) }

func TestIsZero(t *testing.T) {
	require.Equal(t, true, IsZero((*int)(nil)))
	require.Equal(t, true, IsZero(""))
	require.Equal(t, true, IsZero(0))

	require.Equal(t, false, IsZero("0"))
	require.Equal(t, false, IsZero(1))

	require.Equal(t, true, IsZero(invertZero(false)))
	require.Equal(t, false, IsZero(invertZero(true)))
}
