package LeetCode

import "LeetCodeByGo/utils"

func PathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return PathSum(root.Left, targetSum-root.Val) || PathSum(root.Right, targetSum-root.Val)
}
