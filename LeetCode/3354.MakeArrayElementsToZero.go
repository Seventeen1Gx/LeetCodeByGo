package LeetCode

func MakeArrayElements2Zero_1(nums []int) int {
	// 模拟过程
	n := len(nums)
	ans := 0

	sum := 0
	for _, num := range nums {
		sum += num
	}

	f := func(i int, nums []int, goRight bool, sum int) bool {
		for {
			if sum == 0 ||
				i == n-1 && goRight ||
				i == 0 && !goRight {
				break
			}
			// 移动一步
			if goRight {
				i++
			} else {
				i--
			}
			// 移动后碰到非零，则减一后反向
			if nums[i] != 0 {
				nums[i]--
				sum--
				goRight = !goRight
			}
		}
		return sum == 0
	}

	var copyNums []int
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			continue
		}
		// 从 i 位置开始向左走
		copyNums = make([]int, n)
		copy(copyNums, nums)
		if f(i, copyNums, false, sum) {
			ans++
		}
		// 从 i 位置开始向右走
		copy(copyNums, nums)
		if f(i, copyNums, true, sum) {
			ans++
		}
	}

	return ans
}

func MakeArrayElements2Zero_2(nums []int) int {
	// 从位置 i 开始，就是左边减一，右边减一，交替进行
	n := len(nums)
	ans := 0

	sum := 0
	for _, num := range nums {
		sum += num
	}

	var curSum int
	for i := 0; i < n; i++ {
		curSum += nums[i]
		if nums[i] != 0 {
			continue
		}

		if curSum == sum-curSum {
			// 左右相等，开始时向左向右都能成功
			ans++
			ans++
		} else if curSum == sum-curSum+1 {
			// 左边大一，开始时只能向左
			ans++
		} else if curSum+1 == sum-curSum {
			ans++
		}
	}

	return ans
}
