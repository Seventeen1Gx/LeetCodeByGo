package LeetCode

func MyPow_1(x float64, n int) float64 {
	// x15 = x1*x2*x4*x8
	// x6 = x2*x4

	if n == 0 {
		return 1
	}

	if n < 0 {
		return MyPow_1(1/x, -n)
	}

	// 边遍历边获得因子，当遍历到 1 时，把因子乘到结果中

	var ans = 1.0
	for n > 0 {
		if n&1 > 0 {
			ans *= x
		}
		x *= x
		n >>= 1
	}

	return ans
}

func MyPow_2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return MyPow_2(1/x, -n)
	}

	ans := MyPow_2(x, n/2)
	if n%2 == 0 {
		return ans * ans
	}
	return ans * ans * x
}
