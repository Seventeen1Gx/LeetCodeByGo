package LeetCode

import "LeetCodeByGo/utils"

func PopulatingNextRightPointer(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var prev *utils.TreeNode
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if prev != nil {
				prev.Next = node
			}
			prev = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
