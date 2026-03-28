package LeetCode

import "LeetCodeByGo/utils"

func ConstructBTFromPostorderAndInorder(inorder []int, postorder []int) *utils.TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var helper func(s1, s2, n int) *utils.TreeNode
	helper = func(s1, s2, n int) *utils.TreeNode {
		if n == 0 {
			return nil
		}

		rootVal := postorder[s2+n-1]
		root := &utils.TreeNode{Val: rootVal}

		pos := inorderMap[rootVal]
		root.Left = helper(s1, s2, pos-s1)
		root.Right = helper(pos+1, s2+pos-s1, s1+n-pos-1)

		return root
	}

	return helper(0, 0, len(inorder))
}
