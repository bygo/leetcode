package main

// Link: https://leetcode-cn.com/problems/largest-rectangle-in-histogram

// 找左边界与右边界,只统计当前高度的最大延展宽度
func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n
	}
	stack := []int{} // 单调递增
	for i := 0; i < n; i++ {
		for 0 < len(stack) && heights[i] <= heights[stack[len(stack)-1]] {
			r[stack[len(stack)-1]] = i //  pop 时 计算右边界
			stack = stack[:len(stack)-1]
		}
		if 0 == len(stack) { // 最小值 左边界为-1
			l[i] = -1
		} else {
			l[i] = stack[len(stack)-1] // push 时 左边界 为上一个
		}
		stack = append(stack, i)
	}

	var res int
	for i := 0; i < n; i++ {
		res = max(res, (r[i]-l[i]-1)*heights[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
