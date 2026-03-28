package LeetCode

import "LeetCodeByGo/utils"

func SameTree(p, q *utils.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right)
}
