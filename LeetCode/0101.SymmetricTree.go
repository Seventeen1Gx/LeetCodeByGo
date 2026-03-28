package LeetCode

import "LeetCodeByGo/utils"

func SymmetricTree(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	var helper func(p, q *utils.TreeNode) bool
	helper = func(p, q *utils.TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return helper(p.Left, q.Right) && helper(p.Right, q.Left)
	}
	return helper(root.Left, root.Right)
}
