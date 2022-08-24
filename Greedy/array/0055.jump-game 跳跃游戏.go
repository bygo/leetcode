package array

// https://leetcode.cn/problems/jump-game

func canJump(nums []int) bool {
	var idxFar int
	for idx := range nums {
		if idx <= idxFar {
			idxFar = max(idxFar, idx+nums[idx])
		}

	}
	return len(nums)-1 < idxFar
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
