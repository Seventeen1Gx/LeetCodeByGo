package LeetCode

import (
	"LeetCodeByGo/utils"
	"math"
)

func MinDepthBT_1(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	// 递归计算最小深度
	var helper func(root *utils.TreeNode) int
	helper = func(root *utils.TreeNode) int {
		if root == nil {
			return math.MaxInt
		}
		if root.Left == nil && root.Right == nil {
			return 1
		}
		left := helper(root.Left)
		right := helper(root.Right)
		return min(left, right) + 1
	}
	return helper(root)
}

func MinDepthBT_2(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	// 利用遍历，遍历到叶子节点时判断是否最小深度

	ans := math.MaxInt
	var preorder func(root *utils.TreeNode, curDepth int)
	preorder = func(root *utils.TreeNode, curDepth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			ans = min(ans, curDepth)
		}
		preorder(root.Left, curDepth+1)
		preorder(root.Right, curDepth+1)
	}

	preorder(root, 1)
	return ans
}

func MinDepthBT_3(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MaxInt
	root.Depth = 1
	stack := []*utils.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left == nil && node.Right == nil {
			ans = min(ans, node.Depth)
		}
		if node.Right != nil {
			node.Right.Depth = node.Depth + 1
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			node.Left.Depth = node.Depth + 1
			stack = append(stack, node.Left)
		}
	}

	return ans
}

func MinDepthBT_4(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MaxInt
	var inorder func(root *utils.TreeNode, depth int)
	inorder = func(root *utils.TreeNode, depth int) {
		if root.Left != nil {
			inorder(root.Left, depth+1)
		}
		if root.Left == nil && root.Right == nil {
			ans = min(ans, depth)
			return
		}
		if root.Right != nil {
			inorder(root.Right, depth+1)
		}
	}

	inorder(root, 1)
	return ans
}

func MinDepthBT_5(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MaxInt
	cur := root
	cur.Depth = 1
	stack := make([]*utils.TreeNode, 0)
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			d := cur.Depth
			cur = cur.Left
			if cur != nil {
				cur.Depth = d + 1
			}
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.Left == nil && cur.Right == nil {
			ans = min(ans, cur.Depth)
		}

		d := cur.Depth
		cur = cur.Right
		if cur != nil {
			cur.Depth = d + 1
		}
	}

	return ans
}

func MinDepthBT_6(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	ans := math.MaxInt
	var postorder func(root *utils.TreeNode, depth int)
	postorder = func(root *utils.TreeNode, depth int) {
		if root.Left != nil {
			postorder(root.Left, depth+1)
		}
		if root.Right != nil {
			postorder(root.Right, depth+1)
		}
		if root.Left == nil && root.Right == nil {
			ans = min(ans, depth)
		}
	}
	postorder(root, 1)

	return ans
}

func MinDepthBT_7(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	// 后序遍历双栈法
	root.Depth = 1
	stack1, stack2 := []*utils.TreeNode{root}, []*utils.TreeNode{}
	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		stack2 = append(stack2, node)
		if node.Right != nil {
			node.Right.Depth = node.Depth + 1
			stack1 = append(stack1, node.Right)
		}
		if node.Left != nil {
			node.Left.Depth = node.Depth + 1
			stack1 = append(stack1, node.Left)
		}
	}

	ans := math.MaxInt
	for i := len(stack2) - 1; i >= 0; i-- {
		node := stack2[i]
		if node.Left == nil && node.Right == nil {
			ans = min(ans, node.Depth)
		}
	}

	return ans
}

func MinDepthBT_8(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	// 后续遍历单栈法
	ans := math.MaxInt
	root.Depth = 1
	stack := []*utils.TreeNode{root}
	var prev *utils.TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			d := cur.Depth
			cur = cur.Left
			if cur != nil {
				cur.Depth = d + 1
			}
		}
		// 回到上一个节点，它的左子树为空
		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == prev {
			// 右子树为空或者右子树已经处理过了
			// 访问该节点
			stack = stack[:len(stack)-1]
			if cur.Left == nil && cur.Right == nil {
				ans = min(ans, cur.Depth)
			}
			prev = cur
			cur = nil // 通过栈顶返回上一个元素
		} else {
			d := cur.Depth
			cur = cur.Right
			if cur != nil {
				cur.Depth = d + 1
			}
		}
	}

	return ans
}
