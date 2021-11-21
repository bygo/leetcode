package main

// https://leetcode-cn.com/problems/subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}

		cur = append(cur, nums[i])
		dfs(i + 1)
		cur = cur[:len(cur)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}

// 每一位
func subsets(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	for i := 0; i < 1<<n; i++ {
		cur := []int{}
		for j := range nums {
			if i>>j&1 == 1 {
				cur = append(cur, nums[j])
			}
		}
		res = append(res, cur)
	}
	return res
}
