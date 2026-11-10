package assert

import (
	"fmt"
	"testing"
)

func TestErrorType(t *testing.T) {
	t.Run("Direct error type match", func(t *testing.T) {
		err := &customError{}
		ErrorType[*customError](t, err)
	})

	t.Run("Wrapped error type match", func(t *testing.T) {
		err := fmt.Errorf("context: %w", &customError{})
		ErrorType[*customError](t, err)
	})
}

func TestErrorNotType(t *testing.T) {
	t.Run("Different error types", func(t *testing.T) {
		err := &customError{}
		ErrorNotType[*anotherCustomError](t, err)
	})

	t.Run("Nil error should not match any type", func(t *testing.T) {
		var err error
		ErrorNotType[*customError](t, err)
	})
}

func TestErrorIs(t *testing.T) {
	t.Run("Direct error value match", func(t *testing.T) {
		have := errSentinel
		ErrorIs(t, errSentinel, have)
	})

	t.Run("Wrapped error value match", func(t *testing.T) {
		have := fmt.Errorf("additional context: %w", errSentinel)
		ErrorIs(t, errSentinel, have)
	})
}

func TestErrorIsNot(t *testing.T) {
	t.Run("Different error values", func(t *testing.T) {
		ErrorIsNot(t, errSentinel, errAnother)
	})

	t.Run("Nil error should not match sentinel", func(t *testing.T) {
		var have error = nil
		ErrorIsNot(t, errSentinel, have)
	})
}

func TestErrorIsReportsValues(t *testing.T) {
	t.Run("Different error", func(t *testing.T) {
		spy := &spyT{}
		ErrorIs(spy, errSentinel, errAnother)
		True(t, spy.HasFailed)
		Equal(t, fmt.Sprintf("Expected error %q, but got %q", errSentinel, errAnother), spy.Message)
	})

	t.Run("Nil error", func(t *testing.T) {
		spy := &spyT{}
		ErrorIs(spy, errSentinel, nil)
		True(t, spy.HasFailed)
		Equal(t, fmt.Sprintf("Expected error %q, but got nil", errSentinel), spy.Message)
	})
}

func TestErrorIsNotReportsValues(t *testing.T) {
	spy := &spyT{}
	have := fmt.Errorf("context: %w", errSentinel)
	ErrorIsNot(spy, errSentinel, have)
	True(t, spy.HasFailed)
	Equal(t, fmt.Sprintf("Expected error not to match %q, but got %q", errSentinel, have), spy.Message)
}
