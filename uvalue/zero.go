package uvalue

import (
	"reflect"
)

var zeroerT = reflect.TypeOf((*zeroer)(nil)).Elem()

type zeroer interface {
	IsZero() bool
}

// IsZero returns true if the value is nil or its underlying value is equal to the zero value for its type.
// Works only with comparable types due to the current limitations of the language.
// Warning: this function may use slow reflection.
func IsZero[T comparable](v T) bool {
	if v == *new(T) {
		return true
	}

	vof := reflect.Indirect(reflect.ValueOf(v))
	if !vof.IsValid() || vof.IsZero() {
		return true
	}

	if vof.CanConvert(zeroerT) {
		return vof.Convert(zeroerT).Interface().(zeroer).IsZero()
	}

	return false
}
