//go:build !g_ast_dbg && !g_ast_wrt

package assert

// Unreachable is a no-op in this build configuration.
func Unreachable(a ...any) {}

// True is a no-op in this build configuration.
func True(cond bool, a ...any) {}

// False is a no-op in this build configuration.
func False(cond bool, a ...any) {}

// LazyTrue is a no-op in this build configuration.
func LazyTrue(p predicate, a ...any) {}

// LazyFalse is a no-op in this build configuration.
func LazyFalse(p predicate, a ...any) {}
