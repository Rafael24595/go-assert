package check

import (
	"github.com/Rafael24595/go-assert/internal"
)

// Size fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
func Size[K internal.Number](t T, report Report, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("Size(): %T is not measurable or convertible to %T", have, want)
		return
	}

	if got != 0 {
		value, _ := internal.MagnitudeOf(have)
		report("%sExpected %v, but got %v", prefix(message...), want, value)
	}
}

// Empty fails the test if the length of 'have' is not zero.
func Empty(t T, report Report, have any, message ...any) {
	t.Helper()

	Size(t, report, 0, have, message...)
}

// NotEmpty fails the test if the length of 'have' is zero.
func NotEmpty(t T, report Report, have any, message ...any) {
	t.Helper()

	GreaterThan(t, report, 0, have, message...)
}

// LessThan fails the test if have is not less than want.
func LessThan[K internal.Number](t T, report Report, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("LessThan(): %T is not measurable or convertible to %T", have, want)
		return
	}

	if got == -1 {
		return
	}

	custom := prefix(message...)
	report("%sExpected less than %v, but got %v", custom, want, have)
}

// LessOrEqual fails the test if have is not less than or equal to want.
func LessOrEqual[K internal.Number](t T, report Report, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("LessOrEqual(): %T is not measurable or convertible to %T", have, want)
		return
	}

	if got <= 0 {
		return
	}

	custom := prefix(message...)
	report("%sExpected less or equal than %v, but got %v", custom, want, have)
}

// GreaterThan fails the test if have is not greater than want.
func GreaterThan[K internal.Number](t T, report Report, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("GreaterThan(): %T is not measurable or convertible to %T", have, want)
		return
	}

	if got == 1 {
		return
	}

	custom := prefix(message...)
	report("%sExpected greater than %v, but got %v", custom, want, have)
}

// GreaterOrEqual fails the test if have is not greater than or equal to want.
func GreaterOrEqual[K internal.Number](t T, report Report, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareMagnitude(have, want)
	if !ok {
		t.Fatalf("GreaterOrEqual(): %T is not measurable or convertible to %T", have, want)
		return
	}

	if got >= 0 {
		return
	}

	custom := prefix(message...)
	report("%sExpected greater or equal than %v, but got %v", custom, want, have)
}

// Capacity fails the test if the capacity of 'have' does not match 'want'.
// It supports Slice, Array, and Chan.
func Capacity[K internal.Number](t T, report Report, want K, have any, message ...any) {
	t.Helper()

	got, ok := internal.CompareCapacity(have, want)
	if !ok {
		t.Fatalf("Capacity(): %T is not capacity-measurable or convertible to %T", have, want)
		return
	}

	if got != 0 {
		value, _ := internal.CapacityOf(have)
		report("%sExpected capacity %v, but got %v", prefix(message...), want, value)
	}
}
