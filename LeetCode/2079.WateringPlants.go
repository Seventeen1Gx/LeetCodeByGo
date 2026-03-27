package LeetCode

func WateringPlants(plants []int, capacity int) int {
	// 模拟操作
	i := 0
	ans := 0
	n := len(plants)
	remaining := capacity
	for i < n {
		ans++
		remaining -= plants[i]
		if i < n-1 && remaining < plants[i+1] {
			// 不足以浇灌下个植物
			ans += 2 * (i + 1)
			remaining = capacity
		}
		i++
	}
	return ans
}
