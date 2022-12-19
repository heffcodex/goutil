package types

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Complex interface {
	complex64 | complex128
}

type Number interface {
	Int | Uint | Float | Complex
}

type Char interface {
	byte | rune
}

type ID interface {
	Int | Uint | Float | Char | string | uintptr
}
