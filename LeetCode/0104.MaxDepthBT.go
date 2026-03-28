package LeetCode

import "LeetCodeByGo/utils"

func MaxDepthBT(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(MaxDepthBT(root.Left), MaxDepthBT(root.Right)) + 1
}
