package main

import "sort"

// https://leetcode-cn.com/problems/combination-sum-ii/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var dfs func(cur, left int)
	var nums []int
	// left 限定为组合
	// map有可能变成 [1,7] [7,1]
	dfs = func(cur, left int) {
		if cur == 0 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := left; i < len(candidates); i++ {
			// 除非首发 否则相同的只用一次
			// 类似left 限定为组合

			if cur < candidates[i] || left != i && candidates[i-1] == candidates[i] {
				continue
			}
			nums = append(nums, candidates[i])
			dfs(cur-candidates[i], i+1)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(target, 0)
	return res
}
