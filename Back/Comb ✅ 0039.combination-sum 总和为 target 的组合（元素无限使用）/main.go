package main

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	//sort.Ints(candidates)
	var res [][]int
	var dfs func(val, left int)
	var cur []int

	// left 限定为组合
	dfs = func(val, left int) {
		if val == 0 {
			res = append(res, append([]int{}, cur...))
			return
		}

		for i := left; i < len(candidates); i++ {
			if val < candidates[i] {
				continue
				//break
			}
			cur = append(cur, candidates[i])
			dfs(val-candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0)
	return res
}
