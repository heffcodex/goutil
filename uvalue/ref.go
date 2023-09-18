package uvalue

// Ref returns a pointer to a copy (!) the value v.
func Ref[T any](v T) *T {
	return &v
}

// RefOrNil returns a pointer to a copy (!) the value v or nil if v is the zero value.
// Works only with comparable types due to the current limitations of the language.
// WARNING: this invokes the IsZero function for the value v which may use slow reflection to determine if the value is actually zero.
func RefOrNil[T comparable](v T) *T {
	if IsZero(v) {
		return nil
	}

	return Ref(v)
}

// CopyRef returns a pointer to a copy (!) the value pointed by v or nil if v is nil.
func CopyRef[T any](v *T) *T {
	if v == nil {
		return nil
	}

	return Ref(*v)
}
