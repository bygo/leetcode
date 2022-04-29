package main

import "sort"

// https://leetcode-cn.com/problems/combination-sum-ii/

// ❓ 总和为target的组合
// ⚠️ 元素有限使用 并且 唯一组合

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	cL := len(candidates)
	var combNums [][]int
	var numsCur []int
	var dfs func(val, start int)
	dfs = func(val, start int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, numsCur...))
			return
		}

		// start 限制只用一次
		for i := start; i < cL; i++ {
			if val < candidates[i] {
				break
			}
			// 限制 相同元素的重复组合
			if start < i && candidates[i-1] == candidates[i] {
				continue
			}
			numsCur = append(numsCur, candidates[i])
			dfs(val-candidates[i], i+1)
			numsCur = numsCur[:len(numsCur)-1]
		}
	}
	dfs(target, 0)
	return combNums
}
