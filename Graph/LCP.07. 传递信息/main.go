package main

// https://leetcode-cn.com/problems/chuan-di-xin-xi/

func numWays(n int, relation [][]int, k int) int {
	graph := make([][]int, n)
	for _, v := range relation {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	var res int
	var dfs func(x, step int)
	dfs = func(x, step int) {
		if k == step {
			if x == n-1 {
				res++
			}
			return
		}
		for _, v := range graph[x] {
			dfs(v, step+1)
		}
	}
	dfs(0, 0)
	return res
}
