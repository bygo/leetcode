package main

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	//sort.Ints(candidates)
	var combNums [][]int
	var dfs func(val, left int)
	var nums []int

	// left 限定为组合
	dfs = func(val, left int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, nums...))
			return
		}

		for i := left; i < len(candidates); i++ {
			if val < candidates[i] {
				continue
				//break
			}
			nums = append(nums, candidates[i])
			dfs(val-candidates[i], i)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(target, 0)
	return combNums
}
