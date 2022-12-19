package uvalue

func PassIf[T any](obj T, test func() bool) T {
	if test() {
		return obj
	}

	return *new(T)
}
