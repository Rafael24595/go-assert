package check

import (
	"fmt"
	"testing"
)

func TestErrorType(t *testing.T) {
	t.Run("direct match", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorType[*customError](t, r, &customError{})
		})
	})

	t.Run("wrapped match", func(t *testing.T) {
		err := fmt.Errorf("context: %w", &customError{})
		passes(t, func(t T, r Report) {
			ErrorType[*customError](t, r, err)
		})
	})

	t.Run("reports another type", func(t *testing.T) {
		reports(t, "Expected error of type *check.customError, but got *check.anotherCustomError (another)",
			func(t T, r Report) {
				ErrorType[*customError](t, r, &anotherCustomError{})
			})
	})

	t.Run("reports nil", func(t *testing.T) {
		reports(t, "Expected error of type *check.customError, but got nil",
			func(t T, r Report) {
				ErrorType[*customError](t, r, nil)
			})
	})
}

func TestErrorNotType(t *testing.T) {
	t.Run("different types", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorNotType[*anotherCustomError](t, r, &customError{})
		})
	})

	t.Run("nil matches no type", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorNotType[*customError](t, r, nil)
		})
	})

	t.Run("reports the matching type", func(t *testing.T) {
		err := fmt.Errorf("context: %w", &customError{})
		reports(t, "Unexpected error of type *check.customError received",
			func(t T, r Report) {
				ErrorNotType[*customError](t, r, err)
			})
	})
}

func TestErrorIs(t *testing.T) {
	t.Run("direct match", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorIs(t, r, errSentinel, errSentinel)
		})
	})

	t.Run("wrapped match", func(t *testing.T) {
		have := fmt.Errorf("additional context: %w", errSentinel)
		passes(t, func(t T, r Report) {
			ErrorIs(t, r, errSentinel, have)
		})
	})

	t.Run("nil matches nil", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorIs(t, r, nil, nil)
		})
	})

	t.Run("reports both errors", func(t *testing.T) {
		reports(t, `Expected error "sentinel error", but got "another error"`,
			func(t T, r Report) {
				ErrorIs(t, r, errSentinel, errAnother)
			})
	})

	t.Run("reports nil", func(t *testing.T) {
		reports(t, `Expected error "sentinel error", but got nil`,
			func(t T, r Report) {
				ErrorIs(t, r, errSentinel, nil)
			})
	})
}

func TestErrorIsNot(t *testing.T) {
	t.Run("different errors", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorIsNot(t, r, errSentinel, errAnother)
		})
	})

	t.Run("nil does not match a sentinel", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			ErrorIsNot(t, r, errSentinel, nil)
		})
	})

	t.Run("reports a wrapped match", func(t *testing.T) {
		have := fmt.Errorf("context: %w", errSentinel)
		reports(t, `Expected error not to match "sentinel error", but got "context: sentinel error"`,
			func(t T, r Report) {
				ErrorIsNot(t, r, errSentinel, have)
			})
	})
}
