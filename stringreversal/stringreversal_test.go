package stringreversal

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"apple", "elppa"},
		{"hello", "olleh"},
		{"Greetings!", "!sgniteerG"},
	}

	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
