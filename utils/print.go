package utils

import "fmt"

func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%4d ", matrix[i][j])
		}
		fmt.Println()
	}
}
