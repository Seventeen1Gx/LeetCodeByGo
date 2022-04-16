package LeetCode

import (
	"LeetCode/utils"
	"math"
)

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的[中位数]。
//
// 提示：
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10^6 <= nums1[i], nums2[i] <= 10^6

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if (m+n)%2 == 0 {
		// 偶数长度
		return float64(findKthVal(nums1, 0, nums2, 0, (m+n)/2+1)+findKthVal(nums1, 0, nums2, 0, (m+n)/2)) / 2.0
	} else {
		// 奇数长度
		return float64(findKthVal(nums1, 0, nums2, 0, (m+n)/2+1))
	}
}

// findKthVal 寻找两排序数组 nums1[i:] 和 nums2[j:] 的第 k 个数
func findKthVal(nums1 []int, i int, nums2 []int, j int, k int) int {
	// 保证数组 1 的长度小于或等于数组 2 的长度
	m := len(nums1)
	n := len(nums2)
	if m-i > n-j {
		return findKthVal(nums2, j, nums1, i, k)
	}

	// 如果小数组为空，则返回大数组的第 k 个数
	if m-i == 0 {
		return nums2[j+k-1] // k 从 1 开始
	}

	if k == 1 {
		// 表示找第 1 个数，则比较两数组的首元素，取较小值即可
		return utils.Min(nums1[i], nums2[j])
	}

	// 取小数组开头的 k/2 个元素，若不够，则取全部元素
	pa := utils.Min(i+k/2, m)
	// 由 pa 计算 pb，有 [i:pa) 与 [j:pb) 元素个数为 k
	pb := j + k - (pa - i)
	if nums1[pa-1] < nums2[pb-1] {
		// num1[pa-1] 小于合并之后的第 k 小值（反证法，令该数为第 k+1 个数，则 num2[pb-1] 为第 k+2 个数，但他们之前没这么多元素）
		// 故排除数组 1 分界线之前的部分
		return findKthVal(nums1, pa, nums2, j, k-pa+i)
	} else if nums1[pa-1] > nums2[pb-1] {
		// 同理，排除数组 2 分界线之前的部分
		return findKthVal(nums1, i, nums2, pb, k-pb+j)
	} else {
		return nums1[pa-1]
	}
}

// 假设中位数 median 将 nums1 和 nums2 分到两个集合 A 和 B 中
// 且有 A[...] <= median、B[...] > median
// nums1(-1:i) 和 nums2(-1:j) 在 A 中，nums1[i:) 和 nums[j:) 在 B 中
// 根据中位数定义，len(A) = len(B) 或 len(A) = len(B) + 1
// 则可知 i + j = (m + n) / 2 或 i + j = (m + n + 1) / 2，前者 m+n 为偶数，后者为奇数
// 根据整除，可统一用后者表示 i 和 j 的关系
// 则本问题抽象成，在 [0,m] 中找到一个 i，使得 A 中最大数小于 B 中最小数，即 nums1[i-1] <= nums2[j] 且 nums2[j-1] <= nums1[i]
// 又由于有边界问题，这四个数不一定存在
// 为了统一边界问题，在 nums1 和 nums2 的首尾都插入一个元素，并保证有序性

func  findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		// 保证 nums1 为小数组，缩小搜索范围
		return findMedianSortedArrays2(nums2, nums1)
	}

	maxNum := int(math.Pow(10, 6))
	minNum := -maxNum

	totalLeftLen := (m + n + 1) / 2 // 集合 A 中元素总个数

	// 搜索范围 [0,m]
	iMin, iMax := 0, m
	var leftMax_1 int  // nums1[i-1]
	var leftMax_2 int  // nums2[j-1]
	var rightMin_1 int // nums1[i]
	var rightMin_2 int // nums2[j]
	for iMin <= iMax {
		// 二分搜索
		i := iMin + (iMax-iMin)/2
		// 根据 i 与 j 的关系获得 j
		j := totalLeftLen - i

		// 获得四个分界数，数组边界外用额外添加的数
		leftMax_1 = minNum
		if i != 0 {
			leftMax_1 = nums1[i-1]
		}
		leftMax_2 = minNum
		if j != 0 {
			leftMax_2 = nums2[j-1]
		}
		rightMin_1 = maxNum
		if i != m {
			rightMin_1 = nums1[i]
		}
		rightMin_2 = maxNum
		if j != n {
			rightMin_2 = nums2[j]
		}

		if leftMax_1 <= rightMin_2 && leftMax_2 <= rightMin_1 {
			// 找到满足条件的 i 了，计算中位数
			leftMax := utils.Max(leftMax_1, leftMax_2)	// max 保证了会舍去额外添加的数，下面 min 同理
			if (m+n)%2 == 1 {
				return float64(leftMax)
			} else {
				rightMin := utils.Min(rightMin_1, rightMin_2)
				return float64(leftMax+rightMin) / 2.0
			}
		} else if leftMax_1 > rightMin_2 {
			// A 中来源于数组 1 的最大数过于大了，则缩小 i
			iMax = i - 1
		} else {
			// B 中来源与数组 2 的最大数过于大了，则缩小 j，即扩大 i
			iMin = i + 1
		}
	}
	return 0
}
