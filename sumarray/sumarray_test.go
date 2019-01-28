package sumarray

import "testing"

func TestSumarray(t *testing.T) {
	cases := []struct {
		in     []int
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{3, 3, 3, 3, 3}, 15},
		{[]int{0, -1, 3, 10, 1002}, 1014},
	}

	for _, c := range cases {
		got := Sumarray(c.in)
		if got != c.result {
			t.Errorf("Sum(%v) == %v, want %v", c.in, got, c.result)
		}
	}
}
