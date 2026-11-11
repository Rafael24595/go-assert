package assert

import (
	"errors"
	"strings"
	"testing"

	"github.com/Rafael24595/go-assert/internal/spy"
)

type testError struct{}

var errOther = errors.New("other error")

func (e *testError) Error() string { return "test error" }

type failure struct {
	name string
	call func(t T)
}

func expectFailures(t *testing.T, failures []failure) {
	t.Helper()

	for _, f := range failures {
		t.Run(f.name, func(t *testing.T) {
			spy := &spy.T{}

			f.call(spy)
			if !spy.HasFailed || spy.HasFatal {
				t.Errorf("errorf=%v fatalf=%v, want the failure through t.Errorf: the test continues", spy.HasFailed, spy.HasFatal)
			}

			if !strings.HasPrefix(spy.Message, "golang 125 - ") {
				t.Errorf("the caller's message was not forwarded: %q", spy.Message)
			}

			if !spy.IsHelper {
				t.Error("the assertion must call t.Helper()")
			}
		})
	}
}

func expectStops(t *testing.T, failures []failure) {
	t.Helper()

	for _, f := range failures {
		t.Run(f.name, func(t *testing.T) {
			spy := &spy.T{}

			f.call(spy)
			if !spy.HasFatal || spy.HasFailed {
				t.Errorf("errorf=%v fatalf=%v, want t.Fatalf only", spy.HasFailed, spy.HasFatal)
			}
		})
	}
}
