package uarray

const IdxNotFound = -1

// MatchFn returns true if the element is found in the array.
type MatchFn[T any] TestFn[T]

// Value returns MatchFn to match element by the provided value.
func Value[T comparable](value T) MatchFn[T] {
	return func(arr []T, i int) bool {
		return arr[i] == value
	}
}

// Any returns MatchFn to match any of the provided filters.
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

// AnyValue returns MatchFn to match any element from the provided values.
func AnyValue[T comparable](values ...T) MatchFn[T] {
	return func(arr []T, i int) bool {
		return Any(mapValues(values)...)(arr, i)
	}
}

// mapValues is a helper function for AnyValue.
func mapValues[T comparable](values []T) []MatchFn[T] { return Map(values, Value[T]) }

// FindIndex returns the index of the first element that matches the provided MatchFn.
func FindIndex[T any](arr []T, f MatchFn[T]) int {
	for i := range arr {
		if f(arr, i) {
			return i
		}
	}

	return IdxNotFound
}

// Find returns the first element that matches the provided MatchFn.
func Find[T any](arr []T, f MatchFn[T]) (T, bool) {
	idx := FindIndex(arr, f)
	if idx == IdxNotFound {
		return *new(T), false
	}

	return arr[idx], true
}

// Contains returns true if the array contains the provided element.
func Contains[T any](arr []T, f MatchFn[T]) bool {
	return FindIndex(arr, f) != IdxNotFound
}
