package check

// Unreachable fails the test immediately indicating that a code path
// that should never be executed was reached. It always stops the test,
// so it takes no Report.
func Unreachable(t T, message ...any) {
	t.Helper()

	custom := prefix(message...)
	t.Fatalf("%scode should be unreachable", custom)
}
