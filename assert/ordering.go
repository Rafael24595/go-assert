package assert

import (
	"github.com/Rafael24595/go-assert/internal"
	"github.com/Rafael24595/go-assert/internal/check"
)

// Size fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
//
// The test continues after a failure.
func Size[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.Size(t, t.Errorf, want, have, message...)
}

// Empty fails the test if the length of 'have' is not zero.
//
// The test continues after a failure.
func Empty(t T, have any, message ...any) {
	t.Helper()
	check.Empty(t, t.Errorf, have, message...)
}

// NotEmpty fails the test if the length of 'have' is zero.
//
// The test continues after a failure.
func NotEmpty(t T, have any, message ...any) {
	t.Helper()
	check.NotEmpty(t, t.Errorf, have, message...)
}

// LessThan fails the test if have is not less than want.
//
// The test continues after a failure.
func LessThan[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.LessThan(t, t.Errorf, want, have, message...)
}

// LessOrEqual fails the test if have is not less than or equal to want.
//
// The test continues after a failure.
func LessOrEqual[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.LessOrEqual(t, t.Errorf, want, have, message...)
}

// GreaterThan fails the test if have is not greater than want.
//
// The test continues after a failure.
func GreaterThan[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.GreaterThan(t, t.Errorf, want, have, message...)
}

// GreaterOrEqual fails the test if have is not greater than or equal to want.
//
// The test continues after a failure.
func GreaterOrEqual[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.GreaterOrEqual(t, t.Errorf, want, have, message...)
}

// Capacity fails the test if the capacity of 'have' does not match 'want'.
// It supports Slice, Array, and Chan.
//
// The test continues after a failure.
func Capacity[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.Capacity(t, t.Errorf, want, have, message...)
}
