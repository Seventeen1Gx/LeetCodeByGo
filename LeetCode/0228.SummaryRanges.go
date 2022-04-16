package LeetCode

import "strconv"

// 给定一个无重复元素的有序整数数组 nums 。
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
// 列表中的每个区间范围 [a,b] 应该按如下格式输出：
// "a->b" ，如果 a != b
// "a" ，如果 a == b

// 提示：
// 0 <= nums.length <= 20
// -231 <= nums[i] <= 231 - 1
// nums 中的所有值都 互不相同
// nums 按升序排列

func summaryRanges(nums []int) (ret []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {		// 时刻保证下标在有效范围中
		}
		// [left, i-1] 是一个连续区间
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ret = append(ret, s)
	}
	return ret
}
