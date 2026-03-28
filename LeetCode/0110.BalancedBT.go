package LeetCode

import (
	"LeetCodeByGo/utils"
)

func BalancedBT_1(root *utils.TreeNode) bool {
	var depth func(root *utils.TreeNode) int
	depth = func(root *utils.TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := depth(root.Left)
		rightDepth := depth(root.Right)
		if leftDepth == -1 || rightDepth == -1 {
			// 已经不平衡了，剪枝
			return -1
		}
		if leftDepth > rightDepth && leftDepth-rightDepth > 1 {
			return -1
		}
		if leftDepth < rightDepth && rightDepth-leftDepth > 1 {
			return -1
		}
		return max(leftDepth, rightDepth) + 1
	}

	return depth(root) == -1
}

func BalancedBT_2(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	var depth func(root *utils.TreeNode) int
	depth = func(root *utils.TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := depth(root.Left)
		rightDepth := depth(root.Right)
		return max(leftDepth, rightDepth) + 1
	}

	return utils.Abs(depth(root.Left)-depth(root.Right)) <= 1 && BalancedBT_2(root.Left) && BalancedBT_2(root.Right)
}
