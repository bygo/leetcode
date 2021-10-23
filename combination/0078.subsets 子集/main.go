package main

// https://leetcode-cn.com/problems/subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), cur...))
			return
		}
		cur = append(cur, nums[index])
		dfs(index + 1)
		cur = cur[:len(cur)-1]
		dfs(index + 1)
	}
	dfs(0)
	return res
}

func subsets(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		cur := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				cur = append(cur, v)
			}
		}
		res = append(res, cur)
	}
	return res
}
