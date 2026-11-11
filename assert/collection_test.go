package assert

import (
	"testing"
)

func TestCollectionFailures(t *testing.T) {
	expectFailures(t, []failure{
		{"Inside", func(t T) {
			Inside(t, 3, []int{1}, "golang %d", 125)
		}},
		{"NotInside", func(t T) {
			NotInside(t, 1, []int{1}, "golang %d", 125)
		}},
	})
}
