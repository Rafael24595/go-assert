package assert

// T is an interface that represents a testing object, typically *testing.T or *testing.B.
type T interface {
	// Helper marks the calling function as a test helper function.
    Helper()

	// Errorf formats an error message and marks the test as failed.
    Errorf(format string, args ...any)

	// Fatalf formats an error message, marks the test as failed, and stops its execution.
    Fatalf(format string, args ...any)
}
