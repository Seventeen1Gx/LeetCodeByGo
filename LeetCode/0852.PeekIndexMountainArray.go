package LeetCode

func PeekIndexInMountainArray(arr []int) int {
	// 二分法
	n := len(arr)
	left, right := 1, n-2
	for left <= right {
		mid := (left + right) / 2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid]+1 {
			return mid
		} else if arr[mid] > arr[mid-1] {
			// 在上坡
			left = mid
		} else {
			right = mid
		}
	}
	return -1
}
