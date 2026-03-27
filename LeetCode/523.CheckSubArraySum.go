package LeetCode

func CheckSubarraySum_1(nums []int, k int) bool {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	// 穷举所有的子数组
	for i := 0; i < n+1; i++ {
		for j := i + 2; j < n+1; j++ {
			sum := prefixSum[j] - prefixSum[i]
			if sum%k == 0 {
				return true
			}
		}
	}

	return false
}

func CheckSubarraySum_2(nums []int, k int) bool {
	// 哈希法一次遍历
	// prefixSum[j] - prefixSum[i] = n * k
	// prefixSum[j]/k - prefixSum[i]/k = n (整数结果)
	// 即 prefixSum[j] 和 prefixSum[i] 对 k 取余相同
	n := len(nums)
	prefixSum := 0
	hashSet := make(map[int]int)
	hashSet[0] = 0
	for j := 0; j < n+1; j++ {
		if j >= 1 {
			prefixSum += nums[j-1]
		}
		if i, ok := hashSet[prefixSum%k]; ok {
			if j-i >= 2 {
				return true
			}
		} else {
			// 只有不存在才保存，因为我们只要保存这个元素最早出现的位置即可
			hashSet[prefixSum%k] = j
		}
	}

	return false
}
