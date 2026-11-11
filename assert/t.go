package assert

import "github.com/Rafael24595/go-assert/internal/check"

// T is the subset of testing.T the assertions need; *testing.T and *testing.B satisfy it.
type T = check.T
