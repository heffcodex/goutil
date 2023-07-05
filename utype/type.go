package utype

import (
	"golang.org/x/exp/constraints"
)

// Int is a generic signed integer type.
type Int interface {
	constraints.Signed
}

// Uint is a generic unsigned integer type.
type Uint interface {
	constraints.Unsigned
}

// Integer is a generic integer type.
type Integer interface {
	Int | Uint
}

// Float is a generic floating point type.
type Float interface {
	constraints.Float
}

// Real is a generic type containing both floating point and integer types.
type Real interface {
	Integer | Float
}

// Complex is a generic complex type.
type Complex interface {
	constraints.Complex
}

// Number is a generic type containing both Real and Complex types.
type Number interface {
	Real | Complex
}

// Char is a generic character type.
type Char interface {
	~byte | ~rune
}

// ID is a generic comparable identifier type.
type ID interface {
	Real | Char | ~string | ~uintptr
}
