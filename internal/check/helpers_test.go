package check

import (
	"errors"
	"testing"

	"github.com/Rafael24595/go-assert/internal/spy"
)

type customInt int

type customError struct{}

func (e *customError) Error() string { return "custom" }

type anotherCustomError struct{}

func (e *anotherCustomError) Error() string { return "another" }

var (
	errSentinel = errors.New("sentinel error")
	errAnother  = errors.New("another error")
)

type call func(t T, report Report)

func run(c call) (*spy.T, *spy.Report) {
	st := &spy.T{}
	rep := &spy.Report{}

	c(st, rep.Func())
	return st, rep
}

func passes(t *testing.T, c call) {
	t.Helper()

	st, rep := run(c)
	if rep.Called() || st.HasFailed || st.HasFatal {
		t.Errorf("expected the check to pass; report=%v errorf=%v fatalf=%v message=%q%q",
			rep.Called(), st.HasFailed, st.HasFatal, rep.Message, st.Message)
	}

	if !st.IsHelper {
		t.Error("the check must call t.Helper()")
	}
}

func reports(t *testing.T, want string, c call) {
	t.Helper()

	st, rep := run(c)
	if rep.Calls != 1 {
		t.Errorf("expected one report, got %d", rep.Calls)
	}

	if rep.Message != want {
		t.Errorf("report message:\n want %q\n have %q", want, rep.Message)
	}

	if st.HasFailed || st.HasFatal {
		t.Errorf("a mismatch must go through report, not t (errorf=%v fatalf=%v)", st.HasFailed, st.HasFatal)
	}
}

func stops(t *testing.T, want string, c call) {
	t.Helper()

	st, rep := run(c)
	if !st.HasFatal {
		t.Error("expected t.Fatalf")
	}

	if st.Message != want {
		t.Errorf("fatal message:\n want %q\n have %q", want, st.Message)
	}

	if rep.Called() || st.HasFailed {
		t.Errorf("a fatal check must not also report (report=%v errorf=%v)", rep.Called(), st.HasFailed)
	}
}
