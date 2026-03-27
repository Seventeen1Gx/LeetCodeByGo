package LeetCode

func MaxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	uniq := make(map[int]int, 0)
	windowSum := int64(0)
	maxWindowSum := int64(0)
	for right := 0; right < n; right++ {
		num := nums[right]
		windowSum += int64(num)
		uniq[num]++
		left := right - k + 1
		if left < 0 {
			continue
		}

		if len(uniq) == k {
			maxWindowSum = max(maxWindowSum, windowSum)
		}

		num = nums[left]
		windowSum -= int64(num)
		uniq[num]--
		if uniq[num] == 0 {
			delete(uniq, num)
		}
	}

	return maxWindowSum
}
