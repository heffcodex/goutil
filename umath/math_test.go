package umath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		require.Equal(t, 1, Min(1, 2))
		require.Equal(t, 1, Min(2, 1))
	})

	t.Run("uint", func(t *testing.T) {
		require.Equal(t, uint(1), Min(uint(1), uint(2)))
		require.Equal(t, uint(1), Min(uint(2), uint(1)))
	})

	t.Run("float", func(t *testing.T) {
		require.Equal(t, 1.0, Min(1.0, 2.0))
		require.Equal(t, 1.0, Min(2.0, 1.0))
	})
}

func TestMax(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		require.Equal(t, 2, Max(1, 2))
		require.Equal(t, 2, Max(2, 1))
	})

	t.Run("uint", func(t *testing.T) {
		require.Equal(t, uint(2), Max(uint(1), uint(2)))
		require.Equal(t, uint(2), Max(uint(2), uint(1)))
	})

	t.Run("float", func(t *testing.T) {
		require.Equal(t, 2.0, Max(1.0, 2.0))
		require.Equal(t, 2.0, Max(2.0, 1.0))
	})
}

func TestAbs(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		require.Equal(t, 1, Abs(1))
		require.Equal(t, 1, Abs(-1))
	})
}
