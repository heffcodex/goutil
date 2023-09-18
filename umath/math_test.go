package umath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
