package LeetCode

func TwoSum2(numbers []int, target int) []int {
	// 数组已经是有序的了
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}
