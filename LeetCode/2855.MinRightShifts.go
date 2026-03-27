package LeetCode

func MinRightShifts(nums []int) int {
	// 从左到右仅下降一次
	n := len(nums)
	if n <= 1 {
		return 0
	}
	downCnt := 0
	downIdx := -1
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			downCnt++
			downIdx = i + 1
		}
	}
	if downCnt > 1 {
		return -1
	}
	return len(nums) - downIdx
}
