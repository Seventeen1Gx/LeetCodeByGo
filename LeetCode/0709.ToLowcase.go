package LeetCode

func ToLowcase(s string) string {
	var ans []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			ans = append(ans, s[i]-'A'+'a')
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
