package main

// https://leetcode.cn/problems/combinations

func combine(numMax int, needL int) [][]int {
	var combNums [][]int
	var nums []int
	var dfs func(num int)
	dfs = func(num int) {
		curL := len(nums)
		if curL == needL {
			combNums = append(combNums, append([]int{}, nums...))
			return
		}
		if numMax-num < needL-curL {
			return
		}
		nums = append(nums, num+1)
		dfs(num + 1)
		nums = nums[:curL]
		dfs(num + 1)
	}
	dfs(0)
	return combNums
}
