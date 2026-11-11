package require

import (
	"github.com/Rafael24595/go-assert/internal"
	"github.com/Rafael24595/go-assert/internal/check"
)

// Size fails the test if the length of 'have' does not match 'want'.
// It supports Slice, Map, Array, Chan, and String.
//
// The test stops at the first failure.
func Size[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.Size(t, t.Fatalf, want, have, message...)
}

// Empty fails the test if the length of 'have' is not zero.
//
// The test stops at the first failure.
func Empty(t T, have any, message ...any) {
	t.Helper()
	check.Empty(t, t.Fatalf, have, message...)
}

// NotEmpty fails the test if the length of 'have' is zero.
//
// The test stops at the first failure.
func NotEmpty(t T, have any, message ...any) {
	t.Helper()
	check.NotEmpty(t, t.Fatalf, have, message...)
}

// LessThan fails the test if have is not less than want.
//
// The test stops at the first failure.
func LessThan[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.LessThan(t, t.Fatalf, want, have, message...)
}

// LessOrEqual fails the test if have is not less than or equal to want.
//
// The test stops at the first failure.
func LessOrEqual[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.LessOrEqual(t, t.Fatalf, want, have, message...)
}

// GreaterThan fails the test if have is not greater than want.
//
// The test stops at the first failure.
func GreaterThan[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.GreaterThan(t, t.Fatalf, want, have, message...)
}

// GreaterOrEqual fails the test if have is not greater than or equal to want.
//
// The test stops at the first failure.
func GreaterOrEqual[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.GreaterOrEqual(t, t.Fatalf, want, have, message...)
}

// Capacity fails the test if the capacity of 'have' does not match 'want'.
// It supports Slice, Array, and Chan.
//
// The test stops at the first failure.
func Capacity[K internal.Number](t T, want K, have any, message ...any) {
	t.Helper()
	check.Capacity(t, t.Fatalf, want, have, message...)
}
