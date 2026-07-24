package assert

import (
	"math"
	"testing"
)

func TestSize(t *testing.T) {
	Size(t, 0, "")
	Size(t, 0, []int{})
	Size(t, 3, [3]int{1, 2, 3})
	Size(t, 1, map[int]string{1: "one"})
	Size(t, 0, make(chan int))
}

func TestSizeExtended(t *testing.T) {
	t.Run("String length", func(t *testing.T) {
		Size(t, 6, "Gopher")
	})
	t.Run("Map length", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b"}
		Size(t, 2, m)
	})
	t.Run("Channel length", func(t *testing.T) {
		ch := make(chan int, 5)
		ch <- 1
		ch <- 2
		Size(t, 2, ch)
	})
}

func TestEmpty(t *testing.T) {
	Empty(t, "")
	Empty(t, []int{})
	Empty(t, make(map[int]string))
}

func TestNotEmpty(t *testing.T) {
	NotEmpty(t, "golang")
	NotEmpty(t, [3]int{1, 2, 3})
	NotEmpty(t, map[int]string{1: "one"})
}

func TestLessThan(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		LessThan(t, 20, 10)
		LessThan(t, 20.5, 10)
		LessThan(t, 0, -10)
	})

	t.Run("string length", func(t *testing.T) {
		LessThan(t, 10, "go")
	})

	t.Run("slice length", func(t *testing.T) {
		LessThan(t, 5, []int{1, 2})
	})

	t.Run("map length", func(t *testing.T) {
		LessThan(t, 3, map[string]int{
			"go":   1,
			"rust": 2,
		})
	})
}

func TestLessOrEqual(t *testing.T) {
	t.Run("equal numbers", func(t *testing.T) {
		LessOrEqual(t, 10, 10)
		LessOrEqual(t, 10.5, 10)
	})

	t.Run("string length", func(t *testing.T) {
		LessOrEqual(t, 5, "hello")
	})

	t.Run("map length", func(t *testing.T) {
		LessOrEqual(t, 2, map[string]int{
			"a": 1,
			"b": 2,
		})
	})

	t.Run("map length", func(t *testing.T) {
		LessOrEqual(t, 2, map[string]int{
			"go":   1,
			"rust": 2,
		})
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		GreaterThan(t, 5, 10)
		GreaterThan(t, 5.5, 10)
		GreaterThan(t, -10, 0)
		GreaterThan(t, 0, uint64(math.MaxUint64))
	})

	t.Run("string length", func(t *testing.T) {
		GreaterThan(t, 3, "golang")
	})

	t.Run("slice length", func(t *testing.T) {
		GreaterThan(t, 2, []string{
			"go",
			"rust",
			"zig",
		})
	})

	t.Run("map length", func(t *testing.T) {
		GreaterThan(t, 1, map[string]int{
			"go":   1,
			"rust": 2,
		})
	})
}

func TestGreaterOrEqual(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		GreaterOrEqual(t, 5, 5)
		GreaterOrEqual(t, 5, 6)
		GreaterOrEqual(t, 5.5, 6)
		GreaterOrEqual(t, -5, 6)
	})

	t.Run("string length", func(t *testing.T) {
		GreaterOrEqual(t, 6, "golang")
	})

	t.Run("slice length", func(t *testing.T) {
		GreaterOrEqual(t, 3, []string{
			"golang",
			"zig",
			"rust",
		})
	})

	t.Run("map length", func(t *testing.T) {
		GreaterOrEqual(t, 2, map[string]int{
			"go":   1,
			"rust": 2,
		})
	})

	t.Run("custom type", func(t *testing.T) {
		var a customInt = 10
		Greater(t, customInt(5), a)
	})
}

func TestCapacity(t *testing.T) {
	t.Run("slice capacity", func(t *testing.T) {
		s := make([]int, 2, 10)
		Capacity(t, 10, s)
	})

	t.Run("array capacity", func(t *testing.T) {
		arr := [5]int{1, 2, 3}
		Capacity(t, 5, arr)
	})

	t.Run("channel capacity", func(t *testing.T) {
		ch := make(chan int, 8)
		ch <- 1
		ch <- 2
		Capacity(t, 8, ch)
	})

	t.Run("different numeric types for want", func(t *testing.T) {
		s := make([]string, 0, 4)
		Capacity(t, uint(4), s)
		Capacity(t, float64(4.0), s)
		Capacity(t, int64(4), s)
	})

	t.Run("pointer to slice or array", func(t *testing.T) {
		s := make([]int, 0, 6)
		arr := [4]int{1, 2}
		Capacity(t, 6, &s)
		Capacity(t, 4, &arr)
	})
}
