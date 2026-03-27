package LeetCode

import "math"

func MinNumOfSeconds(mountainHeight int, workerTimes []int) (ans int64) {
	// 如果 t 秒能完成工作，那么大于 t 秒都能完成工作
	// 答案具有单调性，于是想到二分查找
	// 二分查找的下界是 1，上界是最慢工人移动完山的时间，找到最小可以移完山的时间
	// 二分查找的每一步，我们需要判断 mid 秒是否所有工人可以一起将山降低
	// 假设在 mid 秒里每个工人都竭尽所能，第 i 个工人降低山高为 k 则
	// workerTimes[i] * (1 + 2 + 3 + ... + k) = workerTimes[i] * (1 + k) * k / 2 <= mid
	// k * k + k - 2 * mid / workerTimes[i] = 0
	// 解方程 k = (-1 + 根号(1+8*mid/workerTimes[i])) / 2

	maxWorkerTime := 0
	for _, workerTime := range workerTimes {
		if workerTime > maxWorkerTime {
			maxWorkerTime = workerTime
		}
	}

	left := 1
	right := maxWorkerTime * (1 + mountainHeight) * mountainHeight / 2
	for left <= right {
		mid := (left + right) / 2
		cnt := 0

		for _, workerTime := range workerTimes {
			work := mid / workerTime
			// 该名工人在 mid 秒最大移动山的高度
			k := (-1 + math.Sqrt(float64(1+8*work))) / 2
			cnt += int(k)
		}

		if cnt >= mountainHeight {
			ans = int64(mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
