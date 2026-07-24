package assert

import (
	"github.com/Rafael24595/go-assert/assert/internal"
)

// Size fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
func Size[K internal.Number](t T, want K, have any, message ...any) {
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
func Empty(t T, have any, message ...any) {
	t.Helper()

	Size(t, 0, have, message...)
}

// NotEmpty fails the test if the length of 'have' is zero.
func NotEmpty(t T, have any, message ...any) {
	t.Helper()

	GreaterThan(t, 0, have, message...)
}

// Deprecated: Use Size instead.
//
// Len fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
func Len[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()

	Size(t, want, have, message...)
}

// LessThan fails the test if have is not less than want.
func LessThan[K internal.Number](t T, want K, have any, message ...any) {
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
func Less[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()

	LessThan(t, want, have, message...)
}

// LessOrEqual fails the test if have is not less than or equal to want.
func LessOrEqual[K internal.Number](t T, want K, have any, message ...any) {
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
func GreaterThan[K internal.Number](t T, want K, have any, message ...any) {
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
func Greater[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()

	GreaterThan(t, want, have, message...)
}

// GreaterOrEqual fails the test if have is not greater than or equal to want.
func GreaterOrEqual[K internal.Number](t T, want K, have any, message ...any) {
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

// Capacity fails the test if the capacity of 'have' does not match 'want'.
// It supports Slice, Array, and Chan.
func Capacity[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareCapacity(have, want)
	if !ok {
		t.Fatalf("Capacity(): %T is not capacity-measurable or convertible to %T", have, want)
	}

	if got != 0 {
		value, _ := internal.CapacityOf(have)
		t.Fatalf("%sExpected capacity %v, but got %v", formatMessage(message...), want, value)
	}
}
