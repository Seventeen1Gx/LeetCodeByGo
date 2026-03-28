package LeetCode

import "LeetCodeByGo/utils"

func SortedArray2BST(nums []int) *utils.TreeNode {
	var helper func(i, j int) *utils.TreeNode
	helper = func(i, j int) *utils.TreeNode {
		if i > j {
			return nil
		}
		rootIdx := (i + j) / 2
		root := &utils.TreeNode{Val: nums[rootIdx]}
		root.Left = helper(i, rootIdx-1)
		root.Right = helper(rootIdx+1, j)
		return root
	}

	return helper(0, len(nums)-1)
}
