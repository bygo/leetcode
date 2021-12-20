package main

// https://leetcode-cn.com/problems/loud-and-rich

func loudAndRich(richer [][]int, quiet []int) []int {
	quietL := len(quiet)
	g := make([][]int, quietL)
	for _, r := range richer {
		g[r[1]] = append(g[r[1]], r[0])
	}

	res := make([]int, quietL)
	for i := range res {
		res[i] = -1
	}
	var dfs func(int)
	dfs = func(x int) {
		if res[x] != -1 {
			return
		}
		res[x] = x
		for _, y := range g[x] {
			dfs(y)
			if quiet[res[y]] < quiet[res[x]] {
				res[x] = res[y]
			}
		}
	}
	for i := 0; i < quietL; i++ {
		dfs(i)
	}
	return res
}
