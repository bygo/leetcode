package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var dfs func(nums []int, target, left int)
	dfs = func(nums []int, target, left int) {
		if target == 0 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := left; i < len(candidates); i++ {
			if target < candidates[i] {
				return
			}
			dfs(append(nums, candidates[i]), target-candidates[i], i)
		}
	}
	dfs(nil, target, 0)
	return res
}
