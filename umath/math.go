package umath

import (
	"math"
	"unsafe"

	"github.com/heffcodex/goutil/v2/utype"
)

// Min returns the smallest of the provided values.
// Applicable to all integer utype.Real-like types, but does not use go builtin arch optimisations.
// For optimised version accepting only floating point utype.Float-like types, see StdMin.
func Min[T utype.Real](v ...T) T {
	if len(v) < 2 {
		panic("at least 2 arguments required")
	}

	m := minTwo(v[0], v[1])

	for i := 2; i < len(v); i++ {
		m = minTwo(m, v[i])
	}

	return m
}

func minTwo[T utype.Real](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// StdMin returns the smallest of the provided values for floating point utype.Float-like types only.
// It uses go builtin arch optimisations if possible, calling math.Min internally.
// If you need to compare only 2 values, it is strongly recommended to use math.Min directly instead.
func StdMin[T utype.Float](v ...T) float64 {
	if len(v) < 2 {
		panic("at least 2 arguments required")
	}

	m := math.Min(float64(v[0]), float64(v[1]))

	for i := 2; i < len(v); i++ {
		m = math.Min(m, float64(v[i]))
	}

	return m
}

// Max returns the biggest of the provided values.
// Applicable to all integer utype.Real-like types, but does not use go builtin arch optimisations.
// For optimised version accepting only floating point utype.Float-like types, see StdMax.
func Max[T utype.Real](v ...T) T {
	if len(v) < 2 {
		panic("at least 2 arguments required")
	}

	m := maxTwo(v[0], v[1])

	for i := 2; i < len(v); i++ {
		m = maxTwo(m, v[i])
	}

	return m
}

func maxTwo[T utype.Real](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// StdMax returns the biggest of the provided values for floating point utype.Float-like types only.
// It uses go builtin arch optimisations if possible, calling math.Max internally.
// If you need to compare only 2 values, it is strongly recommended to use math.Max directly instead.
func StdMax[T utype.Float](v ...T) float64 {
	if len(v) < 2 {
		panic("at least 2 arguments required")
	}

	m := math.Max(float64(v[0]), float64(v[1]))

	for i := 2; i < len(v); i++ {
		m = math.Max(m, float64(v[i]))
	}

	return m
}

// Abs returns the absolute value of x.
// Applicable to all signed integer utype.Int-like types.
// For floating point utype.Float-like types only, see StdAbs.
func Abs[T utype.Int](n T) T {
	mask := n >> unsafe.Sizeof(n)
	return (mask + n) ^ mask
}

// StdAbs returns the absolute value of x for floating point utype.Float-like types only.
// Use Abs for signed integer utype.Int-like types.
func StdAbs[T utype.Float](n T) float64 {
	return math.Abs(float64(n))
}
