package sequentialsearch

import "testing"

func TestSequentialSearch(t *testing.T) {
	cases := []struct {
		d      []int
		v      int
		result bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{3, 3, 3, 3, 2}, 5, false},
		{[]int{0, -1, 3, 10, 1002}, -1, true},
	}

	for _, c := range cases {
		got := SequentialSearch(c.d, c.v)
		if got != c.result {
			t.Errorf("SequentialSearch(%v) == %v, want %v", c.d, got, c.result)
		}
	}
}
