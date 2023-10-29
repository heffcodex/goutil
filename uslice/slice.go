package uslice

// MapFn converts slice elements.
type MapFn[V, R any] func(item V) R

// Map iterates over all the elements of the given slice and applies the given function to each element.
func Map[V, R any, S ~[]V](s S, fn MapFn[V, R]) []R {
	res := make([]R, Len(s))

	for i, item := range s {
		res[i] = fn(item)
	}

	return res
}

// Merge combines elements of the given slices into a single slice with appending.
// Returns a new slice with the results.
func Merge[V any, S ~[]V](s ...S) S {
	res := make(S, Len(s...))
	i := 0

	for _, _s := range s {
		for si := range _s {
			res[i] = _s[si]
			i++
		}
	}

	return res
}

// Len returns an overall length of the given slices.
func Len[V any, S ~[]V](s ...S) int {
	l := 0

	for _, _s := range s {
		l += len(_s)
	}

	return l
}
