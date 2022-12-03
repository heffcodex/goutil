package goutil

import (
	"reflect"
)

var zeroerT = reflect.TypeOf((*zeroer)(nil)).Elem()

type zeroer interface {
	IsZero() bool
}

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

func Ref[T any](v T) *T {
	return &v
}

func RefOrNil[T comparable](v T) *T {
	if IsZero(v) {
		return nil
	}

	return Ref(v)
}
