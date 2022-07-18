package main

// https://leetcode.cn/problems/spiral-matrix-ii

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	var num = 1
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			matrix[top][col] = num
			num++
		}

		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}

		if left < right && top < bottom {
			for col := right - 1; left <= col; col-- {
				matrix[bottom][col] = num
				num++
			}

			for row := bottom - 1; top < row; row-- {
				matrix[row][left] = num
				num++
			}
		}
		top++
		right--
		bottom--
		left++
	}
	return matrix
}
