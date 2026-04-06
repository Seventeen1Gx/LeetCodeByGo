package LeetCode

import "math"

func MaxCityMinPower_1(stations []int, r int, k int) int64 {
	// power[i] 表示第 i 座城市的电量
	// d[i] = power[i] - power[i-1] 且 d[0] = power[0] d[n] = -power[n-1]
	// stations[i] 表示对 power[max(0, i-r), min(n-1, i+r)] += stations[i]
	// 即 d[max(0, i-r)] += stations[i] 、d[min(n, i+r+1)]] -= stations[i]

	// 预先计算建造前每座城市电量
	n := len(stations)
	d := make([]int, n+1)
	for i, station := range stations {
		left := max(0, i-r)
		right := min(n, i+r+1)
		d[left] += station
		d[right] -= station
	}

	minVal := math.MaxInt
	power := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			power[i] = d[i]
		} else {
			power[i] = d[i] + power[i-1]
		}
		minVal = min(minVal, power[i])
	}

	// 二分查找，给定 low 是否存在建造 k 座发电站的方案，使所有城市的电量都 >= low
	// 如果存在，low 可以更大，如果不存在，low 需要更小
	// 从左遍历 a 如果 a[i] < low 则它的右侧最远处（贪心）需要建造发电站满足它
	check := func(low int) bool {
		remainingK := k

		d2 := make([]int, n+1)
		copy(d2, d)

		power := 0
		for i := 0; i < n; i++ {
			power += d2[i]
			if power >= low {
				continue
			}

			if remainingK < low-power {
				return false
			}

			// 建造核电站
			remainingK -= low - power
			pos := min(n-1, i+r)
			j := min(n-1, pos+r)
			d2[i] += low - power
			d2[j+1] -= low - power

			power += low - power
		}
		return true
	}

	ans := 0
	low, high := minVal, minVal+k
	for low <= high {
		mid := low + (high-low)/2
		if check(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return int64(ans)
}
