package goutil

func ArrayPluck[T, V any](arr []T, valueFn func(item T) V) []V {
	res := make([]V, 0, len(arr))

	for _, item := range arr {
		res = append(res, valueFn(item))
	}

	return res
}

func ArrayMap[T, V any](arr []T, mapFn func(item T) V) []V {
	res := make([]V, 0, len(arr))

	for _, item := range arr {
		res = append(res, mapFn(item))
	}

	return res
}

func InArray[T comparable](arr []T, v T) bool {
	for _, item := range arr {
		if item == v {
			return true
		}
	}

	return false
}
