package LeetCode

import "slices"

func RobotCollisions(positions []int, healths []int, directions string) []int {
	n := len(positions)

	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}

	slices.SortFunc(ids, func(i, j int) int { return positions[i] - positions[j] })

	// 按位置从左到右遍历机器人
	stack := make([]int, 0)
	for _, id := range ids {
		if directions[id] == 'R' {
			// 向右的机器人入栈
			stack = append(stack, id)
		} else {
			// 向左的机器人把栈顶健康度低的机器人移除
			for len(stack) > 0 && healths[stack[len(stack)-1]] < healths[id] {
				healths[stack[len(stack)-1]] = 0
				healths[id]--
				stack = stack[:len(stack)-1]
			}
			if healths[id] == 0 {
				continue
			}
			if len(stack) > 0 {
				if healths[stack[len(stack)-1]] == healths[id] {
					// 栈顶和当前相同，一起销毁
					healths[stack[len(stack)-1]] = 0
					healths[id] = 0
					stack = stack[:len(stack)-1]
				} else {
					// 栈顶比当前的大，当前的移除
					healths[id] = 0
					healths[stack[len(stack)-1]]--
					if healths[stack[len(stack)-1]] == 0 {
						stack = stack[:len(stack)-1]
					}
				}
			}
			// 栈空，则当前向左的机器人剩余健康度保住了
		}
	}

	ans := make([]int, 0)
	for _, health := range healths {
		if health != 0 {
			ans = append(ans, health)
		}
	}

	return ans
}
