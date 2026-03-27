package LeetCode

func MaxRectangleInHistogram(heights []int) (ans int) {
	// 柱形图中的最大矩形
	// 枚举每个柱子作为矩形的高，它可以向左或者向右延申到比它矮的柱子
	// 下一个更矮的柱子 -> 单调栈

	n := len(heights)

	// 求左侧更矮的柱子
	// 从左往右遍历
	// 栈存放那些可能是接下来遍历到的柱子的左侧更矮柱子的下标
	// 当前元素比栈顶元素大，则栈顶元素就是更矮的柱子
	// 当前元素比栈顶元素小或等于，栈顶元素不可能是之后元素更矮的柱子，因为当前元素拦着的
	stack := make([]int, 0)
	left := make([]int, n)
	for i, height := range heights {
		m := len(stack)
		for m > 0 && heights[stack[m-1]] >= height {
			stack = stack[:m-1]
			m--
		}

		if m > 0 {
			left[i] = stack[m-1]
		} else {
			left[i] = -1 // 为了让后面计算宽好处理
		}

		stack = append(stack, i)
	}

	// 求右侧更矮的柱子
	// 从左往右遍历（和上面的思路不一样）
	// 栈中存放那些还没确定右侧更矮柱子的下标
	// 当前元素比栈顶元素小，则当前元素就是栈顶元素右侧更矮的柱子
	// 当前元素比栈顶元素大或者等于，则当前元素也不能确定右侧更矮的柱子
	stack = make([]int, 0)
	right := make([]int, n)
	for i, height := range heights {
		m := len(stack)
		for m > 0 && heights[stack[m-1]] > height {
			right[stack[m-1]] = i
			stack = stack[:m-1]
			m--
		}

		stack = append(stack, i)
	}

	// 剩下栈中都是没有右边更矮柱子的，设置成默认值 n
	for _, i := range stack {
		right[i] = n
	}

	// 依次以每个 height 作为矩形的高
	for i, height := range heights {
		width := right[i] - left[i] - 1
		ans = max(ans, height*width)
	}

	return ans
}
