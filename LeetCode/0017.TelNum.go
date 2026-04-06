package LeetCode

func TelNum(digits string) []string {
	n := len(digits)
	ans := make([]string, 0)
	hash := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var backtrace func(i int, cur string)
	backtrace = func(i int, cur string) {
		if i == n {
			ans = append(ans, cur)
			return
		}
		chars := hash[digits[i]]
		for _, ch := range chars {
			backtrace(i+1, cur+string(ch))
		}
	}
	backtrace(0, "")
	return ans
}
