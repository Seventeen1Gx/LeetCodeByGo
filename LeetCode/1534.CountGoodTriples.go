package LeetCode

import (
	"LeetCodeByGo/utils"
	"slices"
)

func CountGoodTriples_1(arr []int, a int, b int, c int) int {
	var ans int
	var n = len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if utils.Abs(arr[i]-arr[j]) <= a &&
					utils.Abs(arr[j]-arr[k]) <= b &&
					utils.Abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}
	return ans
}

func CountGoodTriples_2(arr []int, a int, b int, c int) int {
	var ans int
	var n = len(arr)
	maxA := slices.Max(arr)

	// 枚举 j,k 可以得到 i 的范围:
	// |Ai-Aj|<=a 即 -a <= Ai-Aj <= a 即 Aj-a <= Ai <= Aj+a
	// |Ai-Ak|<=c 即 -c <= Ai-Ak <= c 即 Ak-c <= Ai <= Ak+c
	// 0 <= Ai <= max(arr)
	// 求交集 [max(0, Aj-a, Ak-c), min(max(arr),Aj+a,Aj+c)] 中每个数在 arr 出现的次数之和

	// 假设有一个数组 cnt ，cnt[i] 表示 arr[i] 在 arr 中出现的次数
	// prefixSum 是 cnt 的前缀和，即 cnt[0:i) 的和
	// prefixSum[0] = 0
	// prefixSum[i+1] = prefixSum[i]+cnt[i]

	// prefixSum[x] 表示下标小于 j 的元素中，小于 x 的元素出现次数
	prefixSum := make([]int, maxA+2)

	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			if utils.Abs(arr[j]-arr[k]) > b {
				continue
			}
			l := max(0, arr[j]-a, arr[k]-c)
			r := min(maxA, arr[j]+a, arr[k]+c)
			if l > r {
				continue
			}
			// 闭区间 [l,r] 中每个元素出现多少次，答案就出现多少次
			// 同时我们要保证 i < j
			// ans += sum[l] + sum[l+1] + ... + sum[r]
			ans += prefixSum[r+1] - prefixSum[l]
		}
		// j 后移，i 的取值范围增加
		// 此时 arr[j] 可作为 Ai 出现一次
		for v := arr[j] + 1; v < maxA+2; v++ {
			prefixSum[v]++
		}
	}

	return ans
}

func CountGoodTriples_3(arr []int, a int, b int, c int) int {
	// 枚举中间试探两边
	ans := 0
	n := len(arr)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return arr[i] - arr[j] })

	for _, j := range idx {
		Aj := arr[j]
		var left, right []int
		for _, i := range idx {
			Ai := arr[i]
			if i < j && utils.Abs(Ai-Aj) <= a {
				left = append(left, Ai)
			}
		}
		for _, k := range idx {
			Ak := arr[k]
			if j < k && utils.Abs(Ak-Aj) <= b {
				right = append(right, Ak)
			}
		}

		// 给定两个有序数组，从两个数组中取数，使得两数差值小于 c
		// 取左边的一个数 x，右边的数应该有 |x-y|<=c 即 -c <= x-y <= c 即 x-c <= y <= x+c
		// 那么又回到 right 中在 [x-c,x+c] 范围内有多少元素的问题
		// right 中小于等于 x+c 的元素个数减去小于 x-c 的元素个数
		var p1, p2 int
		for _, x := range left {
			for p1 < len(right) && right[p1] <= x+c {
				p1++
			}
			for p2 < len(right) && right[p2] < x-c {
				p2++
			}
			ans += p1 - p2
		}
	}
	return ans
}
