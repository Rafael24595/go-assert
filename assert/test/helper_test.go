package assert

import (
	"errors"
	"fmt"
)

var (
	errSentinel = errors.New("sentinel error")
	errAnother  = errors.New("another error")
)

type spyT struct {
	HasFailed bool
	HasFatal  bool
	Message   string
}

func (s *spyT) Helper() {}

func (s *spyT) Errorf(format string, args ...any) {
	s.HasFailed = true
	s.Message = fmt.Sprintf(format, args...)
}

func (s *spyT) Fatalf(format string, args ...any) {
	s.HasFatal = true
	s.Message = fmt.Sprintf(format, args...)
}

type customInt int

type customError struct{}

func (e *customError) Error() string {
	return ""
}

type anotherCustomError struct{}

func (e *anotherCustomError) Error() string {
	return ""
}
