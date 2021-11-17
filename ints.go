package goutil

func Uint64RefOrNil(v uint64) *uint64 {
	if v > 0 {
		return &v
	}

	return nil
}

func Float64Ref(f float64) *float64 {
	return &f
}
