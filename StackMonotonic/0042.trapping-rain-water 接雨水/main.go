package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/trapping-rain-water

// 单调递减栈，出栈时 统计当前柱子  横向最大水量
func trap(height []int) int {
	var res int
	var stack []int
	for index, v := range height {
		// 单调递减栈
		for 0 < len(stack) && height[stack[len(stack)-1]] < v { // 大于栈顶
			cur := stack[len(stack)-1] // 洼地
			stack = stack[:len(stack)-1]
			if 0 == len(stack) { // 没有 左挡板了  直接结束
				break
			}
			left := stack[len(stack)-1]             // 左挡板
			w := index - left - 1                   // 宽度
			h := min(height[left], v) - height[cur] // 高度
			res += w * h
		}
		stack = append(stack, index) // 设置挡板
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
