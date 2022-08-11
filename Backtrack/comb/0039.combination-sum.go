package main

// https://leetcode.cn/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	//sort.Ints(candidates)
	var combNums [][]int
	var dfs func(val, idxLeft int)
	var nums []int

	dfs = func(val, idxLeft int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, nums...))
			return
		}

		cL := len(candidates)
		for idx := idxLeft; idx < cL; idx++ {
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
