package LeetCode

func CarPooling(trips [][]int, capacity int) bool {
	// a[i] 表示 车子到 i 时乘客数
	// d[i] 是其差分数组，d[i] = a[i]-a[i-1]
	// 在 i~j 上车 n 个乘客，就是对 a[i]、a[i+1]、... 、a[j] 增加 n ，也就是对 d[i]+n ，d[j]-n

	// 0~100
	d := make([]int, 1001)
	for _, trip := range trips {
		n := trip[0]
		i := trip[1]
		j := trip[2]

		d[i] += n
		d[j] -= n
	}

	a := make([]int, 1001)
	for i := 0; i < 1001; i++ {
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
