package main

import "sort"

// https://leetcode.cn/problems/matchsticks-to-square

func makesquare(matchsticks []int) bool {
	total := 0
	for _, matchstick := range matchsticks {
		total += matchstick
	}
	if total%4 != 0 {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // TODO 从大到小
	cnt := total / 4

	mL := len(matchsticks)
	var edges [4]int
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == mL {
			return true
		}

		for i := range edges {
			edges[i] += matchsticks[idx]
			if edges[i] <= cnt && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}
