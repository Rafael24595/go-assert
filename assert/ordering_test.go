package assert

import (
	"testing"
)

func TestOrderingFailures(t *testing.T) {
	expectFailures(t, []failure{
		{"Size", func(t T) {
			Size(t, 2, []int{1}, "golang %d", 125)
		}},
		{"Empty", func(t T) {
			Empty(t, []int{1}, "golang %d", 125)
		}},
		{"NotEmpty", func(t T) {
			NotEmpty(t, []int{}, "golang %d", 125)
		}},
		{"LessThan", func(t T) {
			LessThan(t, 1, 2, "golang %d", 125)
		}},
		{"LessOrEqual", func(t T) {
			LessOrEqual(t, 1, 2, "golang %d", 125)
		}},
		{"GreaterThan", func(t T) {
			GreaterThan(t, 2, 1, "golang %d", 125)
		}},
		{"GreaterOrEqual", func(t T) {
			GreaterOrEqual(t, 2, 1, "golang %d", 125)
		}},
		{"Capacity", func(t T) {
			Capacity(t, 3, make([]int, 0, 1), "golang %d", 125)
		}},
	})
}

func TestOrderingAlwaysStops(t *testing.T) {
	expectStops(t, []failure{
		{"Size misuse", func(t T) {
			Size(t, 1, true)
		}},
	})
}
