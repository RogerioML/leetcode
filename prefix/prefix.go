package prefix

func LongestCommonPrefix(strs []string) string {
	substrs := make([][]byte, len(strs))
	for i, str := range strs {
		substrs[i] = []byte(str)
	}
	var res []byte
	for i, b := range substrs[0] {
		for j := 1; j < len(substrs); j++ {
			if i >= len(substrs[j]) || substrs[j][i] != b {
				return string(res)
			}
		}
		res = append(res, b)
	}
	return string(res)
}
