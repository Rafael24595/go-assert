package assert

// Unreachable fails the test immediately indicating that a code path
// that should never be executed was reached.
func Unreachable(t T, message ...any) {
	t.Helper()

	custom := formatMessage(message...)
	t.Fatalf("%scode should be unreachable", custom)
}
