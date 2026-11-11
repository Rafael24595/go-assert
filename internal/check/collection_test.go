package check

import (
	"testing"
)

func TestInside(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			Inside(t, r, "awesome", "Go is awesome")
		})
	})

	t.Run("slices", func(t *testing.T) {
		list := []int{10, 20, 30}

		passes(t, func(t T, r Report) {
			Inside(t, r, 20, list)
		})

		passes(t, func(t T, r Report) {
			NotInside(t, r, 40, list)
		})
	})

	t.Run("slice of slices", func(t *testing.T) {
		matrix := [][]int{{1, 2}, {3, 4}}
		passes(t, func(t T, r Report) {
			Inside(t, r, []int{3, 4}, matrix)
		})
	})

	t.Run("slice of structs with maps", func(t *testing.T) {
		type data struct {
			ID   int
			Tags map[string]bool
		}

		list := []data{
			{ID: 1, Tags: map[string]bool{"active": true}},
			{ID: 2, Tags: map[string]bool{"admin": false}},
		}

		target := data{ID: 2, Tags: map[string]bool{"admin": false}}

		passes(t, func(t T, r Report) {
			Inside(t, r, target, list)
		})
	})

	t.Run("map keys", func(t *testing.T) {
		m := map[string]int{"A": 1, "B": 2}

		passes(t, func(t T, r Report) {
			Inside(t, r, "A", m)
		})

		passes(t, func(t T, r Report) {
			NotInside(t, r, "C", m)
		})
	})

	t.Run("map with struct keys", func(t *testing.T) {
		type ID struct{ Num int }
		m := map[ID]bool{{Num: 1}: true}

		passes(t, func(t T, r Report) {
			Inside(t, r, ID{Num: 1}, m)
		})
	})

	t.Run("reports a missing item", func(t *testing.T) {
		reports(t, "Expected collection to contain '40'",
			func(t T, r Report) {
				Inside(t, r, 40, []int{10, 20})
			})

		reports(t, "Expected collection to contain 'rust'",
			func(t T, r Report) {
				Inside(t, r, "rust", "Go is awesome")
			})
	})

	t.Run("reports a present item", func(t *testing.T) {
		reports(t, "Expected collection NOT to contain '10'",
			func(t T, r Report) {
				NotInside(t, r, 10, []int{10, 20})
			})
	})

	t.Run("unsupported containers are a misuse", func(t *testing.T) {
		stops(t, "Inside does not support type int",
			func(t T, r Report) {
				Inside(t, r, 1, 42)
			})

		stops(t, "Inside does not support type int",
			func(t T, r Report) {
				NotInside(t, r, 1, 42)
			})
	})
}
