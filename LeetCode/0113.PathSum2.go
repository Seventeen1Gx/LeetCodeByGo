package LeetCode

import (
	"LeetCodeByGo/utils"
	"slices"
)

func PathSum2_1(root *utils.TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var dfs func(root *utils.TreeNode, targetSum int, path []int)
	dfs = func(root *utils.TreeNode, targetSum int, path []int) {
		if root == nil {
			return
		}
		targetSum -= root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			ans = append(ans, path)
		}

		leftPath := make([]int, len(path))
		copy(leftPath, path)
		dfs(root.Left, targetSum, leftPath)

		rightPath := make([]int, len(path))
		copy(rightPath, path)
		dfs(root.Right, targetSum, rightPath)
	}

	dfs(root, targetSum, nil)
	return ans
}

func PathSum2_2(root *utils.TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var dfs func(root *utils.TreeNode, targetSum int, path []int)
	dfs = func(root *utils.TreeNode, targetSum int, path []int) {
		if root == nil {
			return
		}
		targetSum -= root.Val
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			ans = append(ans, slices.Clone(path))
		}

		dfs(root.Left, targetSum, path)
		dfs(root.Right, targetSum, path)

		path = path[:len(path)-1] // 恢复现场
	}

	dfs(root, targetSum, nil)
	return ans
}
