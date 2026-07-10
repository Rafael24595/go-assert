package internal

import (
	"reflect"
)

// Number represents any built-in numeric type supported by this package.
//
// It includes all signed and unsigned integer types, uintptr, and
// floating-point types.
type Number interface {
	Signed | Unsigned | Float
}

// Signed represents any built-in signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned represents any built-in unsigned integer type, including uintptr.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float represents any built-in floating-point type.
type Float interface {
	~float32 | ~float64
}

func ToNumber[T Number](want T, have any) (T, bool) {
	v := reflect.ValueOf(have)

	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan, reflect.String:
		return T(v.Len()), true
	default:
		target := reflect.TypeOf(want)

		if !v.Type().ConvertibleTo(target) {
			var zero T
			return zero, false
		}

		return v.Convert(target).Interface().(T), true
	}
}
