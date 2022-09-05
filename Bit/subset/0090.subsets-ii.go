package main

import "sort"

// https://leetcode.cn/problems/subsets-ii

func subsetsWithDup(nums []int) [][]int {
	var combNums [][]int
	numsL := len(nums)
	sort.Ints(nums)
	subsetMax := 1 << numsL
	for subset := 0; subset < subsetMax; subset++ {
		var numsCur []int
		var idx int
		for ; idx < numsL; idx++ {
			if subset>>idx&1 == 0 {
				continue
			}
			// 1
			// 10
			// 110 100
			// 1110 1100 1000
			if 0 < idx && subset>>(idx-1)&1 == 0 && nums[idx] == nums[idx-1] { // TODO
				break
			}
			numsCur = append(numsCur, nums[idx])
		}
		if idx != numsL { // TODO
			continue
		}
		combNums = append(combNums, numsCur)
	}
	return combNums
}
