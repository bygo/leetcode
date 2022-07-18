package main

// https://leetcode.cn/problems/n-queens-ii

func totalNQueens(n int) int {
	cols, ds1, ds2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
	var cnt int
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n {
			cnt++
			return
		}
		for j := 0; j < n; j++ {
			d1 := idx - j
			d2 := idx + j
			if cols[j] || ds1[d1] || ds2[d2] {
				continue
			}
			cols[j], ds1[d1], ds2[d2] = true, true, true
			dfs(idx + 1)
			cols[j], ds1[d1], ds2[d2] = false, false, false
		}
	}
	dfs(0)
	return cnt
}
