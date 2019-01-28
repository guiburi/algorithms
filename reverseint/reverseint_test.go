package reverseint

import "testing"

func TestReverseInt(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{15, 51},
		{981, 189},
		{-15, -51},
		{-90, -9},
	}

	for _, c := range cases {
		got := ReverseInt(c.in)
		if got != c.want {
			t.Errorf("ReverseInt(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
