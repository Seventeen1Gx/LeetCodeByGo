package LeetCode

import "LeetCodeByGo/utils"

func reverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	if k == 1 {
		return head
	}

	// 添加一个虚拟头结点
	var dummyHead utils.ListNode
	dummyHead.Next = head

	var (
		prevTail = &dummyHead // 上一段的尾结点
		curHead  = head       // 当前段的头结点
	)

	for curHead != nil {
		// 找到当前段的尾结点
		curTail := curHead
		for i := 0; i < k-1; i++ {
			if curTail != nil {
				curTail = curTail.Next
			}
		}

		if curTail == nil {
			// 当前段不足 K 个，直接接上，不做调整
			prevTail.Next = curHead
			break
		}

		// 下一段的头结点
		nextHead := curTail.Next

		// 反转操作
		reverseGroup(curHead, curTail)

		// 接回链表
		prevTail.Next = curTail
		curHead.Next = nextHead

		// 迭代处理下一段
		prevTail = curHead
		curHead = nextHead
	}

	return dummyHead.Next
}

func reverseGroup(head, tail *utils.ListNode) {
	// 反转单链表的部分
	var prev *utils.ListNode
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		if prev == tail {
			break
		}
	}
}
