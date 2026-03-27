package LeetCode

func FourDivisors(nums []int) int {
	// 预处理每个因子
	mx := 10001
	divisorNum := make([]int, mx) // divisorNum[num] = x 表示数字 num 有 x 个因子
	divisorSum := make([]int, mx) // divisorSum[num] = x 表示数字 num 的因子之和为 x

	for i := 1; i < mx; i++ {
		for num := i; num < mx; num += i {
			divisorNum[num]++ // i i*2 i*3 i*4 ... 都有因子 i
			divisorSum[num] += i
		}
	}

	var ans int
	for _, num := range nums {
		if divisorNum[num] == 4 {
			ans += divisorSum[num]
		}
	}

	return ans
}
