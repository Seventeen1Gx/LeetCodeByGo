package LeetCode

func NumSubarrayProductLessThanK_1(nums []int, k int) int {
	// 有没有前缀积的说法？
	// f[i] = nums[0] * nums[1] * ... * nums[i]
	// f[j] = nums[0] * nums[1] * ... * nums[i] * ... * nums[j]
	// f[j] / f[i] = nums[i+1] * ... * nums[j]

	// 本题中所有元素都是正数
	// 遇到第一个乘积大于 k 的就可以切换起始节点

	ans := 0
	n := len(nums)
	prefixMul := make([]int, n+1)
	prefixMul[0] = 1
	for i := 1; i < n+1; i++ {
		prefixMul[i] = prefixMul[i-1] * nums[i-1]
	}

	// 穷举所有子数组[i:j)
	for i := 0; i < n+1; i++ {
		for j := i + 1; j < n+1; j++ {
			if prefixMul[j]/prefixMul[i] < k {
				ans++
			} else {
				break
			}
		}
	}

	return ans
}

func NumSubarrayProductLessThanK_2(nums []int, k int) int {
	// 滑动窗口(问题具有单调性)
	if k <= 1 {
		return 0
	}

	ans := 0
	prod := 1
	left := 0
	for right, num := range nums { // 枚举右端点
		prod *= num
		for prod >= k {
			// 不满足要求，缩小窗口
			prod /= nums[left]
			left++
		}
		// 此时 [left,right] 是满足条件最长的子数组，所有短于他的也都满足条件
		ans += right - left + 1
	}

	return ans
}
