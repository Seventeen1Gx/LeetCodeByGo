package LeetCode

// 给你两个[非空]的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储[一位]数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0      // 初始进位
	var dummyHead ListNode // 虚头节点
	var tail = &dummyHead  // 尾节点指针

	for l1 != nil || l2 != nil {
		// 对应位相加
		var a, b, c int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		c = a + b + carry
		// 得到新节点
		newNode := &ListNode{
			Val:  c % 10,
		}
		carry = c / 10
		// 链接
		tail.Next = newNode
		tail = tail.Next
	}
	if carry != 0 {
		newNode := &ListNode{
			Val:  carry,
		}
		tail.Next = newNode
	}

	return dummyHead.Next
}
