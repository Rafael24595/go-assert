package assert

import (
	"github.com/Rafael24595/go-assert/internal/check"
)

// Inside fails the test if the container (string, slice, or array) does not include the item.
//
// The test continues after a failure.
func Inside(t T, item, container any, message ...any) {
	t.Helper()
	check.Inside(t, t.Errorf, item, container, message...)
}

// NotInside fails the test if the container includes the item.
//
// The test continues after a failure.
func NotInside(t T, item, container any, message ...any) {
	t.Helper()
	check.NotInside(t, t.Errorf, item, container, message...)
}
