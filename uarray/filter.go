package uarray

type FilterFn[T any] TestFn[T]

func Unique[T comparable]() FilterFn[T] {
	m := make(map[T]struct{})

	return func(arr []T, i int) bool {
		if _, ok := m[arr[i]]; ok {
			return false
		}

		m[arr[i]] = struct{}{}
		return true
	}
}

func Intersection[T comparable](arr []T) FilterFn[T] {
	if len(arr) == 0 {
		return none[T]()
	}

	m := make(map[T]struct{}, len(arr))
	for _, item := range arr {
		m[item] = struct{}{}
	}

	return func(arr []T, i int) bool {
		_, ok := m[arr[i]]
		return ok
	}
}

func none[T any]() FilterFn[T] {
	return func(arr []T, i int) bool {
		return false
	}
}

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
