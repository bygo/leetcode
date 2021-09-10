package main

// https://leetcode-cn.com/problems/set-matrix-zeroes

func setZeroes(matrix [][]int) {
	l1 := len(matrix)
	l2 := len(matrix[0])
	var col bool
	for i := range matrix {
		if matrix[i][0] == 0 {
			col = true
		}
		for j := 1; j < l2; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := l1 - 1; 0 <= i; i-- {
		for j := 1; j < l2; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if col {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
