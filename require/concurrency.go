package require

import (
	"time"

	"github.com/Rafael24595/go-assert/internal/check"
)

// WillClose fails the test if the provided channel does not close or
// receive a signal within the specified timeout duration.
//
// The test stops at the first failure.
func WillClose(t T, ch <-chan struct{}, timeout time.Duration, message ...any) {
	t.Helper()
	check.WillClose(t, t.Fatalf, ch, timeout, message...)
}
