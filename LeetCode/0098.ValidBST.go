package LeetCode

import (
	"LeetCodeByGo/utils"
	"math"
)

func ValidBST_1(root *utils.TreeNode) bool {
	// 中序遍历有序
	flag := true
	var prev = math.MinInt
	var inorder func(root *utils.TreeNode)
	inorder = func(root *utils.TreeNode) {
		if !flag {
			return
		}
		if root == nil {
			return
		}

		inorder(root.Left)
		if root.Val <= prev {
			flag = false
			return
		}
		prev = root.Val
		inorder(root.Right)
	}

	inorder(root)

	return flag
}

func ValidBST_2(root *utils.TreeNode) bool {
	// 二叉搜索树定义，如果左子树存在，子树上所有节点都小于当前节点（右子树同理）
	// 我们用 DFS 遍历所有节点，检查每个节点的值是否在给定的范围内
	var dfs func(root *utils.TreeNode, left, right int) bool
	dfs = func(root *utils.TreeNode, left, right int) bool {
		if root == nil {
			return true
		}
		if root.Val >= right || root.Val <= left {
			return false
		}
		return dfs(root.Left, left, root.Val) && dfs(root.Right, root.Val, right)
	}

	// 为什么开始是无穷大的区间？
	// 节点值的范围由它的祖先限制
	// 而根节点没有父母节点，没人限制它
	// 那我们怎么知道根节点在某个合法范围内
	// 其实子节点比较时，会用到祖先节点，反过来也限制了祖先节点，所以没问题
	return dfs(root, math.MinInt, math.MaxInt)
}

func ValidBST_3(root *utils.TreeNode) bool {
	// 求树的最大值和最小值，再进行比较判断
	var dfs func(root *utils.TreeNode) (int, int)
	dfs = func(root *utils.TreeNode) (int, int) {
		if root == nil {
			return math.MaxInt, math.MinInt
		}
		lMin, lMax := dfs(root.Left)
		rMin, rMax := dfs(root.Right)
		x := root.Val

		if x <= lMax || x >= rMin {
			return math.MinInt, math.MaxInt
		}

		return min(lMin, x), max(rMax, x)
	}

	_, mx := dfs(root)
	return mx != math.MaxInt
}
