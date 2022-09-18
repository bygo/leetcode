package main

// https://leetcode.cn/problems/jump-game

func canJump(nums []int) bool {
	var idxFar int
	for idx := range nums {
		if idxFar < idx { // TODO 跳不过
			return false
		}
		idxFar = max(idxFar, idx+nums[idx])
	}
	return true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
