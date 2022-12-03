package goutil

import (
	"sync"
)

type SyncMap[K comparable, V any] struct{ sync.Map }

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	_value, _ok := m.Map.Load(key)
	if !_ok {
		var nilV V
		return nilV, false
	}

	value, ok = _value.(V)
	return
}

func (m *SyncMap[K, V]) Store(key K, value V) { m.Map.Store(key, value) }

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	_actual, _loaded := m.Map.LoadOrStore(key, value)
	return _actual.(V), _loaded
}

func (m *SyncMap[K, V]) Delete(key ...K) {
	for _, k := range key {
		m.Map.Delete(k)
	}
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(
		func(key, value any) bool {
			return f(key.(K), value.(V))
		},
	)
}

func (m *SyncMap[K, V]) Len() int {
	l := 0

	m.Range(
		func(K, V) bool {
			l++
			return true
		},
	)

	return l
}

func (m *SyncMap[K, V]) Keys() []K {
	keys := make([]K, 0)

	m.Range(
		func(key K, _ V) bool {
			keys = append(keys, key)
			return true
		},
	)

	return keys
}

func (m *SyncMap[K, V]) Values() []V {
	values := make([]V, 0)

	m.Range(
		func(_ K, value V) bool {
			values = append(values, value)
			return true
		},
	)

	return values
}
