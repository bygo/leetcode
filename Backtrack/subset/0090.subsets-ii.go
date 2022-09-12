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
		dfs(idx+1, false)
		if !hasPre && 0 < idx && nums[idx-1] == nums[idx] {
			// 上个没使用
			return
		}

		// TODO hasPre 保证连续相同的值 才可组成子集
		numsCur = append(numsCur, nums[idx])
		dfs(idx+1, true) // f1
		numsCur = numsCur[:len(numsCur)-1]
	}
	dfs(0, false)
	return subsetNums
}
