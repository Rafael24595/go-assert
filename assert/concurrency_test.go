package assert

import (
	"testing"
	"time"
)

func TestConcurrencyFailures(t *testing.T) {
	expectFailures(t, []failure{
		{"WillClose", func(t T) {
			WillClose(t, make(chan struct{}), time.Millisecond, "golang %d", 125)
		}},
	})
}
