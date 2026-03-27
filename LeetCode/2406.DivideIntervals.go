package LeetCode

import (
	"container/heap"
	"slices"
	"sort"
)

func DivideIntervals_1(intervals [][]int) int {
	// 将区间按左端点排序
	// 从左到右取不相交的区间为一组，直到取完

	slices.SortFunc(intervals, func(i, j []int) int {
		if i[0] == j[0] {
			return i[1] - j[1]
		}
		return i[0] - j[0]
	})

	isInterSection := func(a, b []int) bool {
		if len(a) == 0 || len(b) == 0 {
			return false
		}
		return !(a[1] < b[0] || a[0] > b[1])
	}

	ans := 0
	n := len(intervals)
	isDeleted := make([]bool, n)
	for {
		last := []int{}
		allDeleted := true
		for i := range intervals {
			if isDeleted[i] {
				continue
			}
			allDeleted = false
			if isInterSection(last, intervals[i]) {
				continue
			}
			isDeleted[i] = true
			last = intervals[i]
		}
		ans++
		if allDeleted {
			break
		}
	}

	return ans
}

func DivideIntervals_2(intervals [][]int) int {
	// 一次遍历
	slices.SortFunc(intervals, func(i, j []int) int {
		if i[0] == j[0] {
			return i[1] - j[1]
		}
		return i[0] - j[0]
	})

	isInterSection := func(a, b []int) bool {
		if len(a) == 0 || len(b) == 0 {
			return false
		}
		return !(a[1] < b[0] || a[0] > b[1])
	}

	group := make([][][]int, 0)
	for _, interval := range intervals {
		newGroup := true
		for i := range group {
			if !isInterSection(group[i][len(group[i])-1], interval) {
				newGroup = false
				group[i] = append(group[i], interval)
				break
			}
		}
		if newGroup {
			group = append(group, [][]int{interval})
		}
	}

	return len(group)
}

func DivideIntervals_3(intervals [][]int) int {
	// 用最小堆优化 group 数组，从而无需遍历它
	// 最小堆存储每组最右边的区间即可，按右边界值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	hp := IntervalRightHeap{}
	for _, interval := range intervals {
		if hp.Len() == 0 || interval[0] <= hp.IntSlice[0] {
			// 当前区间的左边界，比最小的右边界还小，那他肯定和所有区间都相交
			heap.Push(&hp, interval[1])
		} else {
			// 加入最小区间所在的组，即当前区间成功该组最后的
			hp.IntSlice[0] = interval[1]
			heap.Fix(&hp, 0)
		}
	}

	return len(hp.IntSlice)
}

type IntervalRightHeap struct{ sort.IntSlice }

func (hp *IntervalRightHeap) Push(x interface{}) { hp.IntSlice = append(hp.IntSlice, x.(int)) }

func (hp *IntervalRightHeap) Pop() interface{} {
	x := hp.IntSlice[hp.Len()-1]
	hp.IntSlice = hp.IntSlice[:hp.Len()-1]
	return x
}
