package assert

import (
	"cmp"
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/Rafael24595/go-assert/assert/internal"
)

// Nil fails the test if the provided item is not nil.
func Nil(t *testing.T, item any, message ...any) {
	t.Helper()

	if isNil(item) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected nil value", custom)
}

// NotNil fails the test if the provided item is nil or a nil pointer/interface.
func NotNil(t *testing.T, item any, message ...any) {
	t.Helper()

	if !isNil(item) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sUnexpected nil value", custom)
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
func True(t *testing.T, result bool, message ...any) {
	t.Helper()

	if result {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected true, but got false", custom)
}

// False fails the test if the result is true.
func False(t *testing.T, result bool, message ...any) {
	t.Helper()

	if !result {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected false, but got true", custom)
}

// Equal fails the test if want and have are not equal.
func Equal[T comparable](t *testing.T, want, have T, message ...any) {
	t.Helper()

	if want == have {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected '%v', but got '%v'", custom, want, have)
}

// NotEqual fails the test if want and have are equal.
func NotEqual[T comparable](t *testing.T, want, have T, message ...any) {
	t.Helper()

	if want != have {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sUnexpected '%v'", custom, want)
}

// DeepEqual fails the test if want and have are not deeply equal.
// It uses reflect.DeepEqual to compare complex structures, slices, and maps.
func DeepEqual(t *testing.T, want, have any, message ...any) {
	t.Helper()

	if reflect.DeepEqual(want, have) {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected deep equality.\nWant: %+v\nGot:  %+v", custom, want, have)
}

// InDelta fails the test if the absolute difference between want and have
// is greater than the specified delta.
func InDelta(t *testing.T, want, have, delta float64, message ...any) {
	t.Helper()

	diff := math.Abs(want - have)
	if diff <= delta {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected %f and %f to be within %f of each other", custom, want, have, delta)
}

// Greater fails the test if have is not greater than want.
func Greater[T cmp.Ordered](t *testing.T, want, have T, message ...any) {
	t.Helper()

	if have > want {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected greater than %v, but got %v", custom, want, have)
}

// GreaterOrEqual fails the test if have is not greater than or equal to want.
func GreaterOrEqual[T cmp.Ordered](t *testing.T, want, have T, message ...any) {
	t.Helper()

	if have >= want {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected greater or equal than %v, but got %v", custom, want, have)
}

// Less fails the test if have is not less than want.
func Less[T cmp.Ordered](t *testing.T, want, have T, message ...any) {
	t.Helper()

	if have < want {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected less than %v, but got %v", custom, want, have)
}

// LessOrEqual fails the test if have is not less than or equal to want.
func LessOrEqual[T cmp.Ordered](t *testing.T, want, have T, message ...any) {
	t.Helper()

	if have <= want {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected less or equal than %v, but got %v", custom, want, have)
}

// Error fails the test if the provided error is nil.
func Error(t *testing.T, err error, message ...any) {
	t.Helper()

	if err != nil {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected error found but nothing found", custom)
}

// NotError fails the test if an error is found (non-nil).
func NotError(t *testing.T, err error, message ...any) {
	t.Helper()

	if err == nil {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sUnexpected error found: '%s'", custom, err.Error())
}

// Len fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
func Len(t *testing.T, want int, have any, message ...any) {
	t.Helper()

	v := reflect.ValueOf(have)
	var got int

	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan, reflect.String:
		got = v.Len()
	default:
		t.Fatalf("Len() assert: type %T is not measurable", have)
	}

	if want != got {
		t.Fatalf("%sExpected %d, but got %d", formatMessage(message...), want, got)
	}
}

// Contains fails the test if the container (string, slice, or array) does not include the item.
func Contains(t *testing.T, container any, item any, message ...any) {
	t.Helper()

	if contains(t, container, item) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection to contain '%v'", custom, item)
}

// NotContains fails the test if the container includes the item.
func NotContains(t *testing.T, container any, item any, message ...any) {
	t.Helper()

	if !contains(t, container, item) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection NOT to contain '%v'", custom, item)
}

func contains(t *testing.T, container any, item any) bool {
	t.Helper()

	val := reflect.ValueOf(container)

	switch val.Kind() {
	case reflect.String:
		strItem, ok := item.(string)
		if !ok {
			return false
		}
		return strings.Contains(val.String(), strItem)

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(val.Index(i).Interface(), item) {
				return true
			}
		}
	case reflect.Map:
		itemVal := reflect.ValueOf(item)
		if !itemVal.IsValid() {
			return false
		}

		if itemVal.Type().AssignableTo(val.Type().Key()) {
			return val.MapIndex(itemVal).IsValid()
		}
	default:
		t.Fatalf("Contains does not support type %T", container)
	}
	return false
}

// Panic fails the test if the provided function does not panic.
func Panic(t *testing.T, fn func(), message ...any) {
	t.Helper()

	custom := formatMessage(message...)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%sexpected panic but function did not panic", custom)
		}
	}()

	fn()
}

// PanicWithMessage fails the test if the function does not panic or if the panic message differs.
func PanicWithMessage(t *testing.T, expected string, fn func(), message ...any) {
	t.Helper()

	custom := formatMessage(message...)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("%sexpected panic but function did not panic", custom)
		} else if expected != "" && fmt.Sprint(r) != expected {
			t.Fatalf("%sexpected panic message %q but got %q", custom, expected, fmt.Sprint(r))
		}
	}()

	fn()
}

// NotPanic fails the test if the provided function panics.
func NotPanic(t *testing.T, fn func(), message ...any) {
	t.Helper()

	custom := formatMessage(message...)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%sexpected no panic but got: %v", custom, r)
		}
	}()

	fn()
}

func formatMessage(message ...any) string {
	if len(message) == 0 {
		return ""
	}

	return internal.FormatMessage(message...) + " - "
}
