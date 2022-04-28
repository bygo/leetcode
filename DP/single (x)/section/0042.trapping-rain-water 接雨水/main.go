package main

// https://leetcode-cn.com/problems/trapping-rain-water

// pre
// f(i) = min(preMax,sufMax) - i

func trap(height []int) int {
	hL := len(height)

	preMax := make([]int, hL)
	sufMax := make([]int, hL)

	preMax[0] = height[0]
	sufMax[hL-1] = height[hL-1]
	for i := 1; i < hL; i++ {
		preMax[i] = max(preMax[i-1], height[i])
		sufMax[hL-i-1] = max(sufMax[hL-i], height[hL-i-1])
	}

	var sum int
	for i := 0; i < hL; i++ {
		sum += min(preMax[i], sufMax[i]) - height[i]
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
