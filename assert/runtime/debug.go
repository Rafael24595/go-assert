//go:build g_ast_dbg && !g_ast_wrt

package assert

import "github.com/Rafael24595/go-assert/assert/internal"

// Unreachable panics with a formatted message, indicating that 
// a section of code should not have been reached.
func Unreachable(a ...any) {
	msg := internal.FormatMessage(a)
	panic(msg)
}

// True panics if the provided condition is false.
func True(cond bool, a ...any) {
	if cond {
		return
	}

	msg := internal.FormatMessage(a)
	panic(msg)
}

// False panics if the provided condition is true.
func False(cond bool, a ...any) {
	if !cond {
		return
	}

	msg := internal.FormatMessage(a)
	panic(msg)
}

// LazyTrue evaluates the predicate and panics if the result is false.
func LazyTrue(p predicate, a ...any) {
	if p() {
		return
	}

	msg := internal.FormatMessage(a)
	panic(msg)
}

// LazyFalse evaluates the predicate and panics if the result is true.
func LazyFalse(p predicate, a ...any) {
	if !p() {
		return
	}

	msg := internal.FormatMessage(a)
	panic(msg)
}
