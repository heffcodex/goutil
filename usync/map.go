package usync

import (
	"sync"
)

// Map is a generic thread-safe map.
type Map[K comparable, V any] struct{ sync.Map }

// Load return the value stored in the map for a key if it exists.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	_value, _ok := m.Map.Load(key)
	if !_ok {
		var nilV V
		return nilV, false
	}

	value, ok = _value.(V)
	return
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) { m.Map.Store(key, value) }

// LoadOrStore returns the existing value for the key if present, otherwise storing the given one.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	_actual, _loaded := m.Map.LoadOrStore(key, value)
	return _actual.(V), _loaded
}

// Delete deletes the given keys from the map.
func (m *Map[K, V]) Delete(key ...K) {
	for _, k := range key {
		m.Map.Delete(k)
	}
}

// Range iterates over the map and calls the given function for each key and value present in the map.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(
		func(key, value any) bool {
			return f(key.(K), value.(V))
		},
	)
}

// Len returns the number of elements in the map.
// WARNING: this method is not optimal for large maps.
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

// Keys returns the slice keys of the map keys.
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

// Values returns the slice values of the map values.
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
