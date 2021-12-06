package main

import "sort"

// https://leetcode-cn.com/problems/permutations

func permute(raw []int) [][]int {
	var res [][]int
	sort.Slice(raw, func(i, j int) bool { return raw[i] < raw[j] })
	l := len(raw)
	cur := make([]int, l)
	vis := make([]bool, l)
	var dfs func(int)
	dfs = func(i int) {
		if i == l {
			res = append(res, append([]int{}, cur...))
			return
		}
		for j := 0; j < l; j++ {
			// 已使用 或者 前置未使用，且等于当前值
			if vis[j] {
				continue
			}
			vis[j] = true
			cur[i] = raw[j]
			dfs(i + 1)
			vis[j] = false
		}
	}
	dfs(0)
	return res
}
