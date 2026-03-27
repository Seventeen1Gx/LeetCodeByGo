package LeetCode

func IsPowerOfFour_1(n int) bool {
	if n <= 0 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

func IsPowerOfFour_2(n int) bool {
	// 4 的幂首先是 2 的幂
	// 其次每乘 4 就是左移两格
	// 1 : 0001
	// 4 : 0100
	// 16 : 1 0000
	// 64 : 100 0000
	// ...
	// 奇数位是 1 ，其他位是 0
	return n > 0 && n&(n-1) == 0 && n&0x55555555 > 0
}
