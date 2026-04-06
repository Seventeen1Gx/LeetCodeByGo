package LeetCode

import (
	"LeetCodeByGo/utils"
	"container/heap"
)

func MergeKList(lists []*utils.ListNode) *utils.ListNode {
	dummyNode := &utils.ListNode{}
	tail := dummyNode

	hp := &ListNodeHeap{items: make([]*utils.ListNode, 0)}
	for _, l := range lists {
		if l != nil {
			hp.items = append(hp.items, l)
		}
	}
	heap.Init(hp)

	for hp.Len() > 0 {
		node := heap.Pop(hp).(*utils.ListNode)
		tail.Next = node
		tail = tail.Next
		if node.Next != nil {
			heap.Push(hp, node.Next)
		}
	}

	return dummyNode.Next
}

type ListNodeHeap struct {
	items []*utils.ListNode
}

func (hp *ListNodeHeap) Len() int           { return len(hp.items) }
func (hp *ListNodeHeap) Swap(i, j int)      { hp.items[i], hp.items[j] = hp.items[j], hp.items[i] }
func (hp *ListNodeHeap) Less(i, j int) bool { return hp.items[i].Val < hp.items[j].Val }
func (hp *ListNodeHeap) Push(x interface{}) {
	// 末尾添加元素
	hp.items = append(hp.items, x.(*utils.ListNode))
}
func (hp *ListNodeHeap) Pop() (x interface{}) {
	// 删除末尾元素
	n := len(hp.items)
	x = hp.items[n-1]
	hp.items = hp.items[:n-1]
	return x
}
