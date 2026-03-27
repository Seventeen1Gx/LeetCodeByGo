package LeetCode

import "LeetCodeByGo/utils"

func TwoSum4_1(root *utils.TreeNode, k int) bool {
	// 哈希法，遍历这颗树的过程中，判断有无满足条件的两数
	hashSet := make(map[int]bool)
	exist := false

	var preOrder func(root *utils.TreeNode)
	preOrder = func(root *utils.TreeNode) {
		if root == nil || exist {
			return
		}
		if hashSet[k-root.Val] {
			exist = true
			return
		}
		hashSet[root.Val] = true
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)

	return exist
}

func TwoSum4_2(root *utils.TreeNode, k int) bool {
	// 中序遍历得到有序数组，然后在有序数组上进行操作
	nums := make([]int, 0)

	var inOrder func(root *utils.TreeNode)
	inOrder = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		nums = append(nums, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)

	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == k {
			return true
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			i++
		}
	}

	return false
}
