package assert

import (
	"testing"
	"time"
)

// WillClose fails the test if the provided channel does not close or
// receive a signal within the specified timeout duration.
func WillClose(t *testing.T, ch <-chan struct{}, timeout time.Duration, message ...any) {
	t.Helper()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-ch:
		return
	case <-timer.C:
		custom := formatMessage(message...)
		t.Fatalf("%stimeout: channel did not close within %v", custom, timeout)
	}
}
