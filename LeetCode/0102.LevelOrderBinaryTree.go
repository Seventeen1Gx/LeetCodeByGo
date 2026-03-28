package LeetCode

import "LeetCodeByGo/utils"

func LevelOrder(root *utils.TreeNode) [][]int {
	var ans [][]int
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			level = append(level, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		ans = append(ans, level)
	}

	return ans
}
