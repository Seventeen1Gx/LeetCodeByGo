package utils

import (
	"errors"
)

// 最小堆：一颗二叉树，每个节点满足，如果存在孩子节点，孩子节点的值都大于当前节点值
type MyMinHeap struct {
	elements []int
}

func NewMyHeap() *MyMinHeap {
	return &MyMinHeap{
		elements: make([]int, 0),
	}
}

func (hp *MyMinHeap) Len() int { return len(hp.elements) }

func (hp *MyMinHeap) Push(x int) {
	// 新元素加入到堆结尾
	// 然后将其向上调整，维护堆的性质
	hp.elements = append(hp.elements, x)
	hp.up(len(hp.elements) - 1)
}

func (hp *MyMinHeap) Pop() (int, error) {
	// 获取根节点
	// 根节点和最后一个节点交换，删除结尾节点
	// 将根节点向下调整，维护堆的性质
	n := len(hp.elements)
	if n == 0 {
		return 0, errors.New("heap is empty")
	}
	x := hp.elements[0]
	hp.elements[0] = hp.elements[n-1]
	hp.elements = hp.elements[:n-1]
	hp.down(0)
	return x, nil
}

func (hp *MyMinHeap) up(idx int) {
	for {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 || hp.elements[parentIdx] <= hp.elements[idx] {
			break
		}
		hp.elements[parentIdx], hp.elements[idx] = hp.elements[idx], hp.elements[parentIdx]
		idx = parentIdx
	}
}

func (hp *MyMinHeap) down(idx int) {
	n := len(hp.elements)
	for {
		leftChildIdx := 2*idx + 1
		rightChildIdx := leftChildIdx + 1
		if leftChildIdx >= n {
			// 不存在孩子节点
			break
		}

		smallestIdx := leftChildIdx // 先假设左孩子小
		if rightChildIdx < n && hp.elements[rightChildIdx] < hp.elements[smallestIdx] {
			// 右孩子存在且更小
			smallestIdx = rightChildIdx
		}

		if hp.elements[idx] < hp.elements[smallestIdx] {
			break
		}
		hp.elements[idx], hp.elements[smallestIdx] = hp.elements[smallestIdx], hp.elements[idx]
		idx = smallestIdx
	}
}

type MyMaxHeap []int

func (hp MyMaxHeap) Len() int { return len(hp) }

func (hp MyMaxHeap) Less(i, j int) bool { return hp[i] > hp[j] }

func (hp MyMaxHeap) Swap(i, j int) { hp[i], hp[j] = hp[j], hp[i] }

func (hp *MyMaxHeap) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}

func (hp *MyMaxHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	x := old[n-1]
	*hp = old[:n-1]
	return x
}
