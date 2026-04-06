package LeetCode

func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[k] = nums[i]
		k++
	}
	return k
}
