package main

// https://leetcode.cn/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	//sort.Ints(candidates)
	var combNums [][]int
	var nums []int
	var dfs func(val, idxLeft int)

	dfs = func(val, idxLeft int) {
		if val < 0 {
			return
		}
		if val == 0 {
			combNums = append(combNums, append([]int{}, nums...))
			return
		}

		cL := len(candidates)
		for idx := idxLeft; idx < cL; idx++ {
			nums = append(nums, candidates[idx])
			dfs(val-candidates[idx], idx) // 'idx' can be used indefinitely
			nums = nums[:len(nums)-1]
		}
	}
	dfs(target, 0)
	return combNums
}