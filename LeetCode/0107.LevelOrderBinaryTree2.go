package LeetCode

import (
	"LeetCodeByGo/utils"
	"slices"
)

func LevelOrderBinaryTree2(root *utils.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var level = make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	slices.Reverse(ans)
	return ans
}
