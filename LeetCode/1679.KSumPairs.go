package LeetCode

func KSumPairs(nums []int, k int) int {
	// 两数之和为 k 的数对
	// 同时元素不能重复利用
	ans := 0
	hashSet := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		if cnt := hashSet[k-nums[j]]; cnt > 0 {
			ans++
			hashSet[k-nums[j]]--
		} else { // 还没用到的元素才要进表
			hashSet[nums[j]]++
		}
	}
	return ans
}
