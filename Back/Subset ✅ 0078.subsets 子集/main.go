package main

// https://leetcode-cn.com/problems/subsets

func subsets(nums []int) [][]int {
	var subsetNums [][]int
	var numsCur []int
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			subsetNums = append(subsetNums, append([]int{}, numsCur...))
			return
		}

		numsCur = append(numsCur, nums[i])
		dfs(i + 1)
		numsCur = numsCur[:len(numsCur)-1]
		dfs(i + 1)
	}
	dfs(0)
	return subsetNums
}
