package LeetCode

func AsteroidCollision(asteroids []int) []int {
	ans := make([]int, 0)
	stack := make([]int, 0)
	for _, a := range asteroids {
		if a > 0 {
			// 向右移动的小行星，入栈
			stack = append(stack, a)
		} else {
			// 向左移动的小行星，与相近的向右移动小行星发生碰撞
			// 小于它的行星都发生爆炸
			for len(stack) > 0 && stack[len(stack)-1] < -a {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				if stack[len(stack)-1] == -a {
					// 栈顶和当前行星同时爆炸
					stack = stack[:len(stack)-1]
				}
				// 栈顶大于当前行星，当前行星爆炸
			} else {
				// 栈空，留下当前小行星
				ans = append(ans, a)
			}
		}
	}
	// 最后栈中的都是剩下的
	for _, a := range stack {
		ans = append(ans, a)
	}
	return ans
}
