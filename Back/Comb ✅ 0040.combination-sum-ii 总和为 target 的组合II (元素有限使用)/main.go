package main

import "sort"

// https://leetcode-cn.com/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var dfs func(val, left int)
	var cur []int
	// left 限定为组合
	// 例如路径 [1,7] [7,1] 重复
	dfs = func(val, left int) {
		if val == 0 {
			res = append(res, append([]int{}, cur...))
			return
		}

		for i := left; i < len(candidates); i++ {
			// 除非首发 否则相同的只用一次，剪枝相同路口

			if val < candidates[i] || left < i && candidates[i-1] == candidates[i] {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(val-candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0)
	return res
}
