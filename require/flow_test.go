package require

import (
	"testing"
)

func TestFlowAlwaysStops(t *testing.T) {
	expectStops(t, []failure{
		{"Unreachable", func(t T) {
			Unreachable(t, "golang %d", 125)
		}},
	})
}
