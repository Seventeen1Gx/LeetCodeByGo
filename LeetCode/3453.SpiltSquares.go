package LeetCode

import (
	"math/bits"
	"sort"
)

func SpiltSquares_1(squares [][]int) float64 {
	totalArea := 0
	maxY := -1
	for _, s := range squares {
		l := s[2]
		totalArea += l * l
		maxY = max(maxY, s[1]+l)
	}

	// 计算在 y 之下的正方形面积
	check := func(y float64) bool {
		area := 0.
		for _, s := range squares {
			l := float64(s[2])
			yi := float64(s[1])
			if yi < y {
				area += l * min(y-yi, l)
			}
		}
		return area*2 >= float64(totalArea)
	}

	// 在 (0,maxY) 进行浮点二分
	// 每循环一次搜索范围减半，循环 k 次后，要满足 L/2^k <= 10^-5
	left, right := 0., float64(maxY)
	for range bits.Len(uint(maxY * 1e5)) {
		y := (left + right) / 2
		if check(y) {
			right = y
		} else {
			left = y
		}
	}

	return (left + right) / 2
}

func SpiltSquares_2(squares [][]int) float64 {
	// 整数二分，就是先乘 100000 计算，最后除 100000 得到结果
	totalArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totalArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	const m = 100_000

	// 在 [0,maxY*m] 之间二分搜索
	multiY := sort.Search(maxY*m, func(multiY int) bool {
		area := 0
		for _, sq := range squares {
			yi, l := sq[1], sq[2]
			if yi*m < multiY {
				// 当前正方形在分割线之下
				area += l * min(multiY-yi*m, l*m)
			}
		}
		return area*2 >= totalArea*m
	})

	return float64(multiY) / m
}

func SpiltSquares_3(squares [][]int) float64 {
	// 先按第一种情况找到整数 y 最终答案在 y~y-1 之间
	totalArea := 0
	maxY := 0
	for _, sq := range squares {
		l := sq[2]
		totalArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	calcArea := func(y int) int {
		area := 0
		for _, sq := range squares {
			yi, l := sq[1], sq[2]
			if yi < y {
				area += l * min(l, y-yi)
			}
		}
		return area
	}

	// 在 [0,maxY] 之间二分搜索
	y := sort.Search(maxY, func(y int) bool { return calcArea(y)*2 >= totalArea })
	// 以 y 为分界线，下面的正方形面积
	areaY := calcArea(y)
	// 在 y 和 y-1 之间的面积如下，同时也是长度，因为高度为 1
	diff := areaY - calcArea(y-1)

	// 假设答案为 y1，有 areaY1 = areaY - diff*(y-y1) 又 areaY1*2 = totalArea
	// 故 y1 = y - (areaY - totalArea/2) / diff

	return float64(y) - (float64(areaY*2)-float64(totalArea))/float64(diff*2)
}
