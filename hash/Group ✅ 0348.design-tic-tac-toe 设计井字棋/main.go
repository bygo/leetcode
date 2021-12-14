package main

// https://leetcode-cn.com/problems/design-tic-tac-toe

// ❓井字棋 n不止3

var symbols = map[int]int{
	1: 1,
	2: -1,
}

type TicTacToe struct {
	rowMpCnt []int
	colMpCnt []int
	angMpCnt [2]int
	n        int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rowMpCnt: make([]int, n),
		colMpCnt: make([]int, n),
		angMpCnt: [2]int{},
		n:        n,
	}
}

func (t *TicTacToe) Move(row int, col int, player int) int {
	symbol := symbols[player]
	t.rowMpCnt[row] += symbol
	if t.win(t.rowMpCnt[row]) {
		return player
	}
	t.colMpCnt[col] += symbol
	if t.win(t.colMpCnt[col]) {
		return player
	}

	if col == row {
		t.angMpCnt[0] += symbol
		if t.win(t.angMpCnt[0]) {
			return player
		}
	}
	if col+row == t.n-1 {
		t.angMpCnt[1] += symbol
		if t.win(t.angMpCnt[1]) {
			return player
		}
	}
	return 0
}

func (t *TicTacToe) win(i int) bool {
	return i == t.n || i == -t.n
}
