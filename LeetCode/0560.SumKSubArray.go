package LeetCode

func SubArraySum_1(nums []int, k int) int {
	// 前缀和
	// f[0] = 0
	// f[i] = f[i-1] + nums[i]
	ans := 0
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		f[i] = f[i-1] + nums[i-1]
	}

	// 子数组[i:j)的和 = f[j] - f[i]

	// 暴力穷举
	for i := 0; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			if f[j]-f[i] == k {
				ans++
			}
		}
	}

	return ans
}

func SubArraySum_2(nums []int, k int) int {
	ans := 0
	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		f[i] = f[i-1] + nums[i-1]
	}

	// 子数组[i:j)的和 = f[j] - f[i]

	// 哈希法
	hashSet := make(map[int]int)
	for j := 0; j < n+1; j++ {
		if cnt, ok := hashSet[f[j]-k]; ok {
			ans += cnt
		}
		hashSet[f[j]]++
	}

	return ans
}

func SubArraySum_3(nums []int, k int) int {
	ans := 0
	n := len(nums)
	f := make([]int, n+1)

	// 子数组[i:j)的和 = f[j] - f[i]

	// 哈希法，一次遍历
	hashSet := make(map[int]int)
	for j := 0; j < n+1; j++ {
		if j >= 1 {
			f[j] = f[j-1] + nums[j-1]
		}
		if cnt, ok := hashSet[f[j]-k]; ok {
			ans += cnt
		}
		hashSet[f[j]]++
	}

	return ans
}

func SubArraySum_4(nums []int, k int) int {
	ans := 0
	n := len(nums)
	s := 0

	// 子数组[i:j)的和 = f[j] - f[i]

	// 哈希法，一次遍历，甚至不需要 f 数组
	hashSet := make(map[int]int)
	for j := 0; j < n+1; j++ {
		if j >= 1 {
			s += nums[j-1]
		}
		if cnt, ok := hashSet[s-k]; ok {
			ans += cnt
		}
		hashSet[s]++
	}

	return ans
}
