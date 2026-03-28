package LeetCode

import (
	"LeetCodeByGo/utils"
)

func ZigzagLevelOrderBT(root *utils.TreeNode) [][]int {
	var ans [][]int

	right := true
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}

			if right {
				level[i] = node.Val
			} else {
				level[size-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
		right = !right
	}

	return ans
}
