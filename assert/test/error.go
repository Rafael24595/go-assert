package assert

import (
	"errors"
)

// ErrorType fails the test if the provided error is not of the expected type.
func ErrorType[K error](t T, have error, message ...any) {
	t.Helper()

	var target K
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
func ErrorNotType[K error](t T, have error, message ...any) {
	t.Helper()

	var target K
	if !errors.As(have, &target) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sUnexpected error of type %T received", custom, target)
}

// ErrorIs fails the test if the provided error does not match the expected type, even when wrapped.
func ErrorIs(t T, want, have error, message ...any) {
	t.Helper()

	if errors.Is(have, want) {
		return
	}

	custom := formatMessage(message...)

	if have == nil {
		t.Errorf("%sExpected error %q, but got nil", custom, want)
		return
	}

	t.Errorf("%sExpected error %q, but got %q", custom, want, have)
}

// ErrorIsNot fails the test if the provided error matches the specified type, even when wrapped.
func ErrorIsNot(t T, want, have error, message ...any) {
	t.Helper()

	if !errors.Is(have, want) {
		return
	}

	custom := formatMessage(message...)
	t.Errorf("%sExpected error not to match %q, but got %q", custom, want, have)
}
