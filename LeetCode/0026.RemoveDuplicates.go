package LeetCode

func RemoveDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if j != 0 && nums[i] == nums[j-1] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}
