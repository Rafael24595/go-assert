package check

import (
	"errors"
	"strings"
	"testing"
)

func TestNil(t *testing.T) {
	t.Run("nil interfaces and pointers", func(t *testing.T) {
		var err error = (*customError)(nil)
		var p *int

		passes(t, func(t T, r Report) {
			Nil(t, r, err)
		})

		passes(t, func(t T, r Report) {
			Nil(t, r, p)
		})

		passes(t, func(t T, r Report) {
			Nil(t, r, nil)
		})
	})

	t.Run("nil collections", func(t *testing.T) {
		var s []int
		var m map[string]int
		var c chan int

		passes(t, func(t T, r Report) {
			Nil(t, r, s)
		})

		passes(t, func(t T, r Report) {
			Nil(t, r, m)
		})

		passes(t, func(t T, r Report) {
			Nil(t, r, c)
		})
	})

	t.Run("reports the value and its type", func(t *testing.T) {
		reports(t, "Expected nil, but got boom (*errors.errorString)",
			func(t T, r Report) {
				Nil(t, r, errors.New("boom"))
			})
	})

	t.Run("prefixes the caller's message", func(t *testing.T) {
		reports(t, "loading user 7 - Expected nil, but got 0 (int)",
			func(t T, r Report) {
				Nil(t, r, 0, "loading user %d", 7)
			})
	})
}

func TestNotNil(t *testing.T) {
	t.Run("zero values are not nil", func(t *testing.T) {
		for _, v := range []any{0, "", false, struct{}{}} {
			passes(t, func(t T, r Report) {
				NotNil(t, r, v)
			})
		}
	})

	t.Run("empty collections are not nil", func(t *testing.T) {
		for _, v := range []any{[]int{}, make(map[string]int), make(chan int)} {
			passes(t, func(t T, r Report) {
				NotNil(t, r, v)
			})
		}
	})

	t.Run("reports a nil pointer", func(t *testing.T) {
		var p *int
		reports(t, "Unexpected nil value", func(t T, r Report) {
			NotNil(t, r, p)
		})
	})
}

func TestBoolean(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			True(t, r, 1 < 2)
		})

		passes(t, func(t T, r Report) {
			True(t, r, strings.HasPrefix("Gopher", "Go"))
		})

		reports(t, "Expected true, but got false", func(t T, r Report) {
			True(t, r, false)
		})
	})

	t.Run("False", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			False(t, r, 1 > 2)
		})

		passes(t, func(t T, r Report) {
			False(t, r, strings.Contains("Go", "Java"))
		})

		reports(t, "Expected false, but got true", func(t T, r Report) {
			False(t, r, true)
		})
	})
}

func TestEqual(t *testing.T) {
	t.Run("multiple types", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			Equal(t, r, 123, 123)
		})

		passes(t, func(t T, r Report) {
			Equal(t, r, "gopher", "gopher")
		})

		passes(t, func(t T, r Report) {
			Equal(t, r, true, true)
		})
	})

	t.Run("custom type", func(t *testing.T) {
		passes(t, func(t T, r Report) {
			Equal(t, r, customInt(10), customInt(10))
		})
	})

	t.Run("reports both values", func(t *testing.T) {
		reports(t, "Expected '1', but got '2'", func(t T, r Report) {
			Equal(t, r, 1, 2)
		})
	})
}

func TestNotEqual(t *testing.T) {
	passes(t, func(t T, r Report) {
		NotEqual(t, r, 10, 20)
	})

	passes(t, func(t T, r Report) {
		NotEqual(t, r, "ziglang", "golang")
	})

	reports(t, "Unexpected 'go'", func(t T, r Report) {
		NotEqual(t, r, "go", "go")
	})
}

func TestDeepEqual(t *testing.T) {
	type person struct {
		Name string
		Meta map[string]int
	}

	p1 := person{Name: "Gopher", Meta: map[string]int{"age": 10}}
	p2 := person{Name: "Gopher", Meta: map[string]int{"age": 10}}

	passes(t, func(t T, r Report) {
		DeepEqual(t, r, p1, p2)
	})

	passes(t, func(t T, r Report) {
		DeepEqual(t, r, []int{1, 2, 3}, []int{1, 2, 3})
	})

	reports(t, "Expected deep equality.\nWant: [1 2]\nGot:  [1 2 3]",
		func(t T, r Report) {
			DeepEqual(t, r, []int{1, 2}, []int{1, 2, 3})
		})
}

func TestNotDeepEqual(t *testing.T) {
	type user struct {
		Name string
		Tags []string
	}

	u1 := user{Name: "Gopher", Tags: []string{"go", "backend"}}
	u2 := user{Name: "Ferris", Tags: []string{"rust", "systems"}}
	u3 := user{Name: "Gopher", Tags: []string{"go", "backend"}}

	passes(t, func(t T, r Report) {
		NotDeepEqual(t, r, u1, u2)
	})

	passes(t, func(t T, r Report) {
		NotDeepEqual(t, r, []int{1, 2}, []int{1, 2, 3})
	})

	reports(t, "Expected values to be deeply different.\nWant: {Name:Gopher Tags:[go backend]}\nGot:  {Name:Gopher Tags:[go backend]}",
		func(t T, r Report) {
			NotDeepEqual(t, r, u1, u3)
		})
}

func TestSame(t *testing.T) {
	type data struct{ Value int }

	obj1 := &data{Value: 42}
	obj2 := &data{Value: 42}
	obj3 := obj1

	slice1 := []int{1, 2}
	slice2 := slice1

	passes(t, func(t T, r Report) {
		Same(t, r, obj1, obj3)
	})

	passes(t, func(t T, r Report) {
		Same(t, r, slice1, slice2)
	})

	t.Run("same value, different pointers", func(t *testing.T) {
		reports(t, "Expected both values to reference the same object",
			func(t T, r Report) {
				Same(t, r, obj1, obj2)
			})
	})

	t.Run("value types are a misuse", func(t *testing.T) {
		stops(t, "Same only supports reference types", func(t T, r Report) {
			Same(t, r, 10, 10)
		})
	})
}

func TestNotSame(t *testing.T) {
	type data struct{ Value int }

	obj1 := &data{Value: 10}
	obj2 := &data{Value: 10}
	obj3 := obj1

	passes(t, func(t T, r Report) {
		NotSame(t, r, obj1, obj2)
	})

	passes(t, func(t T, r Report) {
		NotSame(t, r, []int{1, 2}, []int{1, 2})
	})

	t.Run("same object", func(t *testing.T) {
		reports(t, "Expected different references", func(t T, r Report) {
			NotSame(t, r, obj1, obj3)
		})
	})

	t.Run("value types are a misuse", func(t *testing.T) {
		stops(t, "NotSame only supports reference types", func(t T, r Report) {
			NotSame(t, r, 10, 10)
		})
	})
}

func TestInDelta(t *testing.T) {
	passes(t, func(t T, r Report) {
		InDelta(t, r, 0.3, 0.1+0.2, 0.00001)
	})

	passes(t, func(t T, r Report) {
		InDelta(t, r, 100.0, 100.05, 0.1)
	})

	reports(t, "Expected 1.000000 and 1.500000 to be within 0.100000 of each other",
		func(t T, r Report) {
			InDelta(t, r, 1.0, 1.5, 0.1)
		})
}
