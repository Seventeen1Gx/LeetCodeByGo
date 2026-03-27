package LeetCode

// 懒更新，直到 Get 时才计算真实结果
// 假设只有加法操作，使用 add 变量统计累加量，每次 append 元素的时候减去当前累加量
// 假设只有乘法操作，使用 mul 变量统计累乘量，每次 append 元素的时候除以当前累乘量
// 当加法操作和乘法操作共存时：
// 对于 AddAll 方法，由于 [v * mul + add] + inc = v * mul + add + inc，所以将 add += inc
// 对于 MultAll 方法，由于 [v * mul + add] * m = v * mul * m + add * m，所以将 mul *= m 且 add *= m
// 对于 Append 方法，由于 v * mul + add == val，加入元素需要 (val - add) / mul

type Fancy struct {
	seq []int
	add int
	mul int
}

const M = 1e9 + 7

func FancyConstructor() Fancy {
	return Fancy{
		mul: 1,
	}
}

func fancyPow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 > 0 {
			res = res * x % M
		}
		x = x * x % M
		n /= 2
	}
	return res
}

func (this *Fancy) Append(val int) {
	// 在 mod 运算中，除以一个数，等于乘上他的乘法逆元
	this.seq = append(this.seq, (val-this.add+mod)*fancyPow(this.mul, M-2)%M)
}

func (this *Fancy) AddAll(inc int) {
	this.add = (this.add + inc) % mod
}

func (this *Fancy) MultAll(m int) {
	this.mul = this.mul * m % mod
	this.add = this.add * m % mod
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.seq) {
		return -1
	}

	return (this.seq[idx]*this.mul%mod + this.add) % mod
}
