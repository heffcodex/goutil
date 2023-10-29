package uslice

// KeyFn provides a way to extract comparable keys from slice elements.
// For example, it may be ID of the struct or even the element itself (see KeyValue) for primitive types,
// but it is important to keep these keys short to avoid memory bloat.
type KeyFn[K comparable, V any] func(v V) K

// KeyValue is a KeyFn that uses provided value itself as a key.
func KeyValue[V comparable](v V) V {
	return v
}

// Diff works just like DiffIndex and on a base of it, but returns the resulting slice elements of the difference.
// See the corresponding function for more details.
func Diff[K comparable, V any](actual, desired []V, keyFn KeyFn[K, V]) (eq, rm, add []V) {
	if len(actual) == 0 {
		return nil, nil, desired
	}

	if len(desired) == 0 {
		return nil, actual, nil
	}

	iEq, iRm, iAdd := DiffIndex(actual, desired, keyFn)

	eq = make([]V, len(iEq))
	rm = make([]V, len(iRm))
	add = make([]V, len(iAdd))

	maxLen := max(len(eq), len(rm), len(add))

	for i := 0; i < maxLen; i++ {
		if len(eq) > i {
			eq[i] = actual[iEq[i]]
		}

		if len(rm) > i {
			rm[i] = actual[iRm[i]]
		}

		if len(add) > i {
			add[i] = desired[iAdd[i]]
		}
	}

	return eq, rm, add
}

// DiffIndex returns indices of the elements against the provided `actual` and `desired`:
// `eq` - the indices of the elements of `actual` slice that are _equal_ to the desired (i.e. may stay at their current place).
// `rm` - the indices of the elements of `actual` slice that are _not equal_ to the desired (i.e. needs to be removed from `actual`).
// `add` - the indices of the elements of `desired` slice that are _not equal_ to the actual (i.e. needs to be appended to `actual`).
func DiffIndex[K comparable, V any](actual, desired []V, keyFn KeyFn[K, V]) (eq, rm, add []int) {
	if len(actual) == 0 {
		return nil, nil, series(len(desired))
	}

	if len(desired) == 0 {
		return nil, series(len(actual)), nil
	}

	eq = make([]int, 0, min(len(desired), len(actual)))
	rm = make([]int, 0, len(actual))
	add = make([]int, 0, len(desired))

	actualKeys := make([]*K, len(actual))
	actualKIdxSame := make(map[K][]int, len(actual))
	desiredKIdxSame := make(map[K][]int, len(desired))

	for i := range actual {
		k := keyFn(actual[i])

		actualKeys[i] = &k
		actualKIdxSame[k] = append(actualKIdxSame[k], i)
	}

	for i := range desired {
		k := keyFn(desired[i])
		sameActual := actualKIdxSame[k]
		sameDesired := append(desiredKIdxSame[k], i) //nolint: gocritic // appendAssign: false positive

		if len(sameActual) < len(sameDesired) {
			add = append(add, i)
		}

		desiredKIdxSame[k] = sameDesired
	}

	for i := range actual {
		pk := actualKeys[i]

		sameActual, ok := actualKIdxSame[*pk]
		if !ok {
			continue
		}

		sameDesired := desiredKIdxSame[*pk]

		if len(sameActual) <= len(sameDesired) {
			eq = append(eq, sameActual...)
		} else {
			eq = append(eq, sameActual[:len(sameDesired)]...)
			rm = append(rm, sameActual[len(sameDesired):]...)
		}

		delete(actualKIdxSame, *pk)
	}

	return eq, rm, add
}

// series returns a slice of ints filled with values equal to its index.
func series(n int) []int {
	if n == 0 {
		return nil
	}

	s := make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i
	}

	return s
}
