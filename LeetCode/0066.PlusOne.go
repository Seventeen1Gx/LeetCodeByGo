package LeetCode

func plusOne(digits []int) []int {
	n := len(digits)
	ret := make([]int, n+1)

	carry := 1
	for i := n - 1; i >= 0; i-- {
		cur := digits[i] + carry
		ret[i+1] = cur % 10
		carry = cur / 10
	}

	if carry > 0 {
		ret[0] = carry
		return ret
	}

	return ret[1:]
}
