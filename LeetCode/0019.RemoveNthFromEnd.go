package LeetCode

import "LeetCodeByGo/utils"

func RemoveNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	// 双指针
	dummyHead := &utils.ListNode{Next: head}
	p, q := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		p = p.Next // 其中一个指针先走 n 步
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	// 删除 q 之后的节点
	q.Next = q.Next.Next
	return dummyHead.Next
}
