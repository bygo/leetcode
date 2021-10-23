package main

// https://leetcode-cn.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				num := board[row][col] - '1'
				box := (row/3)*3 + col/3
				rows[row][num], cols[col][num], boxes[box][num] = true, true, true
			}
		}
	}
	var dfs func(n int) bool
	dfs = func(n int) bool {
		if n == 81 {
			return true
		}
		var (
			row = n / 9
			col = n % 9
		)
		if board[row][col] != '.' {
			return dfs(n + 1)
		}

		var box = (row/3)*3 + col/3
		for num := 0; num < 9; num++ {
			if !rows[row][num] && !cols[col][num] && !boxes[box][num] {
				board[row][col] = byte(num + '1') // 尝试填入
				rows[row][num], cols[col][num], boxes[box][num] = true, true, true

				if dfs(n + 1) { // 全部成功
					return true
				}
				rows[row][num], cols[col][num], boxes[box][num] = false, false, false // 重置
			}
		}
		board[row][col] = '.'

		// 没有个一个值符合要求
		return false
	}
	dfs(0)
}
