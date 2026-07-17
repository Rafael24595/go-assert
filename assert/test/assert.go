package assert

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

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

// NotDeepEqual fails the test if want and have are deeply equal.
// It uses reflect.DeepEqual to compare complex structures, slices, and maps.
func NotDeepEqual(t *testing.T, want, have any, message ...any) {
	t.Helper()
	
	if !reflect.DeepEqual(want, have) {
		return
	}
	
	custom := formatMessage(message...)
	
	t.Errorf("%sExpected values to be deeply different.\nWant: %+v\nGot:  %+v", custom, want, have)
}

// Same fails the test if want and have do not reference the same object.
func Same(t *testing.T, want, have any, message ...any) {
	t.Helper()

	v1 := reflect.ValueOf(want)
	v2 := reflect.ValueOf(have)

	if !isReference(v1.Kind()) || !isReference(v2.Kind()) {
		t.Fatalf("Same only supports reference types")
	}

	if v1.Pointer() == v2.Pointer() {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected both values to reference the same object", custom)
}

// NotSame fails the test if want and have reference the same object.
func NotSame(t *testing.T, want, have any, message ...any) {
	t.Helper()

	v1 := reflect.ValueOf(want)
	v2 := reflect.ValueOf(have)

	if !isReference(v1.Kind()) || !isReference(v2.Kind()) {
		t.Fatalf("NotSame only supports reference types")
	}

	if v1.Pointer() != v2.Pointer() {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected different references", custom)
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

// Size fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
func Size[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("Size(): %T is not measurable or convertible to %T", have, want)
	}

	if got != 0 {
		value, _ := internal.MagnitudeOf(have)
		t.Fatalf("%sExpected %v, but got %v", formatMessage(message...), want, value)
	}
}

// Empty fails the test if the length of 'have' is not zero.
func Empty(t *testing.T, have any, message ...any) {
	t.Helper()

	Size(t, 0, have, message...)
}

// NotEmpty fails the test if the length of 'have' is zero.
func NotEmpty(t *testing.T, have any, message ...any) {
	t.Helper()

	GreaterThan(t, 0, have, message...)
}

// Deprecated: Use Size instead.
//
// Len fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
func Len[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	Size(t, want, have, message...)
}

// LessThan fails the test if have is not less than want.
func LessThan[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("LessThan(): %T is not measurable or convertible to %T", have, want)
	}

	if got == -1 {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected less than %v, but got %v", custom, want, have)
}

// Deprecated: Use LessThan instead.
//
// Less fails the test if have is not less than want.
func Less[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	LessThan(t, want, have, message...)
}

// LessOrEqual fails the test if have is not less than or equal to want.
func LessOrEqual[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("LessOrEqual(): %T is not measurable or convertible to %T", have, want)
	}

	if got <= 0 {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected less or equal than %v, but got %v", custom, want, have)
}

// GreaterThan fails the test if have is not greater than want.
func GreaterThan[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("GreaterThan(): %T is not measurable or convertible to %T", have, want)
	}

	if got == 1 {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected greater than %v, but got %v", custom, want, have)
}

// Deprecated: Use GreaterThan instead.
//
// Greater fails the test if have is not greater than want.
func Greater[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	GreaterThan(t, want, have, message...)
}

// GreaterOrEqual fails the test if have is not greater than or equal to want.
func GreaterOrEqual[T internal.Number](t *testing.T, want T, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("GreaterOrEqual(): %T is not measurable or convertible to %T", have, want)
	}

	if got >= 0 {
		return
	}

	custom := formatMessage(message...)

	t.Errorf("%sExpected greater or equal than %v, but got %v", custom, want, have)
}

// Inside fails the test if the container (string, slice, or array) does not include the item.
func Inside(t *testing.T, item, container any, message ...any) {
	t.Helper()

	if inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection to contain '%v'", custom, item)
}

// NotInside fails the test if the container includes the item.
func NotInside(t *testing.T, item, container any, message ...any) {
	t.Helper()

	if !inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection NOT to contain '%v'", custom, item)
}

func inside(t *testing.T, item, container any) bool {
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
		t.Fatalf("Inside does not support type %T", container)
	}
	return false
}

// Deprecated: Use Inside instead.
//
// Contains fails the test if the container (string, slice, or array) does not include the item.
func Contains(t *testing.T, container any, item any, message ...any) {
	t.Helper()

	if inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection to contain '%v'", custom, item)
}

// Deprecated: Use NotInside instead.
//
// NotContains fails the test if the container includes the item.
func NotContains(t *testing.T, container any, item any, message ...any) {
	t.Helper()

	if !inside(t, item, container) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected collection NOT to contain '%v'", custom, item)
}

// ErrorType fails the test if the provided error is not of the expected type.
func ErrorType[T error](t *testing.T, have error, message ...any) {
	t.Helper()

	var target T
	// Aquí &target es un puntero real al tipo T (ej: **customError), lo cual preserva el tipo para errors.As
	if errors.As(have, &target) {
		return
	}

	custom := formatMessage(message...)
	if have == nil {
		t.Errorf("%sExpected error of type %T, but got nil", custom, target)
		return
	}

	t.Errorf("%sExpected error of type %T, but got %T (%v)", custom, target, have, have)
}

// ErrorNotType fails the test if the provided error is of the specified type.
func ErrorNotType[T error](t *testing.T, have error, message ...any) {
	t.Helper()

	var target T
	if !errors.As(have, &target) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sUnexpected error of type %T received", custom, target)
}

// ErrorIs fails the test if the provided error does not match the expected type, even when wrapped.
func ErrorIs(t *testing.T, want, have error, message ...any) {
	t.Helper()

	custom := formatMessage(message...)

	if have == nil {
		t.Errorf("%sExpected error of type %T, but got nil", custom, want)
		return
	}

	if errors.Is(have, want) {
		return
	}

	t.Errorf("%sExpected error of type %T, but got %T (%v)", custom, want, have, have)
}

// ErrorIsNot fails the test if the provided error matches the specified type, even when wrapped.
func ErrorIsNot(t *testing.T, want, have error, message ...any) {
	t.Helper()

	if !errors.Is(have, want) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sUnexpected error of type %T received", custom, want)
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

// WillClose fails the test if the provided channel does not close or
// receive a signal within the specified timeout duration.
func WillClose(t *testing.T, ch <-chan struct{}, timeout time.Duration, message ...any) {
	t.Helper()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-ch:
		return
	case <-timer.C:
		custom := formatMessage(message...)
		t.Fatalf("%stimeout: channel did not close within %v", custom, timeout)
	}
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

func formatMessage(message ...any) string {
	if len(message) == 0 {
		return ""
	}

	return internal.FormatMessage(message...) + " - "
}
