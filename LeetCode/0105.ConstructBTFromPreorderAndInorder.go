package LeetCode

import "LeetCodeByGo/utils"

func ConstructBTFromPreorderAndInorder(preorder []int, inorder []int) *utils.TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var helper func(s1, s2, n int) *utils.TreeNode
	helper = func(s1, s2, n int) *utils.TreeNode {
		if n == 0 {
			return nil
		}
		// 前序遍历的首元素就是根
		rootVal := preorder[s1]
		root := &utils.TreeNode{Val: rootVal}

		pos := inorderMap[rootVal]

		root.Left = helper(s1+1, s2, pos-s2)
		root.Right = helper(s1+1+pos-s2, pos+1, s2+n-pos-1)
		return root
	}

	return helper(0, 0, len(preorder))
}
