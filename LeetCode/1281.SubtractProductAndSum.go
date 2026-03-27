package LeetCode

func SubtractProductAndSum(n int) int {
	a, b := 0, 1
	for n > 0 {
		a += n % 10
		b *= n / 10
		n /= 10
	}
	return b - a
}
