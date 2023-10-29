package uslice

// FromMapFn converts a map element to an slice element.
type FromMapFn[K comparable, V, R any] func(k K, v V) R

// UseKeys implements FromMapFn to use the keys of the map as the slice elements.
func UseKeys[K comparable, V any](k K, _ V) K {
	return k
}

// UseValues implements FromMapFn to use the values of the map as the slice elements.
func UseValues[K comparable, V any](_ K, v V) V {
	return v
}

// FromMap converts a map to a slice with the given FromMapFn converter.
func FromMap[K comparable, V, R any, M ~map[K]V](m M, fn FromMapFn[K, V, R]) []R {
	res := make([]R, len(m))
	i := 0

	for k, v := range m {
		res[i] = fn(k, v)
		i++
	}

	return res
}
