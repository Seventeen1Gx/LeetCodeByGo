package LeetCode

// 给你一个整数数组 nums 和两个整数：left 及 right
// 找出 nums 中连续、非空且其中最大元素在范围 [left, right] 内的子数组
// 并返回满足条件的子数组的个数

func NumSubarrayBoundedMax_1(nums []int, left, right int) int {
	n := len(nums)

	// 写完之后思考：这里的 dp 和暴力没啥区别，都是枚举了所有子数组

	// dp[i][j] 表示 nums[i:j] 的最大值
	// dp[i][i] = nums[i]
	// dp[i][i+1] = max(dp[i][i], nums[i+1])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], nums[j])
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] >= left && dp[i][j] <= right {
				ans++
			}
		}
	}

	return ans
}

func NumSubarrayBoundedMax_2(nums []int, left, right int) int {
	// f(x) = 最大值 <= x 的子数组数量
	// 答案 = f(right) - f(left-1)
	// f(x) 计算方法：遍历数组遇到 <= x 的数，当前连续段长度 +1 ，累加长度，遇到 > x 则连续段长度置 0
	// 连续段的含义就是从当前遍历到的数，作为子数组的结尾，新增的子数组数
	l1, l2 := 0, 0
	sumA, sumB := 0, 0
	for _, num := range nums {
		if num <= right {
			l1++
			sumA += l1
		} else {
			l1 = 0
		}
		if num <= left-1 {
			l2++
			sumB += l2
		} else {
			l2 = 0
		}
	}
	return sumA - sumB
}
