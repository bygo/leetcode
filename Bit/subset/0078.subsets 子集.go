package main

// https://leetcode-cn.com/problems/subsets

// ❓ 子集
// ⚠️ 1 <= len(nums) <= 10

func subsets(nums []int) [][]int {
	var subsetNums [][]int
	numsL := len(nums)
	for subset := 0; subset < 1<<numsL; subset++ { // 总共 0 ~ 1<<numsL - 1 种
		var numsCur []int
		for idx := 0; idx < numsL; idx++ { //
			if subset>>idx&1 == 1 { // 对应的idx 是否被使用
				numsCur = append(numsCur, nums[idx])
			}
		}
		subsetNums = append(subsetNums, numsCur)
	}
	return subsetNums
}
