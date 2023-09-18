package uarray

// TestFn is a basic filtering function.
type TestFn[T any] func(arr []T, i int) bool

func Map[T, V any](arr []T, mapFn func(item T) V) []V {
	res := make([]V, 0, len(arr))

	for _, item := range arr {
		res = append(res, mapFn(item))
	}

	return res
}

// Merge combines elements of provided arrays.
func Merge[T any](arrays ...[]T) []T {
	res := make([]T, 0)

	for _, arr := range arrays {
		res = append(res, arr...)
	}

	return res
}
