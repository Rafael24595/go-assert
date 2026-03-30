package internal

import (
	"testing"
)

func TestFormatMessage(t *testing.T) {
	t.Run("Empty message", func(t *testing.T) {
		if got := FormatMessage(); got != "" {
			t.Errorf("Expected empty string, got %q", got)
		}
	})

	t.Run("Single string", func(t *testing.T) {
		msg := "ziglang"
		if got := FormatMessage(msg); got != msg {
			t.Errorf("Expected %q, got %q", msg, got)
		}
	})

	t.Run("Format string with args", func(t *testing.T) {
		got := FormatMessage("hello %s", "golang")
		want := "hello golang"
		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})

	t.Run("Multiple values without format string", func(t *testing.T) {
		got := FormatMessage("valor1", "valor2", 100)
		want := "valor1%!(EXTRA string=valor2, int=100)"
		if got != want {
			t.Errorf("Unexpected format: %q", got)
		}
	})
}
