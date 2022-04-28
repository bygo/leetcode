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

		for i := start; i < len(candidates); i++ {
			if val < candidates[i] {
				continue
				//break
			}
			// back
			nums = append(nums, candidates[i])
			dfs(val-candidates[i], i)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(target, 0)
	return combNums
}
