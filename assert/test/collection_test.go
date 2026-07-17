package assert

import "testing"

func TestInside(t *testing.T) {
	t.Run("Strings", func(t *testing.T) {
		Inside(t, "awesome", "Go is awesome")
	})

	t.Run("Slices", func(t *testing.T) {
		list := []int{10, 20, 30}
		Inside(t, 20, list)
		NotInside(t, 40, list)
	})

	t.Run("Slice of Slices", func(t *testing.T) {
		matrix := [][]int{{1, 2}, {3, 4}}
		target := []int{3, 4}
		Inside(t, target, matrix)
	})

	t.Run("Slice of Structs with Maps", func(t *testing.T) {
		type data struct {
			ID   int
			Tags map[string]bool
		}

		list := []data{
			{ID: 1, Tags: map[string]bool{"active": true}},
			{ID: 2, Tags: map[string]bool{"admin": false}},
		}

		target := data{ID: 2, Tags: map[string]bool{"admin": false}}
		Inside(t, target, list)
	})

	t.Run("Map keys", func(t *testing.T) {
		m := map[string]int{"A": 1, "B": 2}

		Inside(t, "A", m)
		NotInside(t, "C", m)
	})

	t.Run("Map with struct keys", func(t *testing.T) {
		type ID struct{ Num int }
		m := map[ID]bool{{Num: 1}: true}

		Inside(t, ID{Num: 1}, m)
	})
}

func TestContains(t *testing.T) {
	t.Run("Strings", func(t *testing.T) {
		Contains(t, "Go is awesome", "awesome")
	})

	t.Run("Slices", func(t *testing.T) {
		list := []int{10, 20, 30}
		Contains(t, list, 20)
		NotContains(t, list, 40)
	})

	t.Run("Slice of Slices", func(t *testing.T) {
		matrix := [][]int{{1, 2}, {3, 4}}
		target := []int{3, 4}
		Contains(t, matrix, target)
	})

	t.Run("Slice of Structs with Maps", func(t *testing.T) {
		type data struct {
			ID   int
			Tags map[string]bool
		}

		list := []data{
			{ID: 1, Tags: map[string]bool{"active": true}},
			{ID: 2, Tags: map[string]bool{"admin": false}},
		}

		target := data{ID: 2, Tags: map[string]bool{"admin": false}}
		Contains(t, list, target)
	})

	t.Run("Map keys", func(t *testing.T) {
		m := map[string]int{"A": 1, "B": 2}

		Contains(t, m, "A")
		NotContains(t, m, "C")
	})

	t.Run("Map with struct keys", func(t *testing.T) {
		type ID struct{ Num int }
		m := map[ID]bool{{Num: 1}: true}

		Contains(t, m, ID{Num: 1})
	})
}
