package umath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, 1, Abs(1))
		assert.Equal(t, 1, Abs(-1))
	})
}

func TestStdAbs(t *testing.T) {
	t.Parallel()

	t.Run("float32", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, float64(1.0), StdAbs(float32(1.0)))  //nolint:testifylint // float-compare: not now
		assert.Equal(t, float64(1.0), StdAbs(float32(-1.0))) //nolint:testifylint // float-compare: not now
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, float64(1.0), StdAbs(float64(1.0)))  //nolint:testifylint // float-compare: not now
		assert.Equal(t, float64(1.0), StdAbs(float64(-1.0))) //nolint:testifylint // float-compare: not now
	})
}
