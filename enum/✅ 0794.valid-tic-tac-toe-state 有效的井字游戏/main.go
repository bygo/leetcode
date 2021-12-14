package main

import "strings"

// https://leetcode-cn.com/problems/valid-tic-tac-toe-state

func win(board []string, p byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == p && board[i][1] == p && board[i][2] == p ||
			board[0][i] == p && board[1][i] == p && board[2][i] == p {
			return true
		}
	}

	return board[0][0] == p && board[1][1] == p && board[2][2] == p ||
		board[0][2] == p && board[1][1] == p && board[2][0] == p
}

func validTicTacToe(board []string) bool {
	oCnt, xCnt := 0, 0
	for _, row := range board {
		oCnt += strings.Count(row, "O")
		xCnt += strings.Count(row, "X")
	}

	if xCnt != oCnt && xCnt != oCnt+1 { // 不符合顺序
		return false
	}

	if xCnt == oCnt+1 && win(board, 'O') { // 本该x赢
		return false
	}

	if xCnt == oCnt && win(board, 'X') { // 本该O赢
		return false
	}

	return true
}
