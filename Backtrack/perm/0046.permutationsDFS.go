package main

// https://leetcode.cn/problems/permutations

func permute(raw []int) [][]int {
	var permNums [][]int
	rawL := len(raw)
	cur := make([]int, rawL)
	vis := make([]bool, rawL)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == rawL {
			permNums = append(permNums, append([]int{}, cur...))
			return
		}
		for j := 0; j < rawL; j++ {
			// 已使用
			if vis[j] {
				continue
			}
			vis[j] = true
			cur[idx] = raw[j]
			dfs(idx + 1)
			vis[j] = false
		}
	}
	dfs(0)
	return permNums
}
