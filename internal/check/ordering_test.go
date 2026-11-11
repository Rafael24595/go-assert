package check

import (
	"math"
	"testing"
)

func TestSize(t *testing.T) {
	t.Run("supported kinds", func(t *testing.T) {
		ch := make(chan int, 5)
		ch <- 1
		ch <- 2

		passes(t, func(t T, r Report) {
			Size(t, r, 0, "")
		})

		passes(t, func(t T, r Report) {
			Size(t, r, 6, "Gopher")
		})

		passes(t, func(t T, r Report) {
			Size(t, r, 0, []int{})
		})

		passes(t, func(t T, r Report) {
			Size(t, r, 3, [3]int{1, 2, 3})
		})

		passes(t, func(t T, r Report) {
			Size(t, r, 2, map[int]string{1: "a", 2: "b"})
		})

		passes(t, func(t T, r Report) {
			Size(t, r, 0, make(chan int))
		})

		passes(t, func(t T, r Report) {
			Size(t, r, 2, ch)
		})
	})

	t.Run("reports the actual size", func(t *testing.T) {
		reports(t, "Expected 2, but got 1", func(t T, r Report) {
			Size(t, r, 2, []int{1})
		})
	})

	t.Run("unmeasurable values are a misuse", func(t *testing.T) {
		stops(t, "Size(): bool is not measurable or convertible to int",
			func(t T, r Report) {
				Size(t, r, 2, true)
			})
	})
}

func TestEmpty(t *testing.T) {
	passes(t, func(t T, r Report) {
		Empty(t, r, "")
	})

	passes(t, func(t T, r Report) {
		Empty(t, r, []int{})
	})

	passes(t, func(t T, r Report) {
		Empty(t, r, make(map[int]string))
	})

	reports(t, "Expected 0, but got 2", func(t T, r Report) {
		Empty(t, r, []int{1, 2})
	})
}

func TestNotEmpty(t *testing.T) {
	passes(t, func(t T, r Report) {
		NotEmpty(t, r, "golang")
	})

	passes(t, func(t T, r Report) {
		NotEmpty(t, r, [3]int{1, 2, 3})
	})

	passes(t, func(t T, r Report) {
		NotEmpty(t, r, map[int]string{1: "one"})
	})

	reports(t, "Expected greater than 0, but got []", func(t T, r Report) {
		NotEmpty(t, r, []int{})
	})
}

func TestLessThan(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			LessThan(t, r, 20, 10)
		})

		passes(t, func(t T, r Report) {
			LessThan(t, r, 20.5, 10)
		})

		passes(t, func(t T, r Report) {
			LessThan(t, r, 0, -10)
		})
	})

	t.Run("lengths", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			LessThan(t, r, 10, "go")
		})

		passes(t, func(t T, r Report) {
			LessThan(t, r, 5, []int{1, 2})
		})

		passes(t, func(t T, r Report) {
			LessThan(t, r, 3, map[string]int{"go": 1, "rust": 2})
		})
	})

	t.Run("equal is not less", func(t *testing.T) {
		reports(t, "Expected less than 10, but got 10", func(t T, r Report) {
			LessThan(t, r, 10, 10)
		})
	})

	t.Run("unmeasurable values are a misuse", func(t *testing.T) {
		stops(t, "LessThan(): bool is not measurable or convertible to int",
			func(t T, r Report) {
				LessThan(t, r, 1, true)
			})
	})
}

