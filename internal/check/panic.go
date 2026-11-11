package check

import (
	"fmt"
)

// Panic fails the test if the provided function does not panic.
func Panic(t T, report Report, fn func(), message ...any) {
	t.Helper()

	custom := prefix(message...)

	defer func() {
		t.Helper()

		if r := recover(); r == nil {
			report("%sexpected panic but function did not panic", custom)
		}
	}()

	fn()
}

// PanicWithMessage fails the test if the function does not panic or if the panic message differs.
func PanicWithMessage(t T, report Report, expected string, fn func(), message ...any) {
	t.Helper()

	custom := prefix(message...)

	defer func() {
		t.Helper()

		r := recover()
		if r == nil {
			report("%sexpected panic but function did not panic", custom)
			return
		}

		if expected != "" && fmt.Sprint(r) != expected {
			report("%sexpected panic message %q but got %q", custom, expected, fmt.Sprint(r))
		}
	}()

	fn()
}

// NotPanic fails the test if the provided function panics.
func NotPanic(t T, report Report, fn func(), message ...any) {
	t.Helper()

	custom := prefix(message...)

	defer func() {
		t.Helper()

		if r := recover(); r != nil {
			report("%sexpected no panic but got: %v", custom, r)
		}
	}()

	fn()
}
