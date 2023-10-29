package umap

// FromArray returns a new map from the given slice.
func FromArray[K comparable, V any, S ~[]V](s S, keyFn func(item V) K) map[K]V {
	res := make(map[K]V, len(s))

	for _, item := range s {
		res[keyFn(item)] = item
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
