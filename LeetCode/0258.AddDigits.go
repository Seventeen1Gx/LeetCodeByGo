package LeetCode

func AddDigits_1(num int) int {
	for num > 9 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	return num
}

func AddDigits_2(num int) int {
	return (num-1)%9 + 1
}
