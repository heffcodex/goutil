package utype

import (
	"golang.org/x/exp/constraints"
)

type Int interface {
	constraints.Signed
}

type Uint interface {
	constraints.Unsigned
}

type Integer interface {
	Int | Uint
}

type Float interface {
	constraints.Float
}

type Real interface {
	Integer | Float
}

type Complex interface {
	constraints.Complex
}

type Number interface {
	Real | Complex
}

type Char interface {
	~byte | ~rune
}

type ID interface {
	Real | Char | ~string | ~uintptr
}
