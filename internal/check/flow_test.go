package check

import (
	"testing"
)

func TestUnreachable(t *testing.T) {
	t.Run("always stops the test", func(t *testing.T) {
		stops(t, "code should be unreachable", func(t T, _ Report) {
			Unreachable(t)
		})
	})

	t.Run("prefixes the caller's message", func(t *testing.T) {
		stops(t, "unexpected state - code should be unreachable",
			func(t T, _ Report) {
				Unreachable(t, "unexpected state")
			})
	})
}
