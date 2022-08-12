package main

// https://leetcode.cn/problems/check-if-word-can-be-placed-in-crossword

func placeWordInCrossword(board [][]byte, word string) bool {
	l1 := len(board)
	l2 := len(board[0])
	l3 := len(word)
	for i, row := range board {
		for j, _ := range row {
			if i == 0 || board[i-1][j] == '#' {
				var k int
				for k < l3 && i+k < l1 && (board[i+k][j] == ' ' || board[i+k][j] == word[k]) {
					k++
				}
				if k == l3 && (i+k == l1 || board[i+k][j] == '#') {
					return true
				}
			}

			if j == 0 || board[i][j-1] == '#' {
				var k int
				for k < l3 && j+k < l2 && (board[i][j+k] == ' ' || board[i][j+k] == word[k]) {
					k++
				}
				if k == l3 && (j+k == l2 || board[i][j+k] == '#') {
					return true
				}
			}

			if i == l1-1 || board[i+1][j] == '#' {
				var k = 0
				for k < l3 && 0 <= i-k && (board[i-k][j] == ' ' || board[i-k][j] == word[k]) {
					k++
				}
				if k == l3 && (i-k == -1 || board[i-k][j] == '#') {
					return true
				}
			}

			if j == l2-1 || board[i][j+1] == '#' {
				var k = 0
				for k < l3 && 0 <= j-k && (board[i][j-k] == ' ' || board[i][j-k] == word[k]) {
					k++
				}
				if k == l3 && (j-k == -1 || board[i][j-k] == '#') {
					return true
				}
			}
		}
	}
	return false
}
