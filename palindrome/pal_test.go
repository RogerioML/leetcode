package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input int
		want  bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{12321, true},
		{0, true},
	}
	for _, c := range cases {
		got := IsPalindrome(c.input)
		if got != c.want {
			t.Errorf("IsPalindrome(%d) == %t, want %t", c.input, got, c.want)
		}
	}
}
