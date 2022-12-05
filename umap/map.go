package umap

func FromArray[K comparable, V any](arr []V, keyFn func(item V) K) map[K]V {
	res := make(map[K]V, len(arr))

	for _, item := range arr {
		res[keyFn(item)] = item
	}

	return res
}
