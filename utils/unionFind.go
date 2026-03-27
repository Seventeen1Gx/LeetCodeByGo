package utils

type UnionFind struct {
	// 用树表示集合，树的根称为代表元
	// 如果 fa[i] = j，当 j > 0 时，说明元素 i 在以 j 为根的集合中，当 j < 0 时，说明元素 i 就是集合的根，-j 代表这棵树的结点高度
	fa []int
}

// 初始化并查集
func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{fa: make([]int, n)}
	for i := range n {
		uf.fa[i] = -1 // 每个元素单独作为一个集合
	}
	return &uf
}

// 寻找元素 x 所在集合的代表元
func (uf *UnionFind) Find(x int) int {
	if uf.fa[x] < 0 {
		// 元素 x 本身就是根
		return x
	}
	// 往上找，并做路径压缩
	y := uf.Find(uf.fa[x])
	uf.fa[x] = y
	return y
}

// 合并元素 x 和元素 y 所在的集合
func (uf *UnionFind) Merge(x, y int) {
	rootx := uf.Find(x)
	rooty := uf.Find(y)
	if rootx != rooty {
		// 结点数小的集合合并到结点数多的集合中
		sizex := -uf.fa[rootx]
		sizey := -uf.fa[rooty]
		if sizex < sizey {
			uf.fa[rootx] = rooty
			uf.fa[rooty] -= sizex
		} else {
			uf.fa[rooty] = rootx
			uf.fa[rootx] -= sizey
		}
	}
}

// 判断是否是相同集合
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// 获得集合数
func (uf *UnionFind) NumOfSets() int {
	var cnt int
	for _, x := range uf.fa {
		if x < 0 {
			cnt++
		}
	}
	return cnt
}

func (uf *UnionFind) Size(x int) int {
	root := uf.Find(x)
	return -uf.fa[root]
}
