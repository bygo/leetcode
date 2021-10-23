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

func (this *TicTacToe) Move(row int, col int, player int) int {
	symbol := 1
	if player == 2 {
		symbol = -1
	}
	this.row[row] += symbol
	if this.win(this.row[row]) {
		return player
	}
	this.col[col] += symbol
	if this.win(this.col[col]) {
		return player
	}

	if col-row == 0 {
		this.ang[0] += symbol
		if this.win(this.ang[0]) {
			return player
		}
	}
	if col+row == this.n-1 {
		this.ang[1] += symbol
		if this.win(this.ang[1]) {
			return player
		}
	}
	return 0
}

func (this *TicTacToe) win(i int) bool {
	return i == this.n || i == -this.n
}
