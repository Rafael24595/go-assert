package check

import (
	"math"
	"reflect"
)

// Nil fails the test if the provided item is not nil.
func Nil(t T, report Report, item any, message ...any) {
	t.Helper()

	if isNil(item) {
		return
	}

	custom := prefix(message...)
	report("%sExpected nil, but got %v (%T)", custom, item, item)
}

// NotNil fails the test if the provided item is nil or a nil pointer/interface.
func NotNil(t T, report Report, item any, message ...any) {
	t.Helper()

	if !isNil(item) {
		return
	}

	custom := prefix(message...)
	report("%sUnexpected nil value", custom)
}

func isNil(item any) bool {
	if item == nil {
		return true
	}

	v := reflect.ValueOf(item)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return false
}

// True fails the test if the result is false.
func True(t T, report Report, result bool, message ...any) {
	t.Helper()

	if result {
		return
	}

	custom := prefix(message...)
	report("%sExpected true, but got false", custom)
}

// False fails the test if the result is true.
func False(t T, report Report, result bool, message ...any) {
	t.Helper()

	if !result {
		return
	}

	custom := prefix(message...)
	report("%sExpected false, but got true", custom)
}

// Equal fails the test if want and have are not equal.
func Equal[K comparable](t T, report Report, want, have K, message ...any) {
	t.Helper()

	if want == have {
		return
	}

	custom := prefix(message...)
	report("%sExpected '%v', but got '%v'", custom, want, have)
}

// NotEqual fails the test if want and have are equal.
func NotEqual[K comparable](t T, report Report, want, have K, message ...any) {
	t.Helper()

	if want != have {
		return
	}

	custom := prefix(message...)
	report("%sUnexpected '%v'", custom, want)
}

// DeepEqual fails the test if want and have are not deeply equal.
// It uses reflect.DeepEqual to compare complex structures, slices, and maps.
func DeepEqual(t T, report Report, want, have any, message ...any) {
	t.Helper()

	if reflect.DeepEqual(want, have) {
		return
	}

	custom := prefix(message...)
	report("%sExpected deep equality.\nWant: %+v\nGot:  %+v", custom, want, have)
}

// NotDeepEqual fails the test if want and have are deeply equal.
// It uses reflect.DeepEqual to compare complex structures, slices, and maps.
func NotDeepEqual(t T, report Report, want, have any, message ...any) {
	t.Helper()

	if !reflect.DeepEqual(want, have) {
		return
	}

	custom := prefix(message...)
	report("%sExpected values to be deeply different.\nWant: %+v\nGot:  %+v", custom, want, have)
}

// Same fails the test if want and have do not reference the same object.
func Same(t T, report Report, want, have any, message ...any) {
	t.Helper()

	v1 := reflect.ValueOf(want)
	v2 := reflect.ValueOf(have)

	if !isReference(v1.Kind()) || !isReference(v2.Kind()) {
		t.Fatalf("Same only supports reference types")
		return
	}

	if v1.Pointer() == v2.Pointer() {
		return
	}

	custom := prefix(message...)
	report("%sExpected both values to reference the same object", custom)
}

// NotSame fails the test if want and have reference the same object.
func NotSame(t T, report Report, want, have any, message ...any) {
	t.Helper()

	v1 := reflect.ValueOf(want)
	v2 := reflect.ValueOf(have)

	if !isReference(v1.Kind()) || !isReference(v2.Kind()) {
		t.Fatalf("NotSame only supports reference types")
		return
	}

	if v1.Pointer() != v2.Pointer() {
		return
	}

	custom := prefix(message...)
	report("%sExpected different references", custom)
}

// InDelta fails the test if the absolute difference between want and have
// is greater than the specified delta.
func InDelta(t T, report Report, want, have, delta float64, message ...any) {
	t.Helper()

	diff := math.Abs(want - have)
	if diff <= delta {
		return
	}

	custom := prefix(message...)
	report("%sExpected %f and %f to be within %f of each other", custom, want, have, delta)
}

func isReference(k reflect.Kind) bool {
	switch k {
	case reflect.Pointer,
		reflect.Map,
		reflect.Slice,
		reflect.Func,
		reflect.Chan,
		reflect.UnsafePointer:
		return true
	default:
		return false
	}
}
