package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	cases := []struct {
		in     []int
		value  int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},

		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
	}

	for _, c := range cases {
		got := RotateArray(c.in, c.value)
		if !reflect.DeepEqual(got, c.result) {
			t.Errorf("RotateArray(%v) == %v, want %v", c.in, got, c.result)
		}
	}
}
