package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res) //深度优先
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { //结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if target < candidates[i] { //剪枝
			return
		}
		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i, res) //分支
	}
}
