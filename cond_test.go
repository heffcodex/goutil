package goutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTruePercent(t *testing.T) {
	require.Equal(t, 0, TruePercent(false, false, false))
	require.Equal(t, 100, TruePercent(true, true, true))
	require.Equal(t, 67, TruePercent(true, true, false))
}
