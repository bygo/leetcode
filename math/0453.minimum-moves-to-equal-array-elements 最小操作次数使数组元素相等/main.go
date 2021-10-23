package main

// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements

func minMoves(nums []int) int {
	var min = nums[0]
	var sum = nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
		sum += num
	}
	return sum - min*len(nums)
}
