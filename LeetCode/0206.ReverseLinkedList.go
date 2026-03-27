package LeetCode

import "LeetCodeByGo/utils"

func reverseList1(head *utils.ListNode) *utils.ListNode {
	// 三指针迭代
	var prev *utils.ListNode
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func reverseList2(head *utils.ListNode) *utils.ListNode {
	// 递归法
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
