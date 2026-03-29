package LeetCode

import "LeetCodeByGo/utils"

// 上一题是完美二叉树，而这题是普通二叉树
// 但实际没差别，上一题层序遍历的代码还适用
// 这里记录更多解法

func PopulatingNextRightPointer2_1(root *utils.TreeNode) *utils.TreeNode {
	// 记录每层需要处理的尾节点
	var tails []*utils.TreeNode

	// 先序遍历每个节点，depth 是节点所处的深度
	var preorder func(root *utils.TreeNode, depth int)
	preorder = func(root *utils.TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth <= len(tails) {
			tails[depth-1].Next = root
			tails[depth-1] = root
		} else {
			tails = append(tails, root)
		}
		preorder(root.Left, depth+1)
		preorder(root.Right, depth+1)
	}

	preorder(root, 1)
	return root
}

func PopulatingNextRightPointer2_2(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// 层序遍历
	// 记录当前层和下一层
	cur, nxt := []*utils.TreeNode{root}, []*utils.TreeNode{}
	for len(cur) > 0 {
		for i := 0; i < len(cur); i++ {
			if i != 0 {
				cur[i-1].Next = cur[i]
			}
			if cur[i].Left != nil {
				nxt = append(nxt, cur[i].Left)
			}
			if cur[i].Right != nil {
				nxt = append(nxt, cur[i].Right)
			}
		}
		cur = nxt
		nxt = nil
	}
	return root
}

func PopulatingNextRightPointer2_3(root *utils.TreeNode) *utils.TreeNode {
	// 每层按链表去遍历
	dummyNode := &utils.TreeNode{}

	cur := root
	for cur != nil {
		dummyNode.Next = nil // 上一层的链表断开
		nxt := dummyNode     // 下一层的尾部节点开始
		for cur != nil {     // 同层遍历
			if cur.Left != nil {
				nxt.Next = cur.Left
				nxt = nxt.Next
			}
			if cur.Right != nil {
				nxt.Next = cur.Right
				nxt = nxt.Next
			}
			cur = cur.Next
		}
		cur = dummyNode.Next // 从下一层起点开始
	}

	return root
}