func TestLessOrEqual(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			LessOrEqual(t, r, 10, 10)
		})

		passes(t, func(t T, r Report) {
			LessOrEqual(t, r, 10.5, 10)
		})
	})

	t.Run("lengths", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			LessOrEqual(t, r, 5, "hello")
		})

		passes(t, func(t T, r Report) {
			LessOrEqual(t, r, 2, map[string]int{"a": 1, "b": 2})
		})
	})

	t.Run("reports a greater value", func(t *testing.T) {
		reports(t, "Expected less or equal than 10, but got 11", func(t T, r Report) {
			LessOrEqual(t, r, 10, 11)
		})
	})

	t.Run("unmeasurable values are a misuse", func(t *testing.T) {
		stops(t, "LessOrEqual(): bool is not measurable or convertible to int",
			func(t T, r Report) {
				LessOrEqual(t, r, 1, true)
			})
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			GreaterThan(t, r, 5, 10)
		})

		passes(t, func(t T, r Report) {
			GreaterThan(t, r, 5.5, 10)
		})

		passes(t, func(t T, r Report) {
			GreaterThan(t, r, -10, 0)
		})

		passes(t, func(t T, r Report) {
			GreaterThan(t, r, 0, uint64(math.MaxUint64))
		})
	})

	t.Run("lengths", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			GreaterThan(t, r, 3, "golang")
		})

		passes(t, func(t T, r Report) {
			GreaterThan(t, r, 2, []string{"go", "rust", "zig"})
		})

		passes(t, func(t T, r Report) {
			GreaterThan(t, r, 1, map[string]int{"go": 1, "rust": 2})
		})
	})

	t.Run("equal is not greater", func(t *testing.T) {
		reports(t, "Expected greater than 10, but got 10", func(t T, r Report) {
			GreaterThan(t, r, 10, 10)
		})
	})

	t.Run("unmeasurable values are a misuse", func(t *testing.T) {
		stops(t, "GreaterThan(): bool is not measurable or convertible to int",
			func(t T, r Report) {
				GreaterThan(t, r, 1, true)
			})
	})
}

func TestGreaterOrEqual(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, 5, 5)
		})

		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, 5, 6)
		})

		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, 5.5, 6)
		})

		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, -5, 6)
		})
	})

	t.Run("lengths", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, 6, "golang")
		})

		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, 3, []string{"golang", "zig", "rust"})
		})

		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, 2, map[string]int{"go": 1, "rust": 2})
		})
	})

	t.Run("custom type", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			GreaterOrEqual(t, r, customInt(5), customInt(10))
		})
	})

	t.Run("reports a smaller value", func(t *testing.T) {
		reports(t, "Expected greater or equal than 10, but got 9", func(t T, r Report) {
			GreaterOrEqual(t, r, 10, 9)
		})
	})

	t.Run("unmeasurable values are a misuse", func(t *testing.T) {
		stops(t, "GreaterOrEqual(): bool is not measurable or convertible to int",
			func(t T, r Report) {
				GreaterOrEqual(t, r, 1, true)
			})
	})
}

func TestCapacity(t *testing.T) {
	t.Run("supported kinds", func(t *testing.T) {
		ch := make(chan int, 8)
		ch <- 1

		passes(t, func(t T, r Report) {
			Capacity(t, r, 10, make([]int, 2, 10))
		})

		passes(t, func(t T, r Report) {
			Capacity(t, r, 5, [5]int{1, 2, 3})
		})

		passes(t, func(t T, r Report) {
			Capacity(t, r, 8, ch)
		})
	})

	t.Run("different numeric types for want", func(t *testing.T) {
		s := make([]string, 0, 4)

		passes(t, func(t T, r Report) {
			Capacity(t, r, uint(4), s)
		})

		passes(t, func(t T, r Report) {
			Capacity(t, r, float64(4.0), s)
		})

		passes(t, func(t T, r Report) {
			Capacity(t, r, int64(4), s)
		})
	})

	t.Run("pointer to slice or array", func(t *testing.T) {
		s := make([]int, 0, 6)
		arr := [4]int{1, 2}

		passes(t, func(t T, r Report) {
			Capacity(t, r, 6, &s)
		})

		passes(t, func(t T, r Report) {
			Capacity(t, r, 4, &arr)
		})
	})

	t.Run("reports the actual capacity", func(t *testing.T) {
		reports(t, "Expected capacity 3, but got 10", func(t T, r Report) {
			Capacity(t, r, 3, make([]int, 0, 10))
		})
	})

	t.Run("values without capacity are a misuse", func(t *testing.T) {
		stops(t, "Capacity(): string is not capacity-measurable or convertible to int",
			func(t T, r Report) {
				Capacity(t, r, 3, "abc")
			})
	})
}
