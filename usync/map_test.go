package usync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_Load(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)

	v, ok := m.Load(1)
	require.True(t, ok)
	require.Equal(t, 1, v)

	v, ok = m.Load(2)
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestMap_Store(t *testing.T) {
	m := Map[int, int]{}

	m.Store(1, 1)

	v, ok := m.Map.Load(1)
	require.True(t, ok)
	require.Equal(t, 1, v)

	m.Store(1, 2)

	v, ok = m.Map.Load(1)
	require.True(t, ok)
	require.Equal(t, 2, v)
}

func TestMap_LoadOrStore(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)

	v, ok := m.LoadOrStore(1, 1)
	require.True(t, ok)
	require.Equal(t, 1, v)

	v, ok = m.LoadOrStore(2, 2)
	require.False(t, ok)
	require.Equal(t, 2, v)
}

func TestMap_Delete(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)
	m.Map.Store(2, 2)

	m.Delete(1)

	v, ok := m.Map.Load(1)
	require.False(t, ok)
	require.Equal(t, nil, v)

	v, ok = m.Map.Load(2)
	require.True(t, ok)
	require.Equal(t, 2, v)
}

func TestMap_Range(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)
	m.Map.Store(2, 2)

	m.Range(
		func(key int, value int) bool {
			require.True(t, key == 1 || key == 2)
			require.True(t, value == 1 || value == 2)
			return true
		},
	)
}

func TestMap_Len(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)
	m.Map.Store(2, 2)

	require.Equal(t, 2, m.Len())
}

func TestMap_Keys(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)
	m.Map.Store(2, 2)

	keys := m.Keys()
	require.EqualValues(t, []int{1, 2}, keys)
}

func TestMap_Values(t *testing.T) {
	m := Map[int, int]{}

	m.Map.Store(1, 1)
	m.Map.Store(2, 2)

	values := m.Values()
	require.EqualValues(t, []int{1, 2}, values)
}
