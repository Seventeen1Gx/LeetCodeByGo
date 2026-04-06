package LeetCode

func SearchInsertPos(nums []int, target int) int {
	// 首个 >= target 的位置
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target && (mid == 0 || nums[mid-1] != nums[mid]) {
			// 相等情况下的首个
			return mid
		}
		if nums[mid] > target && (mid == 0 || nums[mid-1] < target) {
			// 大于 target 情况下的首个
			return mid
		}

		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
