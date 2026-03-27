package LeetCode

import (
	"strconv"
)

func StringSameNum1(str string) bool {
	for {
		n := len(str)
		if n == 0 || n == 2 {
			return false
		}
		if n == 1 {
			return true
		}
		newStr := ""
		lastC := 0
		isSame := true
		for i := 0; i < n-1; i++ {
			a, _ := strconv.Atoi(str[i : i+1])
			b, _ := strconv.Atoi(str[i+1 : i+2])
			c := (a + b) % 10
			if len(newStr) != 0 && c != lastC {
				isSame = false
			}
			newStr += strconv.Itoa(c)
			lastC = c
		}
		if isSame {
			return true
		}
		str = newStr
	}
}
