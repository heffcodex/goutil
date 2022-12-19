package uvalue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRef(t *testing.T) {
	require.Equal(t, 1, *Ref(1))
	require.Equal(t, (*int)(nil), *Ref((*int)(nil)))
}

func TestRefOrNil(t *testing.T) {
	require.Equal(t, (*int)(nil), RefOrNil(0))
	require.Equal(t, 1, *RefOrNil(1))
}
