package main

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	//sort.Ints(candidates)
	var res [][]int
	var dfs func(cur, left int)
	var nums []int
	// left 限定为组合
	dfs = func(cur, left int) {
		if cur == 0 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}

		for i := left; i < len(candidates); i++ {
			if cur < candidates[i] {
				continue
			}
			nums = append(nums, candidates[i])
			dfs(cur-candidates[i], i)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(target, 0)
	return res
}
