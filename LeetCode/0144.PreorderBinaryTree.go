package LeetCode

import "LeetCodeByGo/utils"

func PreorderBinaryTree_1(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{root.Val}
	left := PreorderBinaryTree_1(root.Left)
	right := PreorderBinaryTree_1(root.Right)
	ans = append(ans, left...)
	ans = append(ans, right...)

	return ans
}

func PreorderBinaryTree_2(root *utils.TreeNode) []int {
	ans := make([]int, 0)

	var preorder func(root *utils.TreeNode)
	preorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)

	return ans
}

func PreorderBinaryTree_3(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)

	// 类似队列写法
	stack := []*utils.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}
