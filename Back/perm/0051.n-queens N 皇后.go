package main

// https://leetcode-cn.com/problems/n-queens

func solveNQueens(n int) [][]string {
	col, as1, as2 := make([]bool, n), map[int]bool{}, map[int]bool{}
	var bufCur = make([]byte, n)
	for i := 0; i < n; i++ {
		bufCur[i] = '.'
	}
	var rowCur = make([]string, n)
	var rows [][]string

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			rows = append(rows, append([]string{}, rowCur...))
			return
		}
		for j := 0; j < n; j++ {
			a1 := idx - j
			a2 := idx + j
			if col[j] || as1[a1] || as2[a2] {
				continue
			}
			col[j], as1[a1], as2[a2] = true, true, true
			bufCur[j] = 'Q'
			rowCur[idx] = string(bufCur)
			bufCur[j] = '.'
			dfs(idx + 1)
			col[j], as1[a1], as2[a2] = false, false, false
		}
	}
	dfs(0)
	return rows
}
