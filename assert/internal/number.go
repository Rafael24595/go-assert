package internal

import (
	"cmp"
	"math"
	"reflect"
)

const maxExactFloatInteger = 1 << 53

type numberKind uint8

const (
	kindSigned numberKind = iota
	kindUnsigned
	kindFloating
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

type magnitude struct {
	kind numberKind
	i    int64
	u    uint64
	f    float64
}

func signedMagnitude[T Signed](v T) magnitude {
	return magnitude{
		kind: kindSigned,
		i:    int64(v),
	}
}

func unsignedMagnitude[T Unsigned](v T) magnitude {
	return magnitude{
		kind: kindUnsigned,
		u:    uint64(v),
	}
}

func floatMagnitude[T Float](v T) magnitude {
	return magnitude{
		kind: kindFloating,
		f:    float64(v),
	}
}

func measureMagnitude(v any) (magnitude, bool) {
	rv := dereferencePointer(
		reflect.ValueOf(v),
	)

	switch rv.Kind() {
	case reflect.Array,
		reflect.Slice,
		reflect.Map,
		reflect.String,
		reflect.Chan:

		return unsignedMagnitude(
			uint64(rv.Len()),
		), true
	}

	return magnitude{}, false
}

func extractMagnitude(v any) (magnitude, bool) {
	rv := dereferencePointer(
		reflect.ValueOf(v),
	)

	switch rv.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:

		return signedMagnitude(
			rv.Int(),
		), true

	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:

		return unsignedMagnitude(
			rv.Uint(),
		), true

	case reflect.Float32,
		reflect.Float64:

		return floatMagnitude(
			rv.Float(),
		), true
	}

	return magnitude{}, false
}

func magnitudeFrom(v any) (magnitude, bool) {
	if n, ok := extractMagnitude(v); ok {
		return n, true
	}

	return measureMagnitude(v)
}

// MagnitudeOf returns the comparable magnitude of v.
//
// For numeric values, it returns the numeric value itself.
//
// For measurable values (arrays, slices, maps, strings, and channels),
// it returns their length.
//
// The second return value reports whether v has a comparable magnitude.
func MagnitudeOf(v any) (any, bool) {
	n, ok := magnitudeFrom(v)
	if !ok {
		return nil, false
	}

	switch n.kind {
	case kindSigned:
		return n.i, true
	case kindUnsigned:
		return n.u, true
	case kindFloating:
		return n.f, true
	}

	return nil, false
}

// CompareMagnitude compares two numeric or measurable values.
//
// Numeric values may be of different built-in numeric types. Measurable values
// (arrays, slices, maps, strings, and channels) are compared using their
// length.
//
// It returns:
//
//	-1 if a < b
//	 0 if a == b
//	 1 if a > b
//
// The second return value reports whether both operands could be compared.
// It is false if either value is neither numeric nor measurable, or if the
// comparison cannot be performed safely without losing numeric precision.
func CompareMagnitude(a, b any) (int, bool) {
	na, ok := magnitudeFrom(a)
	if !ok {
		return 0, false
	}

	nb, ok := magnitudeFrom(b)
	if !ok {
		return 0, false
	}

	return compare(na, nb)
}

func compare(a, b magnitude) (int, bool) {
	switch {
	case a.kind == kindSigned && b.kind == kindSigned:
		return cmp.Compare(a.i, b.i), true

	case a.kind == kindUnsigned && b.kind == kindUnsigned:
		return cmp.Compare(a.u, b.u), true

	case a.kind == kindSigned && b.kind == kindUnsigned:
		if a.i < 0 {
			return -1, true
		}

		return cmp.Compare(uint64(a.i), b.u), true

	case a.kind == kindUnsigned && b.kind == kindSigned:
		if b.i < 0 {
			return 1, true
		}

		return cmp.Compare(a.u, uint64(b.i)), true

	case a.kind == kindFloating:
		return compareFloat(a.f, b)

	case b.kind == kindFloating:
		result, ok := compareFloat(b.f, a)
		if !ok {
			return 0, false
		}

		return -result, true
	}

	return 0, false
}

func compareFloat(f float64, n magnitude) (int, bool) {
	if math.IsNaN(f) {
		return 0, false
	}

	switch n.kind {
	case kindFloating:
		if math.IsNaN(n.f) {
			return 0, false
		}

		return cmp.Compare(f, n.f), true

	case kindSigned:
		if n.i < -maxExactFloatInteger || n.i > maxExactFloatInteger {
			return 0, false
		}

		return cmp.Compare(f, float64(n.i)), true

	case kindUnsigned:
		if n.u > maxExactFloatInteger {
			return 0, false
		}

		return cmp.Compare(f, float64(n.u)), true
	}

	return 0, false
}

func dereferencePointer(v reflect.Value) reflect.Value {
	for {
		if v.Kind() != reflect.Pointer {
			return v
		}

		v = v.Elem()
	}
}
