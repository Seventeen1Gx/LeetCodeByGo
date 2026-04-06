package LeetCode

import "math"

// 整数反转
// 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。

func ReverseInteger(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit // 在赋值之前，需要满足 -2147483648 <= rev*10+digit <= 2147483647，不然就溢出了，因为 rev 是 int32
	}
	return
}

// rev*10+digit <= 2147483647
// rev*10+digit <= 2147483647/10*10 + 7
// (rev - math.MaxInt32/10)*10 <= 7-digit
// 若 rev - math.MaxInt32/10 > 0，已知 digit >= 0，不等式不成立
// 若 rev - math.MaxInt32/10 = 0，则需要 digit <= 7
// 而此时 rev = math.MaxInt32/10，还要加一位，那么 x 和 2147483647 位数相同，则 x 最高位 digit <= 2，即这时不等式总成立
// 若 rev - math.MaxInt32/10 < 0，则左边小于等于 -10，不等式也成立
// 综上，想让 rev*10+digit <= 2147483647 成立，需要 rev - math.MaxInt32/10 <= 0

// 注意这里，digit∈[-9,0]
// -2147483648 <= rev*10+digit
// -2147483648/10*10-8 <= rev*10+digit
// (math.MinInt32/10-rev) * 10 <= digit + 8
// 若 math.MinInt32/10-rev < 0，则左边最大为 -10，右边最小为 -1，故不等式总成立
// 若 math.MinInt32/10-rev = 0，则左边最大为 -10，右边最小为 -1，但同理整数最高位为 -1 或者 -2，故不等式总成立
// 若 math.MinInt32/10-rev > 0，则左边最小为 10，右边最大为 8，故不等式总不成立
