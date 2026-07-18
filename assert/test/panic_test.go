package assert

import "testing"

func TestPanic(t *testing.T) {
	t.Run("Should panic", func(t *testing.T) {
		Panic(t, func() {
			panic("boom")
		})
	})

	t.Run("Panic with message", func(t *testing.T) {
		PanicWithMessage(t, "error crítico", func() {
			panic("error crítico")
		})
	})

	t.Run("Should not panic", func(t *testing.T) {
		NotPanic(t, func() {
			_ = 1 + 1
		})
	})
}
