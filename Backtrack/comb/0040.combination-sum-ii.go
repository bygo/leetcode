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
	var dfs func(val, idxLeft int)
	dfs = func(val, idxLeft int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, numsCur...))
			return
		}

		// idx0 限制只用一次
		for i := idxLeft; i < cL; i++ {
			if val < nums[i] {
				break
			}
			// 限制 相同元素的重复组合
			if idxLeft < i && nums[i-1] == nums[i] {
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
	var dfs func(val, idxLeft int)
	dfs = func(val, idxLeft int) {
		if val == 0 {
			combNums = append(combNums, append([]int{}, numsCur...))
			return
		}

		for idx := idxLeft; idx < cL; idx++ {
			if val < nums[idx] {
				break
			}

			if idxLeft < idx && nums[idx-1] == nums[idx] {
				continue
			}
			numsCur = append(numsCur, nums[idx])
			dfs(val-nums[idx], idx+1)
			numsCur = numsCur[:len(numsCur)-1]
		}
	}

	dfs(target, 0)
	return combNums
}
