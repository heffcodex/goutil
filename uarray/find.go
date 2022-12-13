package uarray

const IdxNotFound = -1

type MatchFn[T any] TestFn[T]

func Value[T comparable](v T) MatchFn[T] {
	return func(item T) bool {
		return item == v
	}
}

func FindIndex[T any](arr []T, f MatchFn[T]) int {
	for i, item := range arr {
		if f(item) {
			return i
		}
	}

	return IdxNotFound
}

func Find[T any](arr []T, f MatchFn[T]) (T, bool) {
	idx := FindIndex(arr, f)
	if idx == IdxNotFound {
		return *new(T), false
	}

	return arr[idx], true
}

func Contains[T any](arr []T, f MatchFn[T]) bool {
	return FindIndex(arr, f) != IdxNotFound
}
