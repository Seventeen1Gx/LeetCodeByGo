package LeetCode

func FourSumAdd(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	count := 0

	/*
		sort.Ints(nums1)
		sort.Ints(nums2)
		sort.Ints(nums3)
		sort.Ints(nums4)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				k := 0
				l := n - 1
				for k < n && l >= 0 {
					sum := nums1[i] + nums2[j] + nums3[k] + nums4[l]
					if sum == 0 {
						// 错误，无法在两个独立的数组上使用双指针法
						// [-1, 0, 1]
						// [-1, 0, 1]
						// [0, 0, 1]
						// [-1, 1, 1]
						// 当 i = j = 0 时，需要 nums3[k] + nums4[l] == 2
						// 找到第一个满足条件的结果时，k=2，l=2
						// 然后 k++ 得到 3，l-- 得到 1，不满足条件跳出循环
						// 但其实还有 k=2，l=1 满足结果条件
						count++
						k++
						l--
					} else if sum > 0 {
						l--
					} else {
						k++
					}
				}
			}
		}
	*/
	hashSet := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			hashSet[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count += hashSet[-nums3[i]-nums4[j]]
		}
	}

	return count
}
