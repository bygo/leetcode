package main

// https://leetcode-cn.com/problems/design-tic-tac-toe

type TicTacToe struct {
	row []int
	col []int
	ang [2]int
	n   int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		row: make([]int, n),
		col: make([]int, n),
		ang: [2]int{},
		n:   n,
	}
}

func (ttt *TicTacToe) Move(row int, col int, player int) int {
	symbol := 1
	if player == 2 {
		symbol = -1
	}
	ttt.row[row] += symbol
	if ttt.win(ttt.row[row]) {
		return player
	}
	ttt.col[col] += symbol
	if ttt.win(ttt.col[col]) {
		return player
	}

	if col-row == 0 {
		ttt.ang[0] += symbol
		if ttt.win(ttt.ang[0]) {
			return player
		}
	}
	if col+row == ttt.n-1 {
		ttt.ang[1] += symbol
		if ttt.win(ttt.ang[1]) {
			return player
		}
	}
	return 0
}

func (ttt *TicTacToe) win(i int) bool {
	return i == ttt.n || i == -ttt.n
}
