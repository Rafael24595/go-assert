package require

import (
	"errors"
	"testing"
)

func TestErrorFailures(t *testing.T) {
	expectFailures(t, []failure{
		{"ErrorType", func(t T) {
			ErrorType[*testError](t, errOther, "golang %d", 125)
		}},
		{"ErrorNotType", func(t T) {
			ErrorNotType[*testError](t, &testError{}, "golang %d", 125)
		}},
		{"ErrorIs", func(t T) {
			ErrorIs(t, errOther, errors.New("x"), "golang %d", 125)
		}},
		{"ErrorIsNot", func(t T) {
			ErrorIsNot(t, errOther, errOther, "golang %d", 125)
		}},
	})
}
