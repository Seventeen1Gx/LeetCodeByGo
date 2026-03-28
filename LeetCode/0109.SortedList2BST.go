package LeetCode

import "LeetCodeByGo/utils"

func SortedList2BST(head *utils.ListNode) *utils.TreeNode {
	var helper func(head *utils.ListNode, n int) *utils.TreeNode
	helper = func(head *utils.ListNode, n int) *utils.TreeNode {
		if n == 0 {
			return nil
		}

		cur := head
		for i := 0; i < n/2; i++ {
			cur = cur.Next
		}
		root := &utils.TreeNode{Val: cur.Val}
		root.Left = helper(head, n/2)
		root.Right = helper(cur.Next, n-n/2-1)
		return root
	}

	n := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}

	return helper(head, n)
}
