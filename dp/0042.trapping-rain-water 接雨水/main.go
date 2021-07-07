package main

// Link: https://leetcode-cn.com/problems/trapping-rain-water

// 前缀
// f(n) = min(leftMax,rightMax) - self
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0] // 第一个是本身
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1] // 最后一个是本身
	for i := n - 2; 0 <= i; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var res int
	for i, h := range height {
		res += min(leftMax[i], rightMax[i]) - h // 水
	}
	return res
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
