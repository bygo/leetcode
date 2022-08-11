package main

import "sort"

// https://leetcode.cn/problems/combination-sum-ii/

// ❓ 总和为target的组合
// ⚠️ 元素有限使用 并且 唯一组合

func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	cL := len(nums)
	var combNums [][]int
	var numsCur []int
	var dfs func(val, start int)
	dfs = func(val, start int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, numsCur...))
			return
		}

		// start 限制只用一次
		for i := start; i < cL; i++ {
			if val < nums[i] {
				break
			}
			// 限制 相同元素的重复组合
			if start < i && nums[i-1] == nums[i] {
				continue
			}
			numsCur = append(numsCur, nums[i])
			dfs(val-nums[i], i+1)
			numsCur = numsCur[:len(numsCur)-1]
		}
	}
	dfs(target, 0)
	return combNums
}

func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	cL := len(nums)
	var combNums [][]int
	var numsCur []int
	var dfs func(val, start int)
	dfs = func(val, start int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, numsCur...))
			return
		}

		for i := start; i < cL; i++ {
			if val < nums[i] {
				break
			}

			if start < i && nums[i-1] == nums[i] {
				continue
			}
			numsCur = append(numsCur, nums[i])
			dfs(val-nums[i], i+1)
			numsCur = numsCur[:len(numsCur)-1]
		}
	}

	dfs(target, 0)
	return combNums
}
