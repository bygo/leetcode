package main

// https://leetcode-cn.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	rows, cols, boxes := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for row := range board {
		for col := range board {
			ch := board[row][col]
			if board[row][col] != '.' {
				ch -= '1'
				box := (row/3)*3 + col/3
				rows[row][ch], cols[col][ch], boxes[box][ch] = true, true, true
			}
		}
	}

	var dfs func(n int) bool
	dfs = func(n int) bool {
		if n == 81 {
			return true
		}

		row := n / 9
		col := n % 9
		if board[row][col] != '.' {
			return dfs(n + 1)
		}

		var num byte
		for ; num < 9; num++ {
			box := (row/3)*3 + col/3
			if !rows[row][num] && !cols[col][num] && !boxes[box][num] {
				rows[row][num], cols[col][num], boxes[box][num] = true, true, true
				board[row][col] = num + '1'
				if dfs(n + 1) {
					return true
				}
				rows[row][num], cols[col][num], boxes[box][num] = false, false, false
			}
		}
		board[row][col] = '.'
		return false
	}
	dfs(0)
}
