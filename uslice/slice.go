package uslice

// MapFn converts slice elements.
type MapFn[V, R any] func(item V) R

// Map iterates over all the elements of the given slices and applies the given function to each element.
// Returns a new merged (see Merge) slice with the results.
func Map[V, R any, S ~[]V](fn MapFn[V, R], s ...S) []R {
	return mapMerge(fn, s...)
}

// Merge combines elements of the given slices into a single slice with appending.
// Returns a new slice with the results.
func Merge[V any, S ~[]V](s ...S) S {
	return mapMerge(func(item V) V { return item }, s...)
}

// Len returns an overall length of the given slices.
func Len[V any, S ~[]V](s ...S) int {
	l := 0

	for _, _s := range s {
		l += len(_s)
	}

	return l
}

func mapMerge[V, R any, S ~[]V](fn MapFn[V, R], s ...S) []R {
	res := make([]R, Len(s...))
	i := 0

	for _, _s := range s {
		for si := range _s {
			res[i] = fn(_s[si])
			i++
		}
	}

	return res
}
