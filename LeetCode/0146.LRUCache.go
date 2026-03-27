package LeetCode

import (
	"LeetCodeByGo/utils"
)

// 双向链表+哈希表
type LRUCache struct {
	size      int
	capacity  int
	dummyHead *utils.ListNode
	keyToNode map[int]*utils.ListNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	dummyHead := &utils.ListNode{}
	dummyHead.Next = dummyHead
	dummyHead.Prev = dummyHead
	return LRUCache{
		size:      0,
		capacity:  capacity,
		dummyHead: dummyHead,
		keyToNode: make(map[int]*utils.ListNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.getNode(key)
	if node == nil {
		return -1
	}
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node := this.getNode(key)
	if node != nil {
		// 结点存在则更新
		node.Val = value
		return
	}

	// 结点不存在则新建
	node = &utils.ListNode{Key: key, Val: value}
	this.keyToNode[key] = node
	this.insertNode(node)

	if this.size == this.capacity {
		tail := this.dummyHead.Prev
		this.removeNode(tail)
		delete(this.keyToNode, tail.Key)
	} else {
		this.size++
	}
}

// 从链表移除结点
func (this *LRUCache) removeNode(node *utils.ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 将结点插入到头部位置
func (this *LRUCache) insertNode(node *utils.ListNode) {
	node.Next = this.dummyHead.Next
	node.Prev = this.dummyHead
	this.dummyHead.Next.Prev = node
	this.dummyHead.Next = node
}

func (this *LRUCache) getNode(key int) *utils.ListNode {
	node := this.keyToNode[key]
	if node == nil {
		return nil
	}
	// 访问结点会刷新结点位置
	this.removeNode(node)
	this.insertNode(node)
	return node
}
