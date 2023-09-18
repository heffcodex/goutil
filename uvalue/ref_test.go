package uvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRef(t *testing.T) {
	assert.Equal(t, 1, *Ref(1))
	assert.Equal(t, (*int)(nil), *Ref((*int)(nil)))
}

func TestRefOrNil(t *testing.T) {
	assert.Equal(t, (*int)(nil), RefOrNil(0))
	assert.Equal(t, 1, *RefOrNil(1))
}

func TestCopyRef(t *testing.T) {
	v := 1

	assert.Equal(t, 1, *CopyRef(&v))
	assert.Nil(t, CopyRef((*int)(nil)))
}
