package main

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/

func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	cL := len(nums)
	var combNums [][]int
	var numsCur []int
	var dfs func(val, idxLeft int)
	dfs = func(val, idxLeft int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, numsCur...))
			return
		}

		for idx := idxLeft; idx < cL; idx++ { // 'idxLeft' restrict `num` to be used only once
			if val < nums[idx] {
				break
			}
			if idxLeft < idx && nums[idx-1] == nums[idx] { // restrict repeating elements
				continue
			}

			numsCur = append(numsCur, nums[idx])
			dfs(val-nums[idx], idx+1) // `idx+1` restrict `num` to be used only once
			numsCur = numsCur[:len(numsCur)-1]
		}
	}
	dfs(target, 0)
	return combNums
}
