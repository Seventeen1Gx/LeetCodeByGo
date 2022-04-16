package LeetCode

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

func isPalindrome1(x int) bool {
	// 利用栈，比较位
	if x < 0 {
		return false
	}

	var n = 0
	var t = x
	for t > 0 {
		t /= 10
		n++
	}

	var stack = make([]int, n/2)
	for i := 0; i < n/2; i++ {
		stack[i] = x % 10
		x /= 10
	}

	if n%2 == 1 {
		x /= 10
	}

	for i := n/2 - 1; i >= 0; i-- {
		if stack[i] != x%10 {
			return false
		}
		x /= 10
	}
	return true
}

func isPalindrome2(x int) bool {
	// 直接反转一半数字进行比较

	if x < 0 || (x%10 == 0 && x != 0) {
		// 负数和末尾为 0 的数字都不是回文数
		return false
	}

	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}

	// 奇数和偶数两种情况
	return x == rev || x == rev/10
}
