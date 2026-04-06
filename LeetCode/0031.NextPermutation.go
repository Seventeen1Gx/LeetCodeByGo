package LeetCode

func NextPermutation(nums []int) {
	// 从后往前找第一个相邻数对 i,i+1 有 nums[i] < nums[i+1]，此时 nums[i+1:] 必然是降序
	// 从后往前在 [i+1:) 范围内找到第一个 nums[j] > nums[i]
	// 交换 nums[i] 和 nums[j]
	// 此时 [i+1:) 依然降序，将其翻转即可
	n := len(nums)
	i := n - 2
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}

	if i != -1 {
		j := n - 1
		for j >= i+1 {
			if nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	l, r := i+1, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
