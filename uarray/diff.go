package uarray

type CmpFn[T any] func(a, d T) bool

func CmpValue[T comparable](a, d T) bool {
	return a == d
}

func Diff[T any](actual, desired []T, cmp CmpFn[T]) (add, rm []T) {
	if len(actual) == 0 {
		return desired, nil
	}

	if len(desired) == 0 {
		return nil, actual
	}

	add = make([]T, 0, len(desired))
	rm = make([]T, 0, len(actual))

	skipActualIdx := make(map[int]struct{}, len(actual))
	skipDesiredIdx := make(map[int]struct{}, len(desired))

	for ia, a := range actual {
		if _, ok := skipActualIdx[ia]; ok {
			continue
		}

		found := false

		for id, d := range desired {
			if _, ok := skipDesiredIdx[id]; ok {
				continue
			}

			if cmp(a, d) {
				found = true
				skipActualIdx[ia] = struct{}{}
				skipDesiredIdx[id] = struct{}{}

				break
			}
		}

		if !found {
			rm = append(rm, a)
		}
	}

	for id, d := range desired {
		if _, ok := skipDesiredIdx[id]; ok {
			continue
		}

		found := false

		for ia, a := range actual {
			if _, ok := skipActualIdx[ia]; ok {
				continue
			}

			if cmp(a, d) {
				found = true
				skipActualIdx[ia] = struct{}{}
				skipDesiredIdx[id] = struct{}{}

				break
			}
		}

		if !found {
			add = append(add, d)
		}
	}

	return add, rm
}
