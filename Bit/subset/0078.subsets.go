package main

// https://leetcode.cn/problems/subsets

func subsets(nums []int) [][]int {
	var subsetNums [][]int
	numsL := len(nums)
	subsetMax := 1 << numsL
	for subset := 0; subset < subsetMax; subset++ { // 总共 0 ~ 1<<numsL - 1 种
		var numsCur []int
		for idx := 0; idx < numsL; idx++ { //
			if subset>>idx&1 == 1 { // used
				numsCur = append(numsCur, nums[idx])
			}
		}
		subsetNums = append(subsetNums, numsCur)
	}
	return subsetNums
}
