package main

import "sort"

// https://leetcode-cn.com/problems/subsets-ii

// ❓ 子集II
// ⚠️ 重复数字 | 不算重复子集

func subsetsWithDup(nums []int) [][]int {
	var combNums [][]int
	sort.Ints(nums)
	numsL := len(nums)
	subsetMax := 1 << numsL
	for subset := 0; subset < subsetMax; subset++ {
		var numsCur []int
		idx := 0
		for ; idx < numsL; idx++ {
			if subset>>idx&1 == 0 {
				continue
			}
			// 是否选择
			if 0 < idx && subset>>(idx-1)&1 == 0 && nums[idx] == nums[idx-1] {
				// 上一个相同的数没有选择，本次也不选择
				break
			}
			numsCur = append(numsCur, nums[idx])
		}
		if idx == numsL {
			combNums = append(combNums, numsCur)
		}
	}
	return combNums
}
