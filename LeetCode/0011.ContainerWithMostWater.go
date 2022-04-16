package LeetCode

import "LeetCode/utils"

func maxArea1(height []int) int {
	// 暴力求解法
	var ans int

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * utils.Min(height[i], height[j])
			if area > ans {
				ans = area
			}
		}
	}
	return ans
}

func maxArea2(height []int) int {
	// 双指针法，两边向中间移动，移动小指针即可，期待它变大。移动大指针已经不管用。
	var ans int
	var i, j = 0, len(height) - 1	// 从两端开始，说明所有列都有机会成为两块板
	for i < j {
		area := (j - i) * utils.Min(height[i], height[j])
		if area > ans {
			ans = area
		}
		if height[i] < height[j] {
			i++	// 说明 height[i] 已经不用考虑它作木板，下面同理（用减治思维想）
		} else {
			j--
		}
	}
	return ans
}
