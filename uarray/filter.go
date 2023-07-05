package uarray

import "github.com/heffcodex/goutil/v2/utype"

type (
	// FilterFn returns true if the element should be included in the resulting array.
	FilterFn[T any] TestFn[T]

	// IDFn returns a unique identifier for the element.
	IDFn[T any, V utype.ID] func(arr []T, i int) V
)

// IDValue implements IDFn to use the value as the ID.
func IDValue[T utype.ID](arr []T, i int) T { return arr[i] }

// Unique returns FilterFn that returns true if the element is unique.
func Unique[T comparable, V utype.ID](id IDFn[T, V]) FilterFn[T] {
	m := make(map[V]struct{})

	return func(arr []T, i int) bool {
		vID := id(arr, i)
		if _, ok := m[vID]; ok {
			return false
		}

		m[vID] = struct{}{}
		return true
	}
}

// Intersection returns FilterFn that returns true if the element is present in both filtering array and the provided one.
func Intersection[T comparable, V utype.ID](arr []T, id IDFn[T, V]) FilterFn[T] {
	if len(arr) == 0 {
		return none[T]()
	}

	m := make(map[V]struct{}, len(arr))
	for i := range arr {
		m[id(arr, i)] = struct{}{}
	}

	return func(arr []T, i int) bool {
		_, ok := m[id(arr, i)]
		return ok
	}
}

// none returns FilterFn that returns false for any element.
func none[T any]() FilterFn[T] {
	return func(arr []T, i int) bool {
		return false
	}
}

// Filter performs filtering on the array.
func Filter[T any](arr []T, fn ...FilterFn[T]) []T {
	res := make([]T, 0, len(arr))

	for i := range arr {
		for _, f := range fn {
			if !f(arr, i) {
				goto next
			}
		}

		res = append(res, arr[i])
	next:
	}

	return res
}

// Count returns the number of filtered elements.
func Count[T any](arr []T, fn ...FilterFn[T]) int {
	count := 0

	for i := range arr {
		for _, f := range fn {
			if !f(arr, i) {
				goto next
			}
		}

		count++
	next:
	}

	return count
}
