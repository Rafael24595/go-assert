package assert

import (
	"testing"
)

func TestUnreachable(t *testing.T) {
	t.Run("fails the test and invokes Helper", func(t *testing.T) {
		spy := &spyT{}

		Unreachable(spy)

		True(t, spy.HasFatal)
		True(t, spy.IsHelper)
	})

	t.Run("formats custom message correctly", func(t *testing.T) {
		spy := &spyT{}

		message := "unexpected state"

		Unreachable(spy, message)

		Inside(t, message, spy.Message)
	})
}
