package twosum

func twoSum(nums []int, target int) []int {
	var ret []int
	for i := range nums {
		for k := range nums {
			if i != k && nums[i]+nums[k] == target {
				ret = append(ret, i)
			}
		}
	}
	return ret
}
