package main

// https://leetcode-cn.com/problems/set-matrix-zeroes

func setZeroes(matrix [][]int) {
	l1 := len(matrix)
	l2 := len(matrix[0])
	var col bool // 第一列是否出现过
	for i := range matrix {
		if matrix[i][0] == 0 {
			col = true
		}
		for j := 1; j < l2; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 所有列 0 ，归第一列
				matrix[0][j] = 0 // 所有行 0， 归第一行
			}
		}
	}

	for i := l1 - 1; 0 <= i; i-- { // 防止第一行被更新，当然也可 存 任何行 任何列
		for j := 1; j < l2; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
		if col { // 第一列是否有 0
			matrix[i][0] = 0
		}
	}
}
