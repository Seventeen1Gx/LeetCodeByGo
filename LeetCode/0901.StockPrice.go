package LeetCode

// 求左侧第一个更大值的下标
// 100 80 60 70 70 60 75 85
// 从左往右遍历，可能是未来某个元素左侧第一个更大值的下标保留在栈中
// 遍历到 100 栈为空，故 100 的跨度是 1，下标入栈
// 遍历到 80 栈顶大于当前值，故 80 的跨度是 1，下标入栈
// 遍历到 60 同理入栈
// 遍历到 70 栈顶元素小于当前值，出栈，因为 60 不可能是接下来元素左侧第一个更大值了 70 挡着的
// 60 出完栈，80 是栈顶，故和栈顶计算差值，得到跨度

type Stock struct {
	idx   int
	price int
}

type StockSpanner struct {
	i     int // 下一个元素的下标
	stack []*Stock
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.stack) > 0 && this.stack[len(this.stack)-1].price <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	ans := this.i + 1
	if len(this.stack) > 0 {
		ans = this.i - this.stack[len(this.stack)-1].idx
	}
	this.stack = append(this.stack, &Stock{idx: this.i, price: price})
	this.i++
	return ans
}
