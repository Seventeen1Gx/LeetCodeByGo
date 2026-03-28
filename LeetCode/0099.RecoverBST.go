package LeetCode

import (
	"LeetCodeByGo/utils"
	"math"
)

func RecoverBST_1(root *utils.TreeNode) {
	// 中序遍历出现两次或一次下降
	// 如果是一次，说明是相邻两个节点交换
	var node1, node2 *utils.TreeNode
	prev := &utils.TreeNode{Val: math.MinInt}

	var inorder func(root *utils.TreeNode)
	inorder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if root.Val <= prev.Val {
			// 出现下降
			if node1 == nil {
				// 第一次出现
				node1, node2 = prev, root
			} else {
				// 第二次出现
				node2 = root
				return
			}
		}
		prev = root
		inorder(root.Right)
	}

	inorder(root)
	node1.Val, node2.Val = node2.Val, node1.Val
}

func RecoverBST_2(root *utils.TreeNode) {
	// 不借助栈的中序遍历方法
	var node1, node2 *utils.TreeNode
	prev := &utils.TreeNode{Val: math.MinInt}

	cur := root
	for cur != nil {
		if cur.Left == nil {
			// 当前访问节点无左孩子
			// 访问当前节点
			if cur.Val <= prev.Val {
				if node1 == nil {
					node1, node2 = prev, cur
				} else {
					node2 = cur
				}
				prev = cur
				cur = cur.Right // 可能走向右子树，也可能走向线索
			}
		} else {
			// 当前节点有左孩子，需要往左走
			// 但是怕丢失 cur 位置
			// 需要找到它的前驱节点建立线索

			// 找 cur 的前驱节点：左子树的最右节点
			pred := cur.Left
			for pred.Right != nil && pred.Right != cur {
				pred = pred.Right
			}
			if pred.Right == nil {
				// 第一次走到这里，建立线索
				// cur 已经保存位置，可以放心往下走
				pred.Right = cur
				cur = cur.Left
			} else {
				// pred.Right == cur
				// 之前线索化过，走到这里说明左边的子树都处理完了
				// 删除线索
				pred.Right = nil
				//访问当前节点
				if cur.Val <= prev.Val {
					if node1 == nil {
						node1, node2 = prev, cur
					} else {
						node2 = cur
					}
					prev = cur
				}
				cur = cur.Right
			}
		}
	}

	node1.Val, node2.Val = node2.Val, node1.Val
}
