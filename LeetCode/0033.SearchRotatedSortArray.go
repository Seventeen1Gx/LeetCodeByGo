package LeetCode

func SearchRotatedSortedArray(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			// nums[mid] 在前半段
			if target > nums[mid] {
				// target > nums[mid] >= nums[0]
				// target 在前半段的 mid 之后
				// 排除 [l,mid]
				l = mid + 1
			} else if target >= nums[0] {
				// target 在前半段的 mid 之前
				// 排除 [mid,r]
				r = mid - 1
			} else {
				// target 在后半段
				l = mid + 1
			}
		} else {
			// nums[mid] 在后半段
			if target >= nums[0] {
				// target 在 前半段
				r = mid - 1
			} else if target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}
