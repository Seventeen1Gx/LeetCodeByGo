package LeetCode

import "LeetCodeByGo/utils"

func PostorderBinaryTree_1(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)
	left := PostorderBinaryTree_1(root.Left)
	right := PostorderBinaryTree_1(root.Right)
	ans = append(ans, left...)
	ans = append(ans, right...)
	ans = append(ans, root.Val)

	return ans
}

func PostorderBinaryTree_2(root *utils.TreeNode) []int {
	ans := make([]int, 0)

	var postorder func(root *utils.TreeNode)
	postorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		ans = append(ans, root.Val)
	}

	postorder(root)

	return ans
}

func PostorderBinaryTree_3(root *utils.TreeNode) []int {
	ans := make([]int, 0)

	// 双栈法
	stack1 := []*utils.TreeNode{root}
	stack2 := []int{}

	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		if node == nil {
			continue
		}

		// 出栈后进入另一个栈，因为要访问完左右后才能访问它
		stack2 = append(stack2, node.Val)

		// 先左后右
		// 右先出栈1，先进栈2
		// 那么从栈1先出来的就是左
		stack1 = append(stack1, node.Left)
		stack1 = append(stack1, node.Right)
	}

	for i := len(stack2) - 1; i >= 0; i-- {
		ans = append(ans, stack2[i])
	}

	return ans
}

func PostorderBinaryTree_4(root *utils.TreeNode) []int {
	ans := make([]int, 0)

	// 单栈法
	cur := root
	stack := []*utils.TreeNode{}
	var prev *utils.TreeNode
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			// 一直走到最左边
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 此时 cur 为 nil，返回到上一个节点，并不出栈
		cur = stack[len(stack)-1]

		// 如果不存在右子树，或者右子树已经处理完毕，出栈
		if cur.Right == nil || cur.Right == prev {
			ans = append(ans, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur // 记录上一个处理的节点
			cur = nil
		} else {
			// 转到右子树去处理
			cur = cur.Right
		}
	}

	return ans
}
