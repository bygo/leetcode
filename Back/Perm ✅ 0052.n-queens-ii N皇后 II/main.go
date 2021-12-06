package main

// https://leetcode-cn.com/problems/n-queens-ii

func totalNQueens(n int) int {
	cols, ds1, ds2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
	var res int
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res++
			return
		}

		for col := 0; col < n; col++ {
			d1 := col - row
			d2 := col + row
			if !cols[col] && !ds1[d1] && !ds2[d2] {
				cols[col], ds1[d1], ds2[d2] = true, true, true
				dfs(row + 1)
				cols[col], ds1[d1], ds2[d2] = false, false, false
			}
		}
	}
	dfs(0)
	return res
}
