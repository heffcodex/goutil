package uarray

// FromMapFn converts a map element to an array element.
type FromMapFn[K comparable, V, R any] func(k K, v V) R

// UseKeys implements FromMapFn to use the keys of the map as the array elements.
func UseKeys[K comparable, V any]() FromMapFn[K, V, K] {
	return func(k K, v V) K {
		return k
	}
}

// UseValues implements FromMapFn to use the values of the map as the array elements.
func UseValues[K comparable, V any]() FromMapFn[K, V, V] {
	return func(k K, v V) V {
		return v
	}
}

// FromMap converts a map to an array.
func FromMap[K comparable, V, R any](m map[K]V, fn FromMapFn[K, V, R]) []R {
	res := make([]R, 0, len(m))

	for k, v := range m {
		res = append(res, fn(k, v))
	}

	return res
}
