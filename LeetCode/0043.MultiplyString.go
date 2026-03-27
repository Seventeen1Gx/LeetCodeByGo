package LeetCode

func MultiplyString(num1 string, num2 string) string {
	// 模拟竖式乘法
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)

	sum := make([]uint8, 1) // "0"
	for j := n - 1; j >= 0; j-- {
		product := make([]uint8, 0)
		for k := 0; k < n-1-j; k++ {
			product = append(product, 0)
		}

		carry := uint8(0)
		digit1 := num2[j] - '0'
		for i := m - 1; i >= 0; i-- {
			digit2 := num1[i] - '0'
			result := digit1*digit2 + carry
			product = append(product, result%10)
			carry = result / 10
		}
		if carry > 0 {
			product = append(product, carry)
		}

		// sum + product
		carry = 0
		result := make([]uint8, max(len(sum), len(product))+1)
		l1, l2, l3 := 0, 0, 0
		for l1 < len(sum) || l2 < len(product) {
			var a uint8
			if l1 < len(sum) {
				a = sum[l1]
				l1++
			}

			var b uint8
			if l2 < len(product) {
				b = product[l2]
				l2++
			}

			c := a + b + carry
			result[l3] = c % 10
			carry = c / 10
			l3++
		}
		if carry > 0 {
			result[l3] = carry
		} else {
			result = result[:l3]
		}
		sum = result
	}

	var ans = make([]uint8, len(sum))
	for i, v := range sum {
		ans[len(sum)-i-1] = '0' + v
	}

	return string(ans)
}
