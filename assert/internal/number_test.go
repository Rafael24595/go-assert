package internal

import (
	"math"
	"testing"
)

func TestCompareMagnitudes(t *testing.T) {
	tests := []struct {
		name string
		a    any
		b    any
		want int
		ok   bool
	}{
		{"int equal", 1, 1, 0, true},
		{"int less", 1, 2, -1, true},
		{"int greater", 2, 1, 1, true},

		{"uint equal", 1, 1, 0, true},
		{"uint less", 1, 2, -1, true},
		{"uint greater", 2, 1, 1, true},

		{"float equal", 3.0, 3.0, 0, true},
		{"float less", 3.0, 4.0, -1, true},
		{"float greater", 4.0, 3.0, 1, true},

		{"signed == unsigned", 5, 5, 0, true},
		{"signed < unsigned", 4, 5, -1, true},
		{"signed > unsigned", 6, 5, 1, true},

		{"negative < unsigned", -1, 0, -1, true},
		{"unsigned > negative", 0, -1, 1, true},

		{"max uint > zero", uint64(math.MaxUint64), 0, 1, true},
		{"zero < max uint", 0, uint64(math.MaxUint64), -1, true},

		{"float == int", 3.0, 3, 0, true},
		{"float > int", 3.1, 3, 1, true},
		{"float < int", 2.9, 3, -1, true},

		{"float == uint", 3.0, 3, 0, true},
		{"float > uint", 3.1, 3, 1, true},
		{"float < uint", 2.9, 3, -1, true},

		{"2^53", float64(1 << 53), uint64(1 << 53), 0, true},
		{"2^53+1", float64((1 << 53) + 1), uint64((1 << 53) + 1), 0, false},

		{"NaN", math.NaN(), 0, 0, false},
		{"NaN NaN", math.NaN(), math.NaN(), 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := CompareMagnitude(tt.a, tt.b)

			if ok != tt.ok {
				t.Fatalf("expected ok=%v, got %v", tt.ok, ok)
			}

			if !ok {
				return
			}

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func TestCompareNumbersSymmetry(t *testing.T) {
	values := []any{
		-10,
		0,
		10,
		int64(math.MaxInt64),

		uint(0),
		uint(10),
		uint64(math.MaxUint32),

		3.14,
		-2.5,

		"hello",
		[]int{1, 2, 3},
	}

	for _, a := range values {
		for _, b := range values {
			ab, okAB := CompareMagnitude(a, b)
			ba, okBA := CompareMagnitude(b, a)

			if okAB != okBA {
				t.Fatalf("%#v vs %#v: ok mismatch", a, b)
			}

			if !okAB {
				continue
			}

			if ab != -ba {
				t.Fatalf("%#v vs %#v: %d != -(%d)", a, b, ab, ba)
			}
		}
	}
}

func TestMagnitudeOf(t *testing.T) {
	type MyInt int
	type MyUint uint
	type MyFloat float64

	tests := []struct {
		name string
		in   any
		want any
		ok   bool
	}{
		{"int", int(42), int64(42), true},
		{"uint", uint(42), uint64(42), true},
		{"float", float64(3.14), float64(3.14), true},

		{"custom int", MyInt(10), int64(10), true},
		{"custom uint", MyUint(10), uint64(10), true},
		{"custom float", MyFloat(1.5), float64(1.5), true},

		{"string", "hello", uint64(5), true},
		{"slice", []int{1, 2, 3}, uint64(3), true},
		{"array", [4]int{}, uint64(4), true},
		{"map", map[string]int{"a": 1, "b": 2}, uint64(2), true},

		{"channel", func() chan int {
			ch := make(chan int, 3)
			ch <- 1
			ch <- 2
			return ch
		}(), uint64(2), true},

		{"bool", true, nil, false},
		{"struct", struct{}{}, nil, false},
		{"nil", nil, nil, false},

		{"pointer", &[]int{1, 2, 3}, uint64(3), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := MagnitudeOf(tt.in)

			if ok != tt.ok {
				t.Fatalf("expected ok=%v, got %v", tt.ok, ok)
			}

			if !ok {
				return
			}

			if got != tt.want {
				t.Fatalf("expected %v (%T), got %v (%T)",
					tt.want, tt.want,
					got, got,
				)
			}
		})
	}
}

func TestCapacityOf(t *testing.T) {
	type MyInt int

	ch := make(chan int, 5)
	ch <- 1
	ch <- 2

	tests := []struct {
		name string
		in   any
		want any
		ok   bool
	}{
		{"int", int(42), int64(42), true},
		{"custom int", MyInt(10), int64(10), true},

		{"slice (len != cap)", make([]int, 2, 8), uint64(8), true},
		{"array", [4]int{}, uint64(4), true},
		{"channel", ch, uint64(5), true},

		{"pointer to slice", &[]int{1, 2, 3}, uint64(3), true},
		{"pointer to array", &[5]int{}, uint64(5), true},

		{"string (unsupported)", "hello", nil, false},
		{"map (unsupported)", map[string]int{"a": 1, "b": 2}, nil, false},

		{"bool", true, nil, false},
		{"struct", struct{}{}, nil, false},
		{"nil", nil, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := CapacityOf(tt.in)

			if ok != tt.ok {
				t.Fatalf("expected ok=%v, got %v", tt.ok, ok)
			}

			if !ok {
				return
			}

			if got != tt.want {
				t.Fatalf("expected %v (%T), got %v (%T)",
					tt.want, tt.want,
					got, got,
				)
			}
		})
	}
}

func TestCompareCapacity(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 1

	slice := make([]int, 2, 5)

	tests := []struct {
		name string
		a    any
		b    any
		want int
		ok   bool
	}{
		{"slice cap vs int equal", slice, 5, 0, true},
		{"slice cap vs int less", slice, 10, -1, true},
		{"slice cap vs int greater", slice, 2, 1, true},

		{"chan cap vs slice cap", ch, slice, 1, true},

		{"slice len=2 cap=5 vs int 2", slice, 2, 1, true},
		{"chan len=1 cap=10 vs int 1", ch, 1, 1, true},

		{"signed vs unsigned", -1, uint(0), -1, true},
		{"float vs int", 3.0, 3, 0, true},

		{"string vs int", "hello", 5, 0, false},
		{"map vs int", map[string]int{"a": 1}, 1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := CompareCapacity(tt.a, tt.b)

			if ok != tt.ok {
				t.Fatalf("expected ok=%v, got %v", tt.ok, ok)
			}

			if !ok {
				return
			}

			if got != tt.want {
				t.Fatalf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func TestCompareCapacitySymmetry(t *testing.T) {
	values := []any{
		make([]int, 1, 5),
		make(chan int, 10),
		[3]int{},
		0,
		10,
		"no-supported-string",
	}

	for _, a := range values {
		for _, b := range values {
			ab, okAB := CompareCapacity(a, b)
			ba, okBA := CompareCapacity(b, a)

			if okAB != okBA {
				t.Fatalf("%#v vs %#v: ok mismatch", a, b)
			}

			if !okAB {
				continue
			}

			if ab != -ba {
				t.Fatalf("%#v vs %#v: %d != -(%d)", a, b, ab, ba)
			}
		}
	}
}
