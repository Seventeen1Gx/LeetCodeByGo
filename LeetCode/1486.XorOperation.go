package LeetCode

func XorOperation(n int, start int) int {
	var ans int
	for i := 0; i < n; i++ {
		ans ^= start + 2*i
	}
	return ans
}
