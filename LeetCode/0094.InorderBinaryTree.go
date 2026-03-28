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

func InorderBinaryTree_4(root *utils.TreeNode) []int {
	// 非递归解法 + 无需栈
	ans := make([]int, 0)

	cur := root
	for cur != nil {
		if cur.Left == nil {
			// 访问当前节点
			ans = append(ans, cur.Val)
			// 向右走一步（可能走向右子树，也可能是走向线索）
			cur = cur.Right
		} else {
			// 有左子树，需要往左走，但是怕丢失当前 cur 的位置
			// 故找 cur 的前驱节点
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}
			if pred.Right == nil {
				// 第一次找到前驱，建立线索后，cur 放心地往左走一步
				pred.Right = cur
				cur = cur.Left
			} else {
				// 恢复树结构
				pred.Right = nil

				// 第二次到达前驱，虽然 cur.left != nil 但是左子树已经处理完了
				// 访问当前节点，然后向右走一步
				ans = append(ans, cur.Val)
				cur = cur.Right
			}
		}
	}

	return ans
}
