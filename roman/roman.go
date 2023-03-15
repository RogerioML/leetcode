package roman

func RomanToInt(s string) int {
	romans := map[byte]int{'I': 1,'V': 5,'X': 10,'L': 50,'C': 100,'D': 500,'M': 1000}
	nums := []byte(s)
	res := romans[nums[len(nums)-1]]
	for i := 0; i < len(nums)-1; i++ {
		if romans[nums[i]] < romans[nums[i+1]] {
			res -= romans[nums[i]]
		} else {
			res += romans[nums[i]]
		}
	}
	return res 
}
