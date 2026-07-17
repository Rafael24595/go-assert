package assert

import (
	"strings"
	"testing"
)

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
	t.Run("multiple types", func(t *testing.T) {
		Equal(t, 123, 123)
		Equal(t, "gopher", "gopher")
		Equal(t, true, true)
	})

	t.Run("custom type", func(t *testing.T) {
		var a customInt = 10
		var b customInt = 10
		Equal(t, a, b)
	})
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 10, 20)
	NotEqual(t, "ziglang", "golang")
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

func TestNotDeepEqual(t *testing.T) {
	type user struct {
		Name string
		Tags []string
	}

	u1 := user{Name: "Gopher", Tags: []string{"go", "backend"}}
	u2 := user{Name: "Ferris", Tags: []string{"rust", "systems"}}
	//u3 := user{Name: "Gopher", Tags: []string{"go", "backend"}}

	NotDeepEqual(t, u1, u2)
	NotDeepEqual(t, []int{1, 2}, []int{1, 2, 3})

	/*t.Run("Should fail if they are deeply equal", func(subT *testing.T) {
		NotDeepEqual(subT, u1, u3)
	})*/
}

func TestSame(t *testing.T) {
	type data struct{ Value int }

	obj1 := &data{Value: 42}
	//obj2 := &data{Value: 42}
	obj3 := obj1

	Same(t, obj1, obj3)

	slice1 := []int{1, 2}
	slice2 := slice1
	Same(t, slice1, slice2)

	/*t.Run("Should fail if they have the same value but different pointers", func(subT *testing.T) {
		Same(subT, obj1, obj2)
	})*/

	/*t.Run("Should fail if not a reference type", func(subT *testing.T) {
		Same(subT, 10, 10)
	})*/
}

func TestNotSame(t *testing.T) {
	type data struct{ Value int }

	obj1 := &data{Value: 10}
	obj2 := &data{Value: 10}
	//obj3 := obj1

	NotSame(t, obj1, obj2)

	s1 := []int{1, 2}
	s2 := []int{1, 2}
	NotSame(t, s1, s2)

	/*t.Run("Should fail if they point to the same object", func(subT *testing.T) {
		NotSame(subT, obj1, obj3)
	})*/
}

func TestInDelta(t *testing.T) {
	InDelta(t, 0.3, 0.1+0.2, 0.00001)
	InDelta(t, 100.0, 100.05, 0.1)
}
