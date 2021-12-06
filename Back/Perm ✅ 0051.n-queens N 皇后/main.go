package main

// https://leetcode-cn.com/problems/n-queens

func solveNQueens(n int) [][]string {
	cols, ds1, ds2 := make([]bool, n), map[int]bool{}, map[int]bool{}
	var b = make([]byte, 0, n)
	for i := 0; i < n; i++ {
		b = append(b, '.')
	}
	var cur = make([]string, n)
	var res [][]string

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res = append(res, append([]string{}, cur...))
			return
		}
		for col := 0; col < n; col++ {
			d1 := row - col
			d2 := row + col
			if !cols[col] && !ds1[d1] && !ds2[d2] {
				cols[col], ds1[d1], ds2[d2] = true, true, true
				b[col] = 'Q'
				cur[row] = string(b)
				b[col] = '.'
				dfs(row + 1)
				cols[col], ds1[d1], ds2[d2] = false, false, false
			}
		}
	}
	dfs(0)
	return res
}
