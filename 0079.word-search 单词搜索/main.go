package main

// https://leetcode-cn.com/problems/word-search

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	l1 := len(board)
	l2 := len(board[0])
	l3 := len(word)
	visited := make([][]bool, l1)
	for i := range visited {
		visited[i] = make([]bool, l2)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == l3-1 {
			return true
		}
		visited[i][j] = true
		for _, dir := range directions {
			ni, nj := i+dir.x, j+dir.y
			if 0 <= ni && ni < l1 && 0 <= nj && nj < l2 && !visited[ni][nj] {
				if check(ni, nj, k+1) {
					return true
				}
			}
		}
		visited[i][j] = false
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
