package uarray

type FromMapFn[K comparable, V, R any] func(k K, v V) R

func FromMap[K comparable, V, R any](m map[K]V, fn FromMapFn[K, V, R]) []R {
	res := make([]R, 0, len(m))

	for k, v := range m {
		res = append(res, fn(k, v))
	}

	return res
}

func UseKeys[K comparable, V any]() FromMapFn[K, V, K] {
	return func(k K, v V) K {
		return k
	}
}

func UseValues[K comparable, V any]() FromMapFn[K, V, V] {
	return func(k K, v V) V {
		return v
	}
}
