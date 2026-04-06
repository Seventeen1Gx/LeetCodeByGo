package LeetCode

func RobotCanReturnOrigin(moves string) bool {
	x, y := 0, 0
	for _, m := range moves {
		if m == 'R' {
			x++
		} else if m == 'L' {
			x--
		} else if m == 'U' {
			y++
		} else if m == 'D' {
			y--
		}
	}

	return x == 0 && y == 0
}
