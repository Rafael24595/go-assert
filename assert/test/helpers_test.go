package assert

import 	"errors"

var (
	errSentinel = errors.New("sentinel error")
	errAnother  = errors.New("another error")
)

type customInt int

type customError struct{}

func (e *customError) Error() string {
	return ""
}

type anotherCustomError struct{}

func (e *anotherCustomError) Error() string {
	return ""
}
