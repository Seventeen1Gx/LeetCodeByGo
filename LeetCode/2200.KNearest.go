package LeetCode

import "LeetCodeByGo/utils"

func KNearest_1(nums []int, key int, k int) (ans []int) {
	n := len(nums)

	// 每个元素距离其左侧 key 元素的距离
	leftKeyDistance := make([]int, n)
	for i := 0; i < n; i++ {
		num := nums[i]
		if num == key {
			leftKeyDistance[i] = 0
		} else if i == 0 || leftKeyDistance[i-1] == n+1 {
			leftKeyDistance[i] = n + 1
		} else {
			leftKeyDistance[i] = leftKeyDistance[i-1] + 1
		}
	}

	// 每个元素距离其右侧 key 元素的距离
	rightKeyDistance := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		if num == key {
			rightKeyDistance[i] = 0
		} else if i == n-1 || rightKeyDistance[i+1] == n+1 {
			rightKeyDistance[i] = n + 1
		} else {
			rightKeyDistance[i] = rightKeyDistance[i+1] + 1
		}
	}

	for i := range nums {
		if leftKeyDistance[i] <= k || rightKeyDistance[i] <= k {
			ans = append(ans, i)
		}
	}

	return ans
}

func KNearest_2(nums []int, key int, k int) (ans []int) {
	// 滑动窗口，[max(0,i-k),min(n-1,i+k)]
	hashSet := make(map[int]int)

	// 特殊处理 [0:k-1]
	// 这样之后就可以逐步移动右端点
	for _, num := range nums[:k] {
		hashSet[num]++
	}

	for i := range nums {
		// 右端点进入窗口
		if i+k < len(nums) {
			hashSet[nums[i+k]]++
		}

		// 更新答案
		if hashSet[key] > 0 {
			ans = append(ans, i)
		}

		// 左端点移出窗口
		if i-k >= 0 {
			hashSet[nums[i-k]]--
		}
	}

	return ans
}

func KNearest_3(nums []int, key int, k int) (ans []int) {
	// 滑动窗口，[max(0,i-k),min(n-1,i+k)]
	// 实际不用记录每个窗口中的元素
	// 只要记录目前为止遇到的所有 key 中，最靠右的那个位置
	last := -k - 1 // 这个初始值是为了后面的判断

	// 处理第一个窗口，然后不断新增右端点进行更新
	for i := k - 1; i >= 0; i-- {
		if nums[i] == k {
			last = i
			break
		}
	}

	for i := range nums {
		// 右端点进入窗口
		if i+k < len(nums) && nums[i+k] == key {
			last = i + k
		}
		// 更新答案，last 还在我们的窗口中
		if last >= i-k {
			// 窗口中含有 Key 元素
			ans = append(ans, i)
		}
		// 因为我们总是记录最靠右的 Key 所以左端点自然就移出窗口了
	}

	return ans
}

// 上面的方法是从位置 i 去看 key
// 下面的方法是从 key 去看位置 i

func KNearest_4(nums []int, key int, k int) (ans []int) {
	// 标记法
	n := len(nums)
	mark := make([]bool, n)

	for i, num := range nums {
		if num == key {
			// 标记 key 前后范围的元素
			left := utils.Max(0, i-k)
			right := utils.Min(n-1, i+k)
			for j := left; j <= right; j++ {
				mark[j] = true
			}
		}
	}

	for i, ok := range mark {
		if ok {
			ans = append(ans, i)
		}
	}

	return ans
}

func KNearest_5(nums []int, key int, k int) (ans []int) {
	// 找到 key 元素，它的左边和右边 k 范围都可以加入答案
	// 为了避免重复，我们维护上次加入到答案的位置
	n := len(nums)
	last := 0
	for i, num := range nums {
		if num != key {
			continue
		}

		last = utils.Max(last, i-k)

		// 将 [i-k,i+k] 加入答案
		for last <= utils.Min(n-1, i+k) {
			ans = append(ans, last)
			last++
		}
	}

	return ans
}
