package main

// https://leetcode.cn/problems/subsets

func subsets(nums []int) [][]int {
	var subsetNums [][]int
	var numsCur []int
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			subsetNums = append(subsetNums, append([]int{}, numsCur...))
			return
		}

		numsCur = append(numsCur, nums[idx])
		dfs(idx + 1) // f1
		numsCur = numsCur[:len(numsCur)-1]

		// 忽略本次 // f2
		dfs(idx + 1)
	}
	dfs(0)
	return subsetNums
}
