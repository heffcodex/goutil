package umap

// MapFn converts map elements.
type MapFn[K, KR comparable, V, VR any] func(k K, v V) (KR, VR)

// FromArray returns a new map from the given slice.
func FromArray[K comparable, V any, S ~[]V](s S, keyFn func(item V) K) map[K]V {
	res := make(map[K]V, len(s))

	for _, item := range s {
		res[keyFn(item)] = item
	}

	return res
}

// Map iterates over all the elements of the given map and applies the given MapFn to each element to construct a new map.
func Map[K, KR comparable, V, VR any, M ~map[K]V](m M, fn MapFn[K, KR, V, VR]) map[KR]VR {
	res := make(map[KR]VR, len(m))

	for k, v := range m {
		kr, vr := fn(k, v)
		res[kr] = vr
	}

	return res
}

// FindAll returns all the values for the given keys in the given map.
func FindAll[K comparable, V any, M ~map[K]V](m M, keys ...K) []V {
	res := make([]V, 0, len(keys))

	for _, key := range keys {
		if item, ok := m[key]; ok {
			res = append(res, item)
		}
	}

	return res
}
