package main

// https://leetcode-cn.com/problems/n-queens

func solveNQueens(n int) [][]string {
	cols, dias1, dias2 := make([]bool, n), map[int]bool{}, map[int]bool{}

	var res [][]string
	var cur = make([]string, n)
	var b = make([]byte, n)
	for i := range b {
		b[i] = '.'
	}

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			tmp := make([]string, n)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		for col := 0; col < n; col++ {
			dia1, dia2 := row-col, row+col
			if !cols[col] && !dias1[dia1] && !dias2[dia2] {
				cols[col], dias1[dia1], dias2[dia2] = true, true, true // 尝试

				// 记录
				b[col] = 'Q'
				cur[row] = string(b)
				b[col] = '.'
				dfs(row + 1)

				cols[col] = false // 重置
				delete(dias1, dia1)
				delete(dias2, dia2)
			}
		}
	}
	dfs(0)
	return res
}
