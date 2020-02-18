package main

func solveSudoku(board [][]byte) {
	row := [9][9]bool{}
	col := [9][9]bool{}
	box := [9][9]bool{}
	for rowK := 0; rowK < 9; rowK++ {
		for colK := 0; colK < 9; colK++ {
			if board[rowK][colK] != '.' {
				num := board[rowK][colK] - '1'
				boxK := (rowK/3)*3 + colK/3
				row[rowK][num], col[colK][num], box[boxK][num] = true, true, true
			}
		}
	}
	fill(board, 0, row, col, box)
}

func fill(board [][]byte, n int, row [9][9]bool, col [9][9]bool, box [9][9]bool) bool {
	if n == 81 {
		return true
	}
	rowK := n / 9
	colK := n % 9
	if board[rowK][colK] != '.' {
		return fill(board, n+1, row, col, box)
	}

	boxK := (rowK/3)*3 + colK/3
	for num := 0; num < 9; num++ {
		if !row[rowK][num] && !col[colK][num] && !box[boxK][num] {
			board[rowK][colK] = byte('1' + num)
			row[rowK][num], col[colK][num], box[boxK][num] = true, true, true
			if fill(board, n+1, row, col, box) { //下一个填充
				return true
			}
			row[rowK][num], col[colK][num], box[boxK][num] = false, false, false //失败回溯
		}
	}
	board[rowK][colK] = '.'
	return false
}
