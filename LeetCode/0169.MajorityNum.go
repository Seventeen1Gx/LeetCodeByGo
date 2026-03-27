package LeetCode

// 绝对众数：即至少出现 n/2+1 次

func majorityElement1(nums []int) int {
	// 暴力法
	var n = len(nums)
	var hashSet = make(map[int]int)

	for _, num := range nums {
		hashSet[num]++
		if hashSet[num] > n/2 {
			return num
		}
	}

	return -1
}

func majorityElement2(nums []int) int {
	// 打擂法：不同元素两两抵消，最终剩的就是绝对众数
	var target int
	var cnt int
	for _, num := range nums {
		if cnt == 0 { // 擂台无人
			target = num
			cnt = 1
		} else if target == num { // 与擂台元素相同
			cnt++
		} else { // 与擂台元素不同
			cnt--
		}
	}

	return target
}
