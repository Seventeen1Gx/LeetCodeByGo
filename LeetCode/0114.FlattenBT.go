package LeetCode

import "LeetCodeByGo/utils"

func FlattenBT(root *utils.TreeNode) {
	var prev *utils.TreeNode
	var preorder func(root *utils.TreeNode)
	preorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		if prev != nil {
			prev.Right = root
			prev.Left = nil
		}
		prev = root
		right := root.Right
		preorder(root.Left)
		preorder(right)
	}

	preorder(root)
}
