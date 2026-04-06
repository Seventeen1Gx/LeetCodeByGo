package LeetCode

func CarPooling(trips [][]int, capacity int) bool {
	// a[i] 表示车子到达 i 时的乘客数，n 为数组 a 的长度
	// d[i] 定义为 a[i] - a[i-1] 其中首尾边界 d[0] = a[0] d[n] = -a[n-1]
	// trips[i] 就是对 a[from_i] 到 a[to_i-1] 增加 numPassengers_i（因为 to_i 下车即 to_i 时已经不在车子上）
	// 那么就是对 d[from_i] + numPassengers_i 、d[to_i] - numPassengers_i

	d := make([]int, 1001)
	for _, trip := range trips {
		n := trip[0]
		i := trip[1]
		j := trip[2]

		d[i] += n
		d[j] -= n
	}

	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		if i == 0 {
			a[i] = d[i]
		} else {
			a[i] = a[i-1] + d[i]
		}
	}

	for _, x := range a {
		if x > capacity {
			return false
		}
	}

	return true
}
