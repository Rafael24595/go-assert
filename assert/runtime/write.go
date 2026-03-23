//go:build g_ast_wrt && !g_ast_dbg

package assert

import "github.com/Rafael24595/go-assert/assert/internal"

// Unreachable writes a formatted message to the default writer 
// indicating that a section of code should not have been executed.
func Unreachable(a ...any) {
	msg := internal.FormatMessage(a)
	writer.Write([]byte(msg))
}

// True writes a formatted message to the default writer if the 
// provided condition is false.
func True(cond bool, a ...any) {
	if cond {
		return
	}

	msg := internal.FormatMessage(a)
	writer.Write([]byte(msg))
}

// False writes a formatted message to the default writer if the 
// provided condition is true.
func False(cond bool, a ...any) {
	if !cond {
		return
	}

	msg := internal.FormatMessage(a)
	writer.Write([]byte(msg))
}

// LazyTrue evaluates the predicate and writes a formatted message 
// if the result is false.
func LazyTrue(p predicate, a ...any) {
	if p() {
		return
	}

	msg := internal.FormatMessage(a)
	writer.Write([]byte(msg))
}

// LazyFalse evaluates the predicate and writes a formatted message 
// if the result is true.
func LazyFalse(p predicate, a ...any) {
	if !p() {
		return
	}

	msg := internal.FormatMessage(a)
	writer.Write([]byte(msg))
}
