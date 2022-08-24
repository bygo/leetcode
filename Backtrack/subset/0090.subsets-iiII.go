package main

import "sort"

// https://leetcode.cn/problems/subsets-ii

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var subsetNums [][]int
	var numsCur []int
	var dfs func(idx int, hasPre bool)
	dfs = func(idx int, hasPre bool) {
		if idx == len(nums) {
			subsetNums = append(subsetNums, append([]int{}, numsCur...))
			return
		}
		// Ignore node // f2
		dfs(idx+1, false)
		if !hasPre && 0 < idx && nums[idx-1] == nums[idx] {
			return
		}

		numsCur = append(numsCur, nums[idx])
		dfs(idx+1, true) // f1
		numsCur = numsCur[:len(numsCur)-1]

	}
	dfs(0, false)
	return subsetNums
}
