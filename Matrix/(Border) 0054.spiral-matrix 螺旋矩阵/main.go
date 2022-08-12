package main

// https://leetcode.cn/problems/spiral-matrix

func spiralOrder(matrix [][]int) []int {
	var rows, cols = len(matrix), len(matrix[0])
	var total = rows * cols
	if total == 0 {
		return nil
	}
	var (
		res                      = make([]int, 0, total)
		left, right, top, bottom = 0, cols - 1, 0, rows - 1
	)
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		for row := top + 1; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}

		if left < right && top < bottom {
			for col := right - 1; left <= col; col-- {
				res = append(res, matrix[bottom][col])
			}
			for row := bottom - 1; top < row; row-- {
				res = append(res, matrix[row][left])
			}
		}
		top++
		right--
		left++
		bottom--
	}
	return res
}
