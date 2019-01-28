package factorial

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct {
		in     int
		result int
	}{
		{3, 6},
		{5, 120},
		{9, 362880},
	}

	for _, c := range cases {
		got := Factorial(c.in)
		if got != c.result {
			t.Errorf("Palindrome(%q) == %v, want %v", c.in, got, c.result)
		}
	}
}