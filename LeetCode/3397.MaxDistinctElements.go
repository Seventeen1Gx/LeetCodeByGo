package LeetCode

import (
	"LeetCodeByGo/utils"
	"math"
	"slices"
)

func MaxDistinctElements(nums []int, k int) (ans int) {
	// 贪心，想象整数在数轴上
	// 对于最小的元素，尽可能往左移动
	// 最小的 a 变成 a-k 令其为 a'
	// 次小的 b 变成 max(b-k,a'+1)，同时它又不能超过 b+k
	// 故最终取 min(max(b-k,a'+1),b+k)
	// 比如 a=2 b=2 c=2 d=2 k=1
	// a'=1 b'=2 c'=3 d'=3
	slices.Sort(nums)
	preNum := math.MinInt
	for _, num := range nums {
		newNum := utils.Min(utils.Max(num-k, preNum+1), num+k)
		if newNum > preNum {
			ans++
			preNum = newNum
		}
	}
	return ans
}
