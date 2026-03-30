package assert

import (
	"fmt"
	"strings"
	"testing"
)

type MyInt int

type customError struct{}
func (e *customError) Error() string { return "" }

func TestNilDeep(t *testing.T) {
    t.Run("Nil interfaces and pointers", func(t *testing.T) {
        var err error = (*customError)(nil)
        var p *int = nil
        Nil(t, err)
        Nil(t, p)
        Nil(t, nil)
    })

    t.Run("Nil collections", func(t *testing.T) {
        var s []int = nil
        var m map[string]int = nil
        var c chan int = nil
        Nil(t, s)
        Nil(t, m)
        Nil(t, c)
    })
}

func TestNotNilDeep(t *testing.T) {
    t.Run("Nil zero values", func(t *testing.T) {
        NotNil(t, 0)
        NotNil(t, "")
        NotNil(t, false)
        NotNil(t, struct{}{})
    })

    t.Run("Nil empty collections", func(t *testing.T) {
        NotNil(t, []int{}) 
        NotNil(t, make(map[string]int))
        NotNil(t, make(chan int))
    })
}

func TestBoolean(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		True(t, 1 < 2)
		True(t, strings.HasPrefix("Gopher", "Go"))
	})
	t.Run("False", func(t *testing.T) {
		False(t, 1 > 2)
		False(t, strings.Contains("Go", "Java"))
	})
}

func TestEqual(t *testing.T) {
	Equal(t, 123, 123)
	Equal(t, "gopher", "gopher")
	Equal(t, true, true)
}

func TestNotEqual(t *testing.T) {
    NotEqual(t, 10, 20)
    NotEqual(t, "ziglang", "golang")
}

func TestNumericComparisons(t *testing.T) {
	Greater(t, 5, 10)
	Less(t, 20, 10)
	GreaterOrEqual(t, 5, 5)
	LessOrEqual(t, 10, 10)
}

func TestInDelta(t *testing.T) {
	InDelta(t, 0.3, 0.1+0.2, 0.00001)
	InDelta(t, 100.0, 100.05, 0.1)
}

func TestDeepEqual(t *testing.T) {
	type person struct {
		Name string
		Meta map[string]int
	}

	p1 := person{Name: "Gopher", Meta: map[string]int{"age": 10}}
	p2 := person{Name: "Gopher", Meta: map[string]int{"age": 10}}

	DeepEqual(t, p1, p2)

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	DeepEqual(t, s1, s2)
}

func TestPanic(t *testing.T) {
	t.Run("Should panic", func(t *testing.T) {
		Panic(t, func() {
			panic("boom")
		})
	})

	t.Run("Panic with message", func(t *testing.T) {
		PanicWithMessage(t, "error crítico", func() {
			panic("error crítico")
		})
	})

	t.Run("Should not panic", func(t *testing.T) {
		NotPanic(t, func() {
			_ = 1 + 1
		})
	})
}

func TestCustomTypes(t *testing.T) {
	var a MyInt = 10
	var b MyInt = 10
	Equal(t, a, b)
	Greater(t, MyInt(5), a)
}

func TestErrors(t *testing.T) {
	t.Run("Error exists", func(t *testing.T) {
		err := fmt.Errorf("fail")
		Error(t, err)
	})

	t.Run("No error expected", func(t *testing.T) {
		var err error = nil
		NotError(t, err)
	})
}

func TestLen(t *testing.T) {
	Len(t, 0, "")
	Len(t, 0, []int{})
	Len(t, 3, [3]int{1, 2, 3})
	Len(t, 0, make(chan int))
}

func TestLenExtended(t *testing.T) {
    t.Run("String length", func(t *testing.T) {
        Len(t, 6, "Gopher")
    })
    t.Run("Map length", func(t *testing.T) {
        m := map[int]string{1: "a", 2: "b"}
        Len(t, 2, m)
    })
    t.Run("Channel length", func(t *testing.T) {
        ch := make(chan int, 5)
        ch <- 1
        ch <- 2
        Len(t, 2, ch)
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
