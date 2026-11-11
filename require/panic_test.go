package require

import (
	"testing"
)

func TestPanicFailures(t *testing.T) {
	expectFailures(t, []failure{
		{"Panic", func(t T) {
			Panic(t, func() {}, "golang %d", 125)
		}},
		{"PanicWithMessage", func(t T) {
			PanicWithMessage(t, "boom", func() {}, "golang %d", 125)
		}},
		{"NotPanic", func(t T) {
			NotPanic(t, func() { panic("boom") }, "golang %d", 125)
		}},
	})
}
