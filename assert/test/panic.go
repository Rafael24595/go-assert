package assert

import (
	"fmt"
	"testing"
)

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
