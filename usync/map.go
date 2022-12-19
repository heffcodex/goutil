package usync

import (
	"sync"
)

type Map[K comparable, V any] struct{ sync.Map }

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	_value, _ok := m.Map.Load(key)
	if !_ok {
		var nilV V
		return nilV, false
	}

	value, ok = _value.(V)
	return
}

func (m *Map[K, V]) Store(key K, value V) { m.Map.Store(key, value) }

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	_actual, _loaded := m.Map.LoadOrStore(key, value)
	return _actual.(V), _loaded
}

func (m *Map[K, V]) Delete(key ...K) {
	for _, k := range key {
		m.Map.Delete(k)
	}
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(
		func(key, value any) bool {
			return f(key.(K), value.(V))
		},
	)
}

func (m *Map[K, V]) Len() int {
	l := 0

	m.Range(
		func(K, V) bool {
			l++
			return true
		},
	)

	return l
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0)

	m.Range(
		func(key K, _ V) bool {
			keys = append(keys, key)
			return true
		},
	)

	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0)

	m.Range(
		func(_ K, value V) bool {
			values = append(values, value)
			return true
		},
	)

	return values
}
