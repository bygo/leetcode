package main

// https://leetcode-cn.com/problems/n-queens-ii

func totalNQueens(n int) int {
	cols, dias1, dias2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
	var dfs func(row int)
	var res int
	dfs = func(row int) {
		if row == n {
			res++
			return
		}
		for col := 0; col < n; col++ {
			dia1 := col - row
			dia2 := col + row
			if !cols[col] && !dias1[dia1] && !dias2[dia2] {
				cols[col], dias1[dia1], dias2[dia2] = true, true, true
				dfs(row + 1)
				cols[col], dias1[dia1], dias2[dia2] = false, false, false
			}
		}
	}
	dfs(0)
	return res
}
