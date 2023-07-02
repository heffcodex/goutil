package uarray

import "github.com/heffcodex/goutil/v2/umath"

type CmpFn[T any] func(a, d T) bool

func CmpValue[T comparable](a, d T) bool {
	return a == d
}

func Diff[T any](actual, desired []T, cmp CmpFn[T]) (eq, add, rm []T) {
	if len(actual) == 0 {
		return nil, desired, nil
	}

	if len(desired) == 0 {
		return nil, nil, actual
	}

	eq = make([]T, 0, umath.Max(len(desired), len(actual)))
	add = make([]T, 0, len(desired))
	rm = make([]T, 0, len(actual))

	skipActualIdx := make(map[int]struct{}, len(actual))
	skipDesiredIdx := make(map[int]struct{}, len(desired))

	for ia := range actual {
		if _, ok := skipActualIdx[ia]; ok {
			continue
		}

		found := false

		for id := range desired {
			if _, ok := skipDesiredIdx[id]; ok {
				continue
			}

			if cmp(actual[ia], desired[id]) {
				found = true
				eq = append(eq, desired[id])
				skipActualIdx[ia] = struct{}{}
				skipDesiredIdx[id] = struct{}{}

				break
			}
		}

		if !found {
			rm = append(rm, actual[ia])
		}
	}

	for id := range desired {
		if _, ok := skipDesiredIdx[id]; ok {
			continue
		}

		found := false

		for ia := range actual {
			if _, ok := skipActualIdx[ia]; ok {
				continue
			}

			if cmp(actual[ia], desired[id]) {
				found = true
				eq = append(add, desired[id])
				skipActualIdx[ia] = struct{}{}
				skipDesiredIdx[id] = struct{}{}

				break
			}
		}

		if !found {
			add = append(add, desired[id])
		}
	}

	return eq, add, rm
}
