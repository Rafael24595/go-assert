package require

import (
	"github.com/Rafael24595/go-assert/internal/check"
)

// ErrorType fails the test if the provided error is not of the expected type.
//
// The test stops at the first failure.
func ErrorType[K error](t T, have error, message ...any) {
	t.Helper()
	check.ErrorType[K](t, t.Fatalf, have, message...)
}

// ErrorNotType fails the test if the provided error is of the specified type.
//
// The test stops at the first failure.
func ErrorNotType[K error](t T, have error, message ...any) {
	t.Helper()
	check.ErrorNotType[K](t, t.Fatalf, have, message...)
}

// ErrorIs fails the test if the provided error does not match the expected type, even when wrapped.
//
// The test stops at the first failure.
func ErrorIs(t T, want, have error, message ...any) {
	t.Helper()
	check.ErrorIs(t, t.Fatalf, want, have, message...)
}

// ErrorIsNot fails the test if the provided error matches the specified type, even when wrapped.
//
// The test stops at the first failure.
func ErrorIsNot(t T, want, have error, message ...any) {
	t.Helper()
	check.ErrorIsNot(t, t.Fatalf, want, have, message...)
}
