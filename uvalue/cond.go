package uvalue

// PassIf returns true if the value satisfies the `test` function.
func PassIf[T any](obj T, test func() bool) T {
	if test() {
		return obj
	}

	return *new(T)
}
