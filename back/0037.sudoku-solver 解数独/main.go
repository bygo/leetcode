package main

// https://leetcode-cn.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	var row = [9][9]bool{}
	var col = [9][9]bool{}
	var box = [9][9]bool{}
	for rowK := 0; rowK < 9; rowK++ {
		for colK := 0; colK < 9; colK++ {
			if board[rowK][colK] != '.' {
				var num = board[rowK][colK] - '1'
				row[rowK][num], col[colK][num], box[rowK/3*3+colK/3][num] = true, true, true
			}
		}
	}

	var dfs func(n int) bool
	dfs = func(n int) bool {
		if n == 81 {
			return true
		}
		rowK := n / 9
		colK := n % 9
		if board[rowK][colK] != '.' {
			return dfs(n + 1)
		}

		var boxK = (rowK/3)*3 + colK/3
		for num := 0; num < 9; num++ {
			if !row[rowK][num] && !col[colK][num] && !box[boxK][num] {
				// 尝试填充
				board[rowK][colK] = byte('1' + num)
				row[rowK][num], col[colK][num], box[boxK][num] = true, true, true

				if dfs(n + 1) { //下一个填充
					return true
				}
				//失败回溯
				row[rowK][num], col[colK][num], box[boxK][num] = false, false, false
			}
		}
		board[rowK][colK] = '.'
		return false
	}
	dfs(0)
}
