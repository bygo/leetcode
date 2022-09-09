package main

// https://leetcode.cn/problems/n-queens

func solveNQueens(n int) [][]string {
	col, as1, as2 := make([]bool, n), map[int]bool{}, map[int]bool{}
	var buf = make([]byte, n)
	var row = make([]string, n)

	for i := 0; i < n; i++ {
		buf[i] = '.'
	}
	var rows [][]string

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			rows = append(rows, append([]string{}, row...))
			return
		}
		for j := 0; j < n; j++ {
			a1 := idx - j
			a2 := idx + j
			if col[j] || as1[a1] || as2[a2] {
				continue
			}
			col[j], as1[a1], as2[a2] = true, true, true
			buf[j] = 'Q'
			row[idx] = string(buf)
			buf[j] = '.'
			dfs(idx + 1)
			col[j], as1[a1], as2[a2] = false, false, false
		}
	}
	dfs(0)
	return rows
}
