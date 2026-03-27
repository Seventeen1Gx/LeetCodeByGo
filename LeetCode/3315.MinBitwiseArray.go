package LeetCode

func MinBitwiseArray(nums []int) []int {
	// 位表示上 x + 1 就是将 x 位形式最右边的 1 变成 0，然后将右边第一个 0 变成 1
	// 比如 10111 -> 11000
	// 然后 x | x+1 就是 10111 | 11000 = 11111
	// 也就是 x | x+1 将 x 最右边的 0 变成 1
	// 那么 ans[i] 就是将 nums[i] 最右边的 0 右边的 1 变为 0
	// 如果 nums[i] 是偶数，则无解，同时题目要求 nums 都是质数，故只有 2 无解

	n := len(nums)
	ans := make([]int, n)
	for i, num := range nums {
		if num == 2 {
			ans[i] = -1
		} else {
			// ^10111 = 01000
			// t & -t 即 lowbit = 1000
			// 10111 与 100 异或操作
			t := ^num
			ans[i] = num ^ ((t & -t) >> 1)
		}
	}
	return ans
}
