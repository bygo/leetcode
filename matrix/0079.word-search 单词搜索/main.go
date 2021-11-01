package main

// https://leetcode-cn.com/problems/word-search

var dirs = [4][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	l1 := len(board)
	l2 := len(board[0])
	top := len(word) - 1
	vis := make([][]bool, l1)
	for i := range vis {
		vis[i] = make([]bool, l2)
	}

	var dfs func(x, y, l int) bool
	dfs = func(x, y, l int) bool {
		if board[x][y] != word[l] {
			return false
		}
		if l == top {
			return true
		}
		vis[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || l1 <= nx || ny < 0 || l2 <= ny || vis[nx][ny] {
				continue
			}
			if dfs(nx, ny, l+1) {
				return true
			}
		}
		vis[x][y] = false
		return false
	}
	for x, row := range board {
		for y := range row {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false
}
