package LeetCode

func TransformerArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i, num := range nums {
		result[i] = nums[((i+num)%n+n)%n]
	}

	return result
}
