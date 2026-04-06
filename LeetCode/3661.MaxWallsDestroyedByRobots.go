package LeetCode

import "sort"

// 与机器人重合的墙壁，只能由重合的机器人能打到，其他人打不到
// 机器人 i 向左攻击范围：
// - 左边存在其他机器人：[max(xi-di, x_(i-1)+1, xi]
// - 左边不存在其他机器人：[xi-di, xi]
// 机器人 i 向右攻击范围：[xi, min(xi+di, x_(i+1)-1)] / [xi, xi+di]
// 确定范围后，通过二分法查找能够摧毁的墙壁数
// left[i] 和 right[i] 表示第 i 个机器人向左/向右能摧毁的墙壁数
// num[i] 表示第 i 个机器人和第 i-1 个机器人之间的墙壁数

// dp[i][0] 表示第 i 个机器人向左开一枪后，前 i 个机器人总共摧毁墙壁的最大数量
// dp[i][1] 表示第 i 个机器人向右开一枪后，前 i 个机器人总共摧毁墙壁的最大数量（视右边无机器人）
//
// 初始条件：
// dp[0][0] = left[0], dp[0][1] = right[0]
// 递推公式：
// dp[i][0] = max(dp[i-1][0]+left[i], dp[i-1][1]-right[i-1]+min(right[i-1]+left[i], num[i]))
// dp[i][1] = max(dp[i-1][0]+right[i], dp[i-1][1]+right[i])
// 解释：
// dp[i-1][0] 第 i-1 机器人向左开枪的情况下，前 i-1 个机器人能摧毁的最大墙壁数量，加上第 i 机器人向左开枪能够摧毁的最大墙壁数量 left[i]
// dp[i-1][1] 第 i-1 机器人向右开枪的情况下，前 i-1 个机器人能摧毁的最大墙壁数量，扣除第 i-1 机器人向右开枪能够摧毁的最大墙壁数量 right[i-1]
// 再加上 right[i-1] + left[i] 以及 num[i] 的最小值，这里是 i-1 和 i 相向开一枪，看中间的墙壁数量是否够，不够就取墙壁数
//
// 最终答案 max(dp[n-1][0], dp[n-1][1])

func MaxWallsDestroyedByRobots(robots []int, distance []int, walls []int) int {
	n := len(robots)

	// 机器人位置对应它的射击距离
	robots2Dist := make(map[int]int)
	for i, robot := range robots {
		robots2Dist[robot] = distance[i]
	}

	// 位置排序
	sort.Ints(robots)
	sort.Ints(walls)

	// 准备
	left, right, num := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		// 找到大于等于 robots[i] 向左射击的最远位置的墙壁位置
		var leftPos int
		if i >= 1 {
			leftBound := max(robots[i]-robots2Dist[robots[i]], robots[i-1]+1)
			leftPos = sort.SearchInts(walls, leftBound)
		} else {
			leftPos = sort.SearchInts(walls, robots[i]-robots2Dist[robots[i]])
		}
		pos1 := sort.SearchInts(walls, robots[i]+1) // robots[i] 右边（不包含它自己）最近的墙壁下标
		left[i] = pos1 - leftPos

		// 找到大于 robots[i] 向右射击的最远位置的墙壁位置
		var rightPos int
		if i < n-1 {
			rightBound := min(robots[i]+robots2Dist[robots[i]], robots[i+1]-1)
			rightPos = sort.SearchInts(walls, rightBound+1)
		} else {
			rightPos = sort.SearchInts(walls, robots[i]+robots2Dist[robots[i]]+1)
		}
		pos2 := sort.SearchInts(walls, robots[i]) // robots[i] 右边（包含它自己）最近的墙壁下标
		right[i] = rightPos - pos2

		if i == 0 {
			continue
		}
		pos3 := sort.SearchInts(walls, robots[i-1])
		num[i] = pos1 - pos3
	}

	// 使用滚动数组
	// preLeft = dp[i-1][0], preRight = dp[i-1][1]
	// curLeft = dp[i][0], curRight = dp[i][1]
	preLeft, preRight := left[0], right[0]
	for i := 1; i < n; i++ {
		curLeft := max(preLeft+left[i], preRight-right[i-1]+min(right[i-1]+left[i], num[i]))
		curRight := max(preLeft, preRight) + right[i]
		preLeft = curLeft
		preRight = curRight
	}

	return max(preLeft, preRight)
}
