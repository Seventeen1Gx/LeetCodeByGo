package LeetCode

import "fmt"

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

func convert1(s string, numRows int) string {
	// 仿真法

	// 先排除只有一行的
	if numRows == 1 {
		return s
	}

	var rows = make([][]uint8, numRows)
	var n = len(s)
	var goDown bool
	var curRow int
	for i := 0; i < n; i++ {
		rows[curRow] = append(rows[curRow], s[i])
		if curRow == 0 || curRow == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			curRow++
		} else {
			curRow--
		}
	}
	var ans string
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(rows[i]); j++ {
			ans += fmt.Sprintf("%c", rows[i][j])
		}
	}
	return ans
}

func convert2(s string, numRows int) string {
	// 规律法

	if numRows == 1 {
		return s
	}

	var ans string
	for curRow := 0; curRow < numRows; curRow++ {
		var goDown = true
		for i, step := curRow, 0; i < len(s); i += step {
			ans += fmt.Sprintf("%c", s[i])
			if curRow == 0 || curRow == numRows-1 {
				step = 2 * (numRows - 1)
			} else if goDown {
				// 向下，再向上
				step = 2 * (numRows - curRow - 1)
				goDown = !goDown
			} else {
				// 向上，再向下
				step = 2 * (curRow - 0)
				goDown = !goDown
			}
		}
	}
	return ans
}
