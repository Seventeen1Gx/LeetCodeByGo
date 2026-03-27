package LeetCode

func NumOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	windowSum := 0
	ans := 0
	for right := 0; right < n; right++ {
		windowSum += arr[right]
		left := right - k + 1
		if left < 0 {
			continue
		}

		if float64(windowSum)/float64(k) >= float64(threshold) {
			ans++
		}
		windowSum -= arr[left]
	}
	return ans
}
