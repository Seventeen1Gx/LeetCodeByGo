package LeetCode

type NumArray struct {
	// PrefixSum[i] 表示 nums[0:i) 的和
	PrefixSum []int
}

func NumArrayConstructor(nums []int) NumArray {
	numArray := NumArray{PrefixSum: make([]int, len(nums)+1)}
	for i, num := range nums {
		numArray.PrefixSum[i+1] = numArray.PrefixSum[i] + num
	}
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.PrefixSum[right+1] - this.PrefixSum[left]
}
