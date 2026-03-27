package LeetCode

func ProductExceptSelf(nums []int) []int {
	// pre[i] = nums[0] * nums[1] * ... * nums[i-1]
	// suf[i] = nums[i] * nums[i+1] * ... * nums[L-1]
	// 为什么 suf 不从 i+1 开始乘，因为如果是那样 suf[0] = nums[1] * ...* nums[L-1]
	// 所有元素的乘积变成了 suf[-1]
	// 强行要这样，就只能 suf[0] 表示 suf[-1] 这种，会导致很奇怪

	/*
		n := len(nums)
		ans := make([]int, n)
		pre, suf := make([]int, n+1), make([]int, n+1)
		pre[0], suf[n] = 1, 1

		for i := 1; i < n+1; i++ {
			pre[i] = nums[i-1] * pre[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			suf[i] = nums[i] * suf[i+1]
		}

		for i := range nums {
			ans[i] = pre[i] * suf[i+1]
		}
	*/

	// 其实我们不需要所有元素乘在一起的结果
	// pre[i] = nums[0] * ... * nums[i-1]
	// suf[i] = nums[i+1] * ... * nums[L-1]

	n := len(nums)
	ans := make([]int, n)
	pre, suf := make([]int, n), make([]int, n)
	pre[0], suf[n-1] = 1, 1

	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	for i := range nums {
		ans[i] = pre[i] * suf[i]
	}

	return ans
}
