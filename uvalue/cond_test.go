package uvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassIf(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, PassIf(1, func() bool { return true }))
	assert.Equal(t, 0, PassIf(1, func() bool { return false }))
}
