package assert

import (
	"github.com/Rafael24595/go-assert/internal/check"
)

// Panic fails the test if the provided function does not panic.
//
// The test continues after a failure.
func Panic(t T, fn func(), message ...any) {
	t.Helper()
	check.Panic(t, t.Errorf, fn, message...)
}

// PanicWithMessage fails the test if the function does not panic or if the panic message differs.
//
// The test continues after a failure.
func PanicWithMessage(t T, expected string, fn func(), message ...any) {
	t.Helper()
	check.PanicWithMessage(t, t.Errorf, expected, fn, message...)
}

// NotPanic fails the test if the provided function panics.
//
// The test continues after a failure.
func NotPanic(t T, fn func(), message ...any) {
	t.Helper()
	check.NotPanic(t, t.Errorf, fn, message...)
}
