package main

import "sort"

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

func permutation(s string) []string {
	var res []string
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	perm := make([]byte, 0, n)
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			res = append(res, string(perm))
			return
		}
		for j, b := range vis {
			if b || 0 < j && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			dfs(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	dfs(0)
	return res
}
