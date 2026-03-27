package LeetCode

func AddString(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	carry := uint8(0)

	ans := make([]uint8, max(n, m)+1)
	for i := 0; i < m || i < n; i++ {
		var a uint8
		if n-i-1 >= 0 {
			a = num1[n-i-1] - '0'
		}

		var b uint8
		if m-i-1 >= 0 {
			b = num2[m-i-1] - '0'
		}

		c := a + b + carry
		ans[len(ans)-i-1] = '0' + c%10
		carry = c / 10
	}
	if carry > 0 {
		ans[0] = '0' + carry
	} else {
		ans = ans[1:]
	}

	return string(ans)
}
