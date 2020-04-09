package calculator

import (
	"bytes"
	"testing"
)

func TestMock(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"0", "Num 0\n"},
		{"8", "Num 8\n"},
		{"1+2", "Num 1\nAdd\nNum 2\n"},
	}

	for _, c := range cases {
		buf := &bytes.Buffer{}
		m := NewMock(buf)

		InterpretString(c.in, m)

		got := buf.String()
		if got != c.want {
			t.Errorf("%q => %q, want %q.\n", c.in, got, c.want)
			return
		}
	}
}
