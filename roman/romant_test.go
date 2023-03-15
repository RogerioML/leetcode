package roman

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		roman string
		want  int
	}{
		{"I", 1},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, c := range cases {
		got := RomanToInt(c.roman)
		if got != c.want {
			t.Errorf("RomanToInt(%q) == %d, want %d", c.roman, got, c.want)
		}
	}
}
