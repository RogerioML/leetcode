package prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"hello", "hello world", "help"}, "hel"},
		{[]string{"", "test"}, ""},
		{[]string{"apple", "apple", "apple"}, "apple"},
		{[]string{"roger", "rogerio", "rogiero"}, "rog"},
	}
	for _, c := range cases {
		got := LongestCommonPrefix(c.strs)
		if got != c.want {
			t.Errorf("longestCommonPrefix(%v) == %s, want %s", c.strs, got, c.want)
		}
	}
}
