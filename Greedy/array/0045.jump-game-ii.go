package main

// https://leetcode.cn/problems/jump-game-ii

func jump(nums []int) int {
	var steps, idxCur, idxFar int
	top := len(nums) - 1
	for idx := 0; idx < top; idx++ {
		idxFar = max(idxFar, idx+nums[idx]) // TODO 最远
		if idxCur == idx {                  // TODO 必须跳跃,最后一次不需要跳跃
			idxCur = idxFar
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
