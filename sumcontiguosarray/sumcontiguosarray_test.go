package sumcontiguosarray

import "testing"

func TestMaxSubArraySum(t *testing.T) {
	cases := []struct {
		in     []int
		result int
	}{
		{[]int{1, -2, 3, 4, -4, 6, -14, 8, 2}, 10},
		{[]int{10, -2, 32, -44, -4, 6, -14, 8, 2}, 40},
	}

	for _, c := range cases {
		got := MaxSubArraySum(c.in)
		if got != c.result {
			t.Errorf("Sum(%v) == %v, want %v", c.in, got, c.result)
		}
	}
}
