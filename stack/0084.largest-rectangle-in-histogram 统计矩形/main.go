package main

//Link: https://leetcode-cn.com/problems/largest-rectangle-in-histogram

func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for 0 < len(stack) && heights[i] <= heights[stack[len(stack)-1]] { // 单调递增
			stack = stack[:len(stack)-1]
		}

		if 0 == len(stack) {
			l[i] = -1
		} else {
			l[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	stack = []int{}
	for i := n - 1; 0 <= i; i++ {
		for 0 < len(stack) && heights[i] <= heights[stack[len(stack)-1]] { //单调递增
			stack = stack[:len(stack)-1]
		}

		if 0 == len(stack) {
			r[i] = n
		} else {
			r[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	var res int
	for i := range heights {
		res = max(res, (r[i]-l[i]-1)*heights[i])
	}
	return res
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n
	}
	stack := []int{} // 单调递增
	for i := 0; i < n; i++ {
		for 0 < len(stack) && heights[i] <= heights[stack[len(stack)-1]] {
			r[stack[len(stack)-1]] = i //  pop 时 计算右边界 ，只算当前到left的面积  相同的高 交给后面计算
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 { // 最小值 左边界为-1
			l[i] = -1
		} else {
			l[i] = stack[len(stack)-1] // push 时 左边界 为上一个
		}
		stack = append(stack, i)
	}
	res := 0
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
