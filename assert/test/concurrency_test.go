package assert

import (
	"testing"
	"time"
)

func TestWillClose(t *testing.T) {
	t.Run("Success when channel closes in time", func(t *testing.T) {
		ch := make(chan struct{})

		go func() {
			time.Sleep(10 * time.Millisecond)
			close(ch)
		}()

		WillClose(t, ch, 100*time.Millisecond)
	})

	t.Run("Success when channel receives signal", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		ch <- struct{}{}

		WillClose(t, ch, 100*time.Millisecond)
	})
}
