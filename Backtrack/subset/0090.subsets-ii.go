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

func subsetsWithDup(nums []int) [][]int {
	numsL := len(nums)
	sort.Ints(nums)
	var subsetsNums [][]int
	var numsCur []int
	var dfs func(idx int, hasPre bool)
	dfs = func(idx int, hasPre bool) {
		if idx == numsL {
			subsetsNums = append(subsetsNums, append([]int(nil), numsCur...))
			return
		}

		dfs(idx+1, false)
		if !hasPre && 0 < idx && nums[idx-1] == nums[idx] {
			return
		}

		numsCur = append(numsCur, nums[idx])
		dfs(idx+1, true)
		numsCur = numsCur[:len(numsCur)-1]
	}

	dfs(0, false)
	return subsetsNums

}
