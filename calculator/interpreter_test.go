package calculator

import (
	"testing"
)

func TestInterpreter(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"=", 0},
		{"1+1", 1},
		{"1+1=", 2},
		{"3*4=", 12},
	}

	for _, c := range cases {
		i := NewInterpreter()

		InterpretString(c.in, i)

		got := i.GetResult()
		if got != c.want {
			t.Errorf("%q => %v, want %v.i\n", c.in, got, c.want)
		}
	}
}
