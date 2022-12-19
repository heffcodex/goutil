package uarray

const IdxNotFound = -1

type MatchFn[T any] TestFn[T]

func Value[T comparable](value T) MatchFn[T] {
	return func(arr []T, i int) bool {
		return arr[i] == value
	}
}

func Not[T comparable](f MatchFn[T]) MatchFn[T] {
	return func(arr []T, i int) bool {
		return !f(arr, i)
	}
}

func Any[T comparable](fn ...MatchFn[T]) MatchFn[T] {
	return func(arr []T, i int) bool {
		for _, f := range fn {
			if f(arr, i) {
				return true
			}
		}

		return false
	}
}

func AnyValue[T comparable](values ...T) MatchFn[T] {
	return func(arr []T, i int) bool {
		return Any(mapValues(values)...)(arr, i)
	}
}

func mapValues[T comparable](values []T) []MatchFn[T] { return Map(values, Value[T]) }

func FindIndex[T any](arr []T, f MatchFn[T]) int {
	for i := range arr {
		if f(arr, i) {
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
