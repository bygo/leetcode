package main

// can i jump to last point ？
// https://leetcode-cn.com/problems/jump-game

// find farthest point
func canJump(nums []int) bool {
	var point int
	for i, num := range nums {
		if i <= point {
			point = max(point, i+num)
		}
	}
	return len(nums)-1 <= point
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
