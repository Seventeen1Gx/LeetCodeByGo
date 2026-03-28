package LeetCode

import (
	"LeetCodeByGo/utils"
)

func UniqueBinarySearchTree2(n int) []*utils.TreeNode {
	var generateBST func(left, right int) []*utils.TreeNode
	generateBST = func(left, right int) []*utils.TreeNode {
		if left > right {
			return []*utils.TreeNode{nil}
		}

		ans := make([]*utils.TreeNode, 0)
		for i := left; i <= right; i++ {
			// 以 i 作为根
			leftSubTrees := generateBST(left, i-1)
			rightSubTrees := generateBST(i+1, right)
			for _, leftSubTree := range leftSubTrees {
				for _, rightSubTree := range rightSubTrees {
					root := &utils.TreeNode{Val: i}
					root.Left = leftSubTree
					root.Right = rightSubTree
					ans = append(ans, root)
				}
			}
		}
		return ans
	}

	return generateBST(1, n)
}
