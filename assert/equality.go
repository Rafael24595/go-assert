package assert

import (
	"github.com/Rafael24595/go-assert/internal/check"
)

// Nil fails the test if the provided item is not nil.
//
// The test continues after a failure.
func Nil(t T, item any, message ...any) {
	t.Helper()
	check.Nil(t, t.Errorf, item, message...)
}

// NotNil fails the test if the provided item is nil or a nil pointer/interface.
//
// The test continues after a failure.
func NotNil(t T, item any, message ...any) {
	t.Helper()
	check.NotNil(t, t.Errorf, item, message...)
}

// True fails the test if the result is false.
//
// The test continues after a failure.
func True(t T, result bool, message ...any) {
	t.Helper()
	check.True(t, t.Errorf, result, message...)
}

// False fails the test if the result is true.
//
// The test continues after a failure.
func False(t T, result bool, message ...any) {
	t.Helper()
	check.False(t, t.Errorf, result, message...)
}

// Equal fails the test if want and have are not equal.
//
// The test continues after a failure.
func Equal[K comparable](t T, want, have K, message ...any) {
	t.Helper()
	check.Equal(t, t.Errorf, want, have, message...)
}

// NotEqual fails the test if want and have are equal.
//
// The test continues after a failure.
func NotEqual[K comparable](t T, want, have K, message ...any) {
	t.Helper()
	check.NotEqual(t, t.Errorf, want, have, message...)
}

// DeepEqual fails the test if want and have are not deeply equal.
// It uses reflect.DeepEqual to compare complex structures, slices, and maps.
//
// The test continues after a failure.
func DeepEqual(t T, want, have any, message ...any) {
	t.Helper()
	check.DeepEqual(t, t.Errorf, want, have, message...)
}

// NotDeepEqual fails the test if want and have are deeply equal.
// It uses reflect.DeepEqual to compare complex structures, slices, and maps.
//
// The test continues after a failure.
func NotDeepEqual(t T, want, have any, message ...any) {
	t.Helper()
	check.NotDeepEqual(t, t.Errorf, want, have, message...)
}

// Same fails the test if want and have do not reference the same object.
//
// The test continues after a failure.
func Same(t T, want, have any, message ...any) {
	t.Helper()
	check.Same(t, t.Errorf, want, have, message...)
}

// NotSame fails the test if want and have reference the same object.
//
// The test continues after a failure.
func NotSame(t T, want, have any, message ...any) {
	t.Helper()
	check.NotSame(t, t.Errorf, want, have, message...)
}

// InDelta fails the test if the absolute difference between want and have
// is greater than the specified delta.
//
// The test continues after a failure.
func InDelta(t T, want, have, delta float64, message ...any) {
	t.Helper()
	check.InDelta(t, t.Errorf, want, have, delta, message...)
}
