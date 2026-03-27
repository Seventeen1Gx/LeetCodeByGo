package LeetCode

func TrapRain_1(height []int) int {
	// 每列单独计算雨水量累加
	// 每列的雨水 = min(该列左侧最高的柱子的高度，该列右侧最高的柱子高度) - 当前列的高度
	ans := 0
	n := len(height)

	right := make([]int, n)
	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	left := make([]int, n)
	left[0] = height[0]
	for i := 1; i < n-1; i++ {
		left[i] = max(left[i-1], height[i])
		ans += min(left[i], right[i]) - height[i]
	}

	return ans
}

func TrapRain_2(height []int) int {
	// 相向双指针：边移动边保留当前方向最高的高度，每次计算更矮的指针位置的那列雨水，然后移动一位，最终顶峰相见
	// 为什么：因为矮的位置，它一侧由遍历过程中得到的 max 拦截，另一侧至少由高的柱子拦截，因为就算没遍历到的还有更高的，也不影响结果
	ans := 0
	n := len(height)

	left, right := 0, n-1
	leftMax, righMax := 0, 0
	for left < right {
		heightLeft := height[left]
		heightRight := height[right]
		leftMax = max(leftMax, heightLeft)
		righMax = max(righMax, heightRight)
		if leftMax < righMax {
			ans += leftMax - heightLeft
			left++
		} else {
			ans += righMax - heightRight
			right--
		}
	}

	return ans
}

func TrapRain_3(height []int) int {
	// 单调栈：横着算
	// 从左往右遍历，遍历到的数如果比栈顶小，不能确定它盛水多少，继续遍历
	// 遍历到的数如果比栈顶大（上一个更大的），则栈顶出栈，计算盛水量
	ans := 0
	n := len(height)

	stack := make([]int, 0)
	for right := 0; right < n; right++ {
		rightH := height[right]
		for len(stack) > 0 && height[stack[len(stack)-1]] < rightH {
			curH := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			leftH := height[left]
			area := (min(leftH, rightH) - curH) * (right - left - 1)
			if area > 0 {
				ans += area
			}
		}
		stack = append(stack, right)
	}

	return ans
}
