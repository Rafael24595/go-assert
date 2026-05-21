package internal

import (
	"reflect"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
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
