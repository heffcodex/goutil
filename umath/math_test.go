package umath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 1, Min(1, 2, 3))
		assert.Equal(t, 1, Min(3, 2, 1))
	})

	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(1), Min(uint(1), uint(2), uint(3)))
		assert.Equal(t, uint(1), Min(uint(3), uint(2), uint(1)))
	})

	t.Run("float", func(t *testing.T) {
		assert.Equal(t, 1.0, Min(1.0, 2.0, 3.0))
		assert.Equal(t, 1.0, Min(3.0, 2.0, 1.0))
	})
}

func TestStdMin(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float64(1.0), StdMin(float32(1.0), float32(2.0), float32(3.0)))
		assert.Equal(t, float64(1.0), StdMin(float32(3.0), float32(2.0), float32(1.0)))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(1.0), StdMin(float64(1.0), float64(2.0), float64(3.0)))
		assert.Equal(t, float64(1.0), StdMin(float64(3.0), float64(2.0), float64(1.0)))
	})
}

func TestMax(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 3, Max(1, 2, 3))
		assert.Equal(t, 3, Max(3, 2, 1))
	})

	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(3), Max(uint(1), uint(2), uint(3)))
		assert.Equal(t, uint(3), Max(uint(3), uint(2), uint(1)))
	})

	t.Run("float", func(t *testing.T) {
		assert.Equal(t, 3.0, Max(1.0, 2.0, 3.0))
		assert.Equal(t, 3.0, Max(3.0, 2.0, 1.0))
	})
}

func TestStdMax(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float64(3.0), StdMax(float32(1.0), float32(2.0), float32(3.0)))
		assert.Equal(t, float64(3.0), StdMax(float32(3.0), float32(2.0), float32(1.0)))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(3.0), StdMax(float64(1.0), float64(2.0), float64(3.0)))
		assert.Equal(t, float64(3.0), StdMax(float64(3.0), float64(2.0), float64(1.0)))
	})
}

func TestAbs(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 1, Abs(1))
		assert.Equal(t, 1, Abs(-1))
	})
}

func TestStdAbs(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float64(1.0), StdAbs(float32(1.0)))
		assert.Equal(t, float64(1.0), StdAbs(float32(-1.0)))
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(1.0), StdAbs(float64(1.0)))
		assert.Equal(t, float64(1.0), StdAbs(float64(-1.0)))
	})
}
