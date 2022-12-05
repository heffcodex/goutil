package goutil

import "strings"

func ArrayFilter[T any](arr []T, filterFn func(item T) bool) []T {
	res := make([]T, 0, len(arr))

	for _, item := range arr {
		if filterFn(item) {
			res = append(res, item)
		}
	}

	return res
}

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

func ArrayIntersect[T comparable](arr1, arr2 []T) []T {
	if len(arr1) == 0 || len(arr2) == 0 {
		return nil
	}

	minArr, maxArr := arr1, arr2
	if len(arr1) > len(arr2) {
		minArr, maxArr = arr2, arr1
	}

	m := make(map[T]struct{}, len(minArr))
	res := make([]T, 0, len(minArr))

	for _, item := range minArr {
		m[item] = struct{}{}
	}

	for _, item := range maxArr {
		if _, ok := m[item]; ok {
			res = append(res, item)
		}
	}

	return res
}

func ArrayUnique[T comparable](arr []T) []T {
	if len(arr) == 0 {
		return nil
	}

	m := make(map[T]struct{}, len(arr))
	for _, item := range arr {
		m[item] = struct{}{}
	}

	res := make([]T, 0, len(m))
	for item := range m {
		res = append(res, item)
	}

	return res
}

func ArrayUniqueStrings[T string](arr []T, caseSensitive bool) []T {
	if len(arr) == 0 {
		return nil
	} else if caseSensitive {
		return ArrayUnique(arr)
	}

	m := make(map[string]T, len(arr))
	for _, item := range arr {
		m[strings.ToLower(string(item))] = item
	}

	res := make([]T, 0, len(m))
	for _, item := range m {
		res = append(res, item)
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