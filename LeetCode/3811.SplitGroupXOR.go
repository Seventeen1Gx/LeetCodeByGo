package LeetCode

func AlternatingXOR_1(nums []int, target1 int, target2 int) int {
	// 每个元素可以加入前一个块或者开启新的块
	cnt := 0
	var backtrace func(i int, cur int, target1, target2 int)
	backtrace = func(i int, cur int, target1, target2 int) {
		if i == len(nums) {
			if cur == target1 {
				cnt++
				cnt %= mod
			}
			return
		}

		// 加入前一个块
		backtrace(i+1, cur^nums[i], target1, target2)

		// 开启新的块
		if cur == target1 {
			backtrace(i+1, nums[i], target2, target1)
		}
	}

	backtrace(0, 0, target1, target2)

	return cnt
}

func AlternatingXOR_2(nums []int, target1 int, target2 int) int {
	// 假设整个数组的异或和为 s ，如果最后一段的异或和是 target，则剩余元素的异或和为 s ^ target
	// 借用前缀和的思路，fi = n0 ^ n1 ^ ... ^ n(i-1)
	// fj ^ fi = ni ^ n(i+1) ^ ... ^ n(j-1)
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		f[i] = f[i-1] ^ nums[i-1]
	}

	// dp[i][0] 表示 nums[:i) 以 target1 开始并以 target1 结尾的有效分组方式
	// dp[i][1] 表示 nums[:i) 以 target1 开始并以 target2 结尾的有效分组方式

	// 状态转移方程
	// dp[i][0] = sum[ dp[j][1] 且 nums[j:i) 的异或和为 target1 ]
	// dp[i][1] = sum[ dp[j][0] 且 nums[j:i) 的异或和为 target2 ]

	// e.g. dp[3][0]
	// 如果 nums[0]^nums[1]^nums[2] == f[0]^f[3] == target1，则 dp[3][0] += dp[0][1]
	// 如果 nums[1]^nums[2] == f[1]^f[3] == target1，则 dp[3][0] += dp[1][1]
	// 如果 nums[2] == f[2]^f[3] == target1，则 dp[3][0] += dp[2][1]

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0], dp[0][1] = 0, 1

	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			if f[j]^f[i] == target1 {
				dp[i][0] += dp[j][1]
			}
			if f[j]^f[i] == target2 {
				dp[i][1] += dp[j][0]
			}
		}
	}

	return dp[n][0] + dp[n][1]
}

func AlternatingXOR_3(nums []int, target1 int, target2 int) int {
	const mod = 1_000_000_007
	f1 := map[int]int{}     // f1[x] 表示异或和为 x 的情况下，满足题目要求，且最后一段是 target1 的方案数
	f2 := map[int]int{0: 1} // f2[x] 表示异或和为 x 的情况下，满足题目要求，且最后一段是 target2 的方案数
	preSum := 0
	for i, x := range nums {
		preSum ^= x                 // 前缀异或和边遍历边计算，也就是 nums[0:i] 的全部元素异或结果
		last1 := f2[preSum^target1] // [0,i] 的最后一段的异或和是 target1 的方案数
		last2 := f1[preSum^target2] // [0,i] 的最后一段的异或和是 target2 的方案数
		if i == len(nums)-1 {
			return (last1 + last2) % mod
		}
		f1[preSum] = (f1[preSum] + last1) % mod
		f2[preSum] = (f2[preSum] + last2) % mod
	}
	panic("unreachable")
}
