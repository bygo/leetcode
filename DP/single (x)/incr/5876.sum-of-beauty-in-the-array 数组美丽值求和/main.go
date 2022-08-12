package main

// https://leetcode.cn/problems/sum-of-beauty-in-the-array

//
func sumOfBeauties(nums []int) int {
	numsL := len(nums)
	var preMax = make([]int, numsL)
	var sufMin = make([]int, numsL)
	var min = 1<<63 - 1
	var max = -1 << 63

	for i := 0; i < numsL; i++ {
		if max < nums[i] {
			max = nums[i]
		}
		preMax[i] = max

		if nums[numsL-i-1] < min {
			min = nums[numsL-i-1]
		}
		sufMin[numsL-i-1] = min
	}

	var score int
	top := numsL - 1
	for i := 1; i < top; i++ {
		if preMax[i-1] < nums[i] && nums[i] < sufMin[i+1] {
			score += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			score += 1
		}
	}
	return score
}
