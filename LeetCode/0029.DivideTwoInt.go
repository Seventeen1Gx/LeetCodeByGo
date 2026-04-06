package LeetCode

import "math"

// 除法的本质就是被除数能减除数多少次
func DivideTwoInt(dividend int, divisor int) int {
	// 排除一些特殊情况
	if divisor == 1 {
		return dividend
	}
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		// 除法上溢出
		return math.MaxInt32
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}

	// 结果符号
	negtive := (dividend > 0) != (divisor > 0)

	// 统一用负数处理，避免正溢出
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	// 快速减（倍增）
	res := 0
	for dividend <= divisor { // |dividend| >= |divisor|
		val := divisor
		count := 1
		for val >= math.MinInt32/2 && val+val >= dividend {
			val += val
			count += count
		}
		dividend -= val
		res += count
	}

	if negtive {
		res = -res
	}

	return res
}
