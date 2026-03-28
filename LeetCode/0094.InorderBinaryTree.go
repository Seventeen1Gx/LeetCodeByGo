package LeetCode

import "LeetCodeByGo/utils"

func InorderBinaryTree_1(root *utils.TreeNode) []int {
	ans := make([]int, 0)

	if root == nil {
		return nil
	}

	left := InorderBinaryTree_1(root.Left)
	right := InorderBinaryTree_1(root.Right)

	ans = append(ans, left...)
	ans = append(ans, root.Val)
	ans = append(ans, right...)

	return ans
}

func InorderBinaryTree_2(root *utils.TreeNode) []int {
	ans := make([]int, 0)

	var inorder func(root *utils.TreeNode)
	inorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return ans
}

func InorderBinaryTree_3(root *utils.TreeNode) []int {
	// 非递归解法
	ans := make([]int, 0)

	cur := root
	stack := make([]*utils.TreeNode, 0)
	for len(stack) > 0 || cur != nil {
		for cur != nil { // 一直走到虚拟叶子结点
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 出栈
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, cur.Val)
		// 向右走一步
		cur = cur.Right
	}

	return ans
}
