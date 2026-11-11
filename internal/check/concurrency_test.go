package check

import (
	"testing"
	"time"
)

func TestWillClose(t *testing.T) {
	t.Run("channel closes in time", func(t *testing.T) {
		ch := make(chan struct{})

		go func() {
			time.Sleep(10 * time.Millisecond)
			close(ch)
		}()

		passes(t, func(t T, r Report) {
			WillClose(t, r, ch, 100*time.Millisecond)
		})
	})

	t.Run("channel receives a signal", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		ch <- struct{}{}

		passes(t, func(t T, r Report) {
			WillClose(t, r, ch, 100*time.Millisecond)
		})
	})

	t.Run("reports a timeout", func(t *testing.T) {
		reports(t, "timeout: channel did not close within 10ms",
			func(t T, r Report) {
				WillClose(t, r, make(chan struct{}), 10*time.Millisecond)
			})
	})
}
