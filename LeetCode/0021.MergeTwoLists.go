package LeetCode

import (
	"LeetCodeByGo/utils"
	"math"
)

func MergeTwoLists(list1, list2 *utils.ListNode) *utils.ListNode {
	dummyNode := &utils.ListNode{}
	l1, l2, tail := list1, list2, dummyNode
	for l1 != nil || l2 != nil {
		v1 := math.MaxInt
		if l1 != nil {
			v1 = l1.Val
		}
		v2 := math.MaxInt
		if l2 != nil {
			v2 = l2.Val
		}
		if v1 < v2 {
			tail.Next = l1
			l1 = l1.Next

		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	return dummyNode.Next
}
