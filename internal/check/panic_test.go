package check

import (
	"testing"
)

func TestPanic(t *testing.T) {
	boom := func() { panic("boom") }
	calm := func() { _ = 1 + 1 }

	t.Run("Panic", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			Panic(t, r, boom)
		})

		reports(t, "expected panic but function did not panic", func(t T, r Report) {
			Panic(t, r, calm)
		})
	})

	t.Run("PanicWithMessage", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			PanicWithMessage(t, r, "boom", boom)
		})

		passes(t, func(t T, r Report) {
			PanicWithMessage(t, r, "", boom)
		})

		reports(t, `expected panic message "bang" but got "boom"`,
			func(t T, r Report) {
				PanicWithMessage(t, r, "bang", boom)
			})

		reports(t, "expected panic but function did not panic",
			func(t T, r Report) {
				PanicWithMessage(t, r, "boom", calm)
			})
	})

	t.Run("NotPanic", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			NotPanic(t, r, calm)
		})

		reports(t, "expected no panic but got: boom", func(t T, r Report) {
			NotPanic(t, r, boom)
		})
	})
}
