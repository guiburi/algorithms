package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	cases := []struct {
		in     string
		result bool
	}{
		{"abba", true},
		{"abcdfe", false},
		{"madam", true},
	}

	for _, c := range cases {
		got := Palindrome(c.in)
		if got != c.result {
			t.Errorf("Palindrome(%q) == %v, want %v", c.in, got, c.result)
		}
	}
}