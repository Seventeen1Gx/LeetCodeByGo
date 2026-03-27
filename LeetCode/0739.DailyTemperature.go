package LeetCode

func DailyTemperature_1(temperatures []int) []int {
	// 以 1 4 3 5 5 2 3 6 为例
	// 从左往右遍历
	// 栈存储那些还未明确更高温度的元素
	// 遍历到 1 我们不知道它的下一个更大元素是什么故入栈
	// 遍历到 4 此时它比栈顶 1 大，那它就是栈顶 1 的下一个更大元素
	// 处理 1 的结果，1 出栈，此时我们不知道 4 的下一个最大元素，故 4 入栈
	// 遍历到 3 此时它比栈顶 4 小，也无法确定，故入栈
	// 遍历到 5 此时它比栈顶 3 小，它是 3 的下一个更大元素
	// 处理 3 的结果，3 出栈，此时 5 依旧比栈顶元素 4 大
	// 处理 4 的结果，4 出栈，此时我们不知道 5 的下一个最大元素，故 5 入栈
	// ...
	// 可以看到，我们栈内的元素始终是递减的

	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i, t := range temperatures {
		m := len(stack)
		for m > 0 && temperatures[stack[m-1]] < t {
			ans[stack[m-1]] = i - stack[m-1]
			stack = stack[:m-1]
			m--
		}
		stack = append(stack, i)
	}

	return ans
}

func DailyTemperature_2(temperatures []int) []int {
	// 以 1 4 3 5 5 2 3 6 为例
	// 从右往左遍历
	// 栈存储那些可能是之前元素的下一个更高温度的元素
	// 遍历到 6 它可能是之前元素的下一个更高温度，故入栈
	// 遍历到 3 此时它比栈顶 6 小，故 3 的下一个更高温度是 6
	// 但 3 也可能是之前元素的下一个更高温度，故入栈
	// 遍历到 2 此时它比栈顶 3 小，故 2 的下一个更高温度是 3
	// 但 2 也可能是之前元素的下一个更高温度，故入栈
	// 遍历到 5 此时它比栈顶 2 大，2 不可能是之前元素的更高温度了，2 出栈，同理 3 也出栈
	// 此时 5 比栈顶 6 小，故 5 的下一个更高温度是 6
	// ...
	// 可以看到，我们栈内的元素始终是递减的

	n := len(temperatures)
	ans := make([]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		m := len(stack)
		for m > 0 && temperatures[stack[m-1]] <= temperatures[i] {
			stack = stack[:m-1]
			m--
		}

		if m > 0 {
			ans[i] = stack[m-1] - i
		}

		stack = append(stack, i)
	}

	return ans
}
