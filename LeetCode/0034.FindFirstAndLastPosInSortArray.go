package LeetCode

func SearchFirstAndLastPosInSortedArray(nums []int, target int) []int {
	n := len(nums)

	ans := []int{-1, -1}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target && (n == 1 || mid == 0 || nums[mid-1] != target) {
			ans[0] = mid
			break
		}
		// nums[mid] != target || mid != 0 && nums[mid-1] == target
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if ans[0] == -1 {
		return ans
	}

	l, r = 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target && (n == 1 || mid == n-1 || nums[mid+1] != target) {
			ans[1] = mid
			break
		}
		// nums[mid] != target || mid != n-1 && nums[mid+1] == target
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			l = mid + 1
		}
	}

	return ans
}
