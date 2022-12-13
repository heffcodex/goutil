package uarray

type TestFn[T any] func(item T) bool

func FromMap[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))

	for _, item := range m {
		res = append(res, item)
	}

	return res
}

func Map[T, V any](arr []T, mapFn func(item T) V) []V {
	res := make([]V, 0, len(arr))

	for _, item := range arr {
		res = append(res, mapFn(item))
	}

	return res
}
