package prefix

import "fmt"

func LongestCommonPrefix(strs []string) string {
	var res string

	substrs := make([][]byte, len(strs))
	for i, str := range strs {
		substrs[i] = []byte(str)
	}

	for _, sub := range substrs {
		fmt.Println(sub)
	}

	return res
}
