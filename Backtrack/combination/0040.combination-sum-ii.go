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

		for idx := idxLeft; idx < cL; idx++ {
			if val < nums[idx] {
				break
			}
			if idxLeft < idx && nums[idx-1] == nums[idx] { // TODO 重复判断
				continue
			}

			numsCur = append(numsCur, nums[idx])
			dfs(val-nums[idx], idx+1) // TODO idx只能使用一次
			numsCur = numsCur[:len(numsCur)-1]
		}
	}
	dfs(target, 0)
	return combNums
}
