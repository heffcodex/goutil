package uarray

type FilterFn[T any] func(item T) bool

func Unique[T comparable]() FilterFn[T] {
	m := make(map[T]struct{})

	return func(item T) bool {
		if _, ok := m[item]; ok {
			return false
		}

		m[item] = struct{}{}
		return true
	}
}

func Intersection[T comparable](arr []T) FilterFn[T] {
	if len(arr) == 0 {
		return None[T]()
	}

	m := make(map[T]struct{}, len(arr))
	for _, item := range arr {
		m[item] = struct{}{}
	}

	return func(item T) bool {
		_, ok := m[item]
		return ok
	}
}

func All[T any]() FilterFn[T] {
	return func(item T) bool {
		return true
	}
}

func None[T any]() FilterFn[T] {
	return func(item T) bool {
		return false
	}
}

func Filter[T any](arr []T, fn ...FilterFn[T]) []T {
	res := make([]T, 0, len(arr))

	for _, item := range arr {
		for _, f := range fn {
			if !f(item) {
				goto next
			}
		}

		res = append(res, item)
	next:
	}

	return res
}