package main

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	//sort.Ints(candidates)
	var combNums [][]int
	var dfs func(val, start int)
	var nums []int

	dfs = func(val, start int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, nums...))
			return
		}

		for idx := start; idx < len(candidates); idx++ {
			if val < candidates[idx] {
				continue
				//break
			}
			// back
			nums = append(nums, candidates[idx])
			dfs(val-candidates[idx], idx)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(target, 0)
	return combNums
}
