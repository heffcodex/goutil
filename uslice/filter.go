package uslice

// FilterFn returns true if the element should be included in the resulting slice.
type FilterFn[V any] func(item V) bool

// Unique returns FilterFn that returns `true` if the element is unique.
func Unique[K comparable, V any](keyFn KeyFn[K, V]) FilterFn[V] {
	m := make(map[K]struct{})

	return func(v V) bool {
		k := keyFn(v)
		if _, ok := m[k]; ok {
			return false
		}

		m[k] = struct{}{}

		return true
	}
}

// Intersection returns FilterFn that returns `true` if the element is present in both filtering slice and the provided one.
func Intersection[K comparable, V any, S ~[]V](s S, keyFn KeyFn[K, V]) FilterFn[V] {
	if len(s) == 0 {
		return func(_ V) bool {
			return false
		}
	}

	m := make(map[K]struct{}, len(s))

	for _, v := range s {
		m[keyFn(v)] = struct{}{}
	}

	return func(v V) bool {
		_, ok := m[keyFn(v)]
		return ok
	}
}

// Filter performs filtering on the slice.
// Given filters are applied in the order they are provided with AND logic.
func Filter[V any, S ~[]V](s S, fn ...FilterFn[V]) S {
	res := make(S, 0, len(s))

	applyFilters(s, fn, func(item V) {
		res = append(res, item)
	})

	return res
}

// Count returns the number of filtered elements.
// Given filters are applied in the order they are provided with AND logic.
func Count[V any, S ~[]V](s S, fn ...FilterFn[V]) int {
	count := 0

	applyFilters(s, fn, func(V) {
		count++
	})

	return count
}

// applyFilters applies the given filters to the given slice and calls the given function on each succeed element.
// Given filters are applied in the order they are provided with AND logic.
func applyFilters[V any, S ~[]V](s S, fns []FilterFn[V], onMatch func(item V)) {
	for _, item := range s {
		for _, fn := range fns {
			if !fn(item) {
				goto next
			}
		}

		onMatch(item)
	next:
	}
}
