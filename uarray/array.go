package uarray

func Map[T, V any](arr []T, mapFn func(item T) V) []V {
	res := make([]V, 0, len(arr))

	for _, item := range arr {
		res = append(res, mapFn(item))
	}

	return res
}

func FindIndex[T comparable](arr []T, v T) int {
	for i, item := range arr {
		if item == v {
			return i
		}
	}

	return -1
}

func Contains[T comparable](arr []T, v T) bool {
	return FindIndex(arr, v) != -1
}
