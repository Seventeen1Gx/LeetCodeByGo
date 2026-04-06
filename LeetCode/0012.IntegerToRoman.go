package LeetCode

// 1 <= num <= 3999
func IntToRoman_1(num int) string {
	ans := ""
	for num > 0 {
		if num >= 1000 {
			ans += "M"
			num -= 1000
		} else if num >= 900 {
			ans += "CM"
			num -= 900
		} else if num >= 500 {
			ans += "D"
			num -= 500
		} else if num >= 400 {
			ans += "CD"
			num -= 400
		} else if num >= 100 {
			ans += "C"
			num -= 100
		} else if num >= 90 {
			ans += "XC"
			num -= 90
		} else if num >= 50 {
			ans += "L"
			num -= 50
		} else if num >= 40 {
			ans += "XL"
			num -= 40
		} else if num >= 10 {
			ans += "X"
			num -= 10
		} else if num >= 9 {
			ans += "IX"
			num -= 9
		} else if num >= 5 {
			ans += "V"
			num -= 5
		} else if num >= 4 {
			ans += "IV"
			num -= 4
		} else if num >= 1 {
			ans += "I"
			num -= 1
		}
	}
	return ans
}

func IntToRoman_2(num int) string {
	maps := []struct {
		val int
		str string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	ans := ""
	for _, v := range maps {
		for num >= v.val {
			ans += v.str
			num -= v.val
		}
		if num == 0 {
			break
		}
	}

	return ans
}
