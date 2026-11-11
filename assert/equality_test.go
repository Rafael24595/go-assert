package assert

import (
	"testing"
)

func TestEqualityFailures(t *testing.T) {
	expectFailures(t, []failure{
		{"Nil", func(t T) {
			Nil(t, 1, "golang %d", 125)
		}},
		{"NotNil", func(t T) {
			NotNil(t, nil, "golang %d", 125)
		}},
		{"True", func(t T) {
			True(t, false, "golang %d", 125)
		}},
		{"False", func(t T) {
			False(t, true, "golang %d", 125)
		}},
		{"Equal", func(t T) {
			Equal(t, 1, 2, "golang %d", 125)
		}},
		{"NotEqual", func(t T) {
			NotEqual(t, 1, 1, "golang %d", 125)
		}},
		{"DeepEqual", func(t T) {
			DeepEqual(t, []int{1}, []int{2}, "golang %d", 125)
		}},
		{"NotDeepEqual", func(t T) {
			NotDeepEqual(t, []int{1}, []int{1}, "golang %d", 125)
		}},
		{"Same", func(t T) {
			Same(t, new(int), new(int), "golang %d", 125)
		}},
		{"NotSame", func(t T) {
			p := new(int)
			NotSame(t, p, p, "golang %d", 125)
		}},
		{"InDelta", func(t T) {
			InDelta(t, 1, 2, 0.1, "golang %d", 125)
		}},
	})
}
