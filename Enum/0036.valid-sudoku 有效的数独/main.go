package main

// https://leetcode.cn/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	row := [9]map[byte]bool{}
	col := [9]map[byte]bool{}
	box := [9]map[byte]bool{}

	for i := 0; i < 9; i++ {
		row[i] = map[byte]bool{}
		col[i] = map[byte]bool{}
		box[i] = map[byte]bool{}
	}

	for rowK, rowV := range board {
		for colK, colV := range rowV {
			if board[rowK][colK] != '.' {
				if row[rowK][colV] == true || col[colK][colV] == true || box[(rowK/3)*3+colK/3][colV] == true {
					return false
				}
				row[rowK][colV] = true
				col[colK][colV] = true
				box[(rowK/3)*3+colK/3][colV] = true
			}
		}
	}
	return true
}
