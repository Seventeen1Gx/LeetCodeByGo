package LeetCode

func RomanToInteger(s string) int {
	hash := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	i, ans := 0, 0
	for i < len(s) {
		if i+1 < len(s) && hash[s[i:i+2]] > 0 {
			ans += hash[s[i:i+2]]
			i += 2
		} else {
			ans += hash[s[i:i+1]]
			i++
		}
	}

	return ans
}
