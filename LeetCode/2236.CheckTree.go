package LeetCode

import "LeetCodeByGo/utils"

func CheckTree(root *utils.TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
