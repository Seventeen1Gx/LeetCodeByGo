package LeetCode

// 给定一个 n 叉树的根节点 root，返回其节点值的[前序遍历]。

// n 叉树在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

type NAryTreeNode struct {
	Val      int
	Children []*NAryTreeNode
}

func preorder(root *NAryTreeNode) []int {
	if root == nil {
		return nil
	}

	var ret = []int{root.Val}
	for _, child := range root.Children {
		childRet := preorder(child)
		ret = append(ret, childRet...)
	}
	return ret
}
