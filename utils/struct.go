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
}
