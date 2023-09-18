package umath

import (
	"math"
	"unsafe"

	"github.com/heffcodex/goutil/v2/utype"
)

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
