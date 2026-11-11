package assert

import (
	"github.com/Rafael24595/go-assert/internal/check"
)

// Unreachable fails the test immediately indicating that a code path
// that should never be executed was reached. It always stops the test.
func Unreachable(t T, message ...any) {
	t.Helper()
	check.Unreachable(t, message...)
}
