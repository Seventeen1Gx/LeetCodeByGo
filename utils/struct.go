package utils

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Prev *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
	Depth int
}

func BuildList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	dummyNode := &ListNode{}
	cur := dummyNode
	for _, v := range arr {
		node := &ListNode{Val: v}
		cur.Next = node
		cur = cur.Next
	}
	return dummyNode.Next
}

func BuildBinaryTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// 左子树
		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右子树
		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
