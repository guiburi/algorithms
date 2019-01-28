package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		d      []int
		v      int
		result bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{100, 99, 88, 54, 20}, 54, false},
		{[]int{-1, 0, 3, 10, 1002}, -1, true},
	}

	for _, c := range cases {
		got := BinarySearch(c.d, c.v)
		if got != c.result {
			t.Errorf("BinarySearch(%v) == %v, want %v", c.d, got, c.result)
		}
	}
}
