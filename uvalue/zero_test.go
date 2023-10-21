package uvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type invertZero bool

func (iz invertZero) IsZero() bool { return !bool(iz) }

func TestIsZero(t *testing.T) {
	assert.True(t, IsZero((*int)(nil)))
	assert.True(t, IsZero(""))
	assert.True(t, IsZero(0))

	assert.False(t, IsZero("0"))
	assert.False(t, IsZero(1))

	assert.True(t, IsZero(invertZero(false)))
	assert.False(t, IsZero(invertZero(true)))
}
