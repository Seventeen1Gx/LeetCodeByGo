package LeetCode

func UniqueBinarySearchTree_1(n int) int {
	var numsOfBST func(left, right int) int
	numsOfBST = func(left, right int) int {
		if left >= right {
			return 1
		}
		ans := 0
		for i := left; i <= right; i++ {
			ans += numsOfBST(left, i-1) * numsOfBST(i+1, right)
		}
		return ans
	}
	return numsOfBST(1, n)
}

func UniqueBinarySearchTree_2(n int) int {
	// 动态规划
	// G(n) 表示 n 个不同节点组成的不同二叉搜索树数
	// F(i, n) 表示 1-n 个不同节点，以 i 为根的不同二叉搜索树数
	// G(n) = sum(F(i,n))
	// F(i,n) = G(i-1) * G(n-i) 即 G(n) = sum(G(i-1)*G(n-i)), i∈[1,n]
	// G(0) = G(1) = 1

	if n <= 1 {
		return 1
	}

	g := make([]int, n+1)
	g[0], g[1] = 1, 1

	for i := 2; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}

	return g[n]
}
