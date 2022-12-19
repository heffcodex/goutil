package umath

import (
	"math"
	"unsafe"

	"github.com/heffcodex/goutil/v2/utype"
)

func Min[T utype.Real](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Max[T utype.Real](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// Abs returns the absolute value of x.
// Applicable to all signed integer types. For floating point types, see StdAbs.
func Abs[T utype.Int](n T) T {
	mask := n >> unsafe.Sizeof(n)
	return (mask + n) ^ mask
}

// StdAbs returns the absolute value of x for floating point types.
// Use Abs for signed integer types.
func StdAbs[T utype.Float](n T) float64 {
	return math.Abs(float64(n))
}
