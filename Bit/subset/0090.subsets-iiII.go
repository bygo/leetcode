package main

import "sort"

// https://leetcode.cn/problems/subsets-ii

//func subsetsWithDup(nums []int) [][]int {
//	var combNums [][]int
//	sort.Ints(nums)
//	numsL := len(nums)
//	subsetMax := 1 << numsL
//	for subset := 0; subset < subsetMax; subset++ {
//		var numsCur []int
//		var idx int
//		for ; idx < numsL; idx++ {
//			if subset<<idx&1 == 0 {
//				continue
//			}
//			if idx < numsL-1 && subset<<(idx+1)&1 == 0 && nums[idx] == nums[idx+1] {
//				break
//			}
//			numsCur = append(numsCur, nums[idx])
//		}
//		if idx == numsL {
//			combNums = append(combNums, numsCur)
//		}
//	}
//	return combNums
//}

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
			if 0 < idx && subset>>(idx-1)&1 == 0 && nums[idx] == nums[idx-1] {
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
