package main

// https://leetcode.cn/problems/maximal-rectangle

func maximalRectangle(matrix [][]byte) int {
	if 0 == len(matrix) {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	nums := make([][]int, m)
	for i, row := range matrix {
		nums[i] = make([]int, n)
		for j, b := range row {
			if b == '1' {
				if j == 0 {
					nums[i][j] = 1
				} else {
					nums[i][j] = nums[i][j-1] + 1
				}
			}
		}
	}

	var res int
	for i, row := range matrix {
		for j, b := range row {
			if b == '1' {
				width := nums[i][j]
				res = max(res, width)
				for k := i; 0 <= k; k-- {
					width = min(width, nums[k][j])
					res = max(res, width*(i-k+1))
				}
			}
		}
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

func maximalRectangle(matrix [][]byte) int {
	if 0 == len(matrix) {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	nums := make([][]int, m)
	for i, row := range matrix {
		nums[i] = make([]int, n)
		for j, b := range row {
			if b == '1' {
				if j == 0 {
					nums[i][j] = 1
				} else {
					nums[i][j] = nums[i][j-1] + 1
				}
			}
		}
	}

	var res int

	for j := 0; j < n; j++ { // 每一列循环
		top, bottom := make([]int, m), make([]int, m)
		for i := 0; i < m; i++ {
			bottom[i] = m // 最长 等于数组数
		}
		stack := []int{}
		for i := 0; i < m; i++ {
			for 0 < len(stack) && nums[i][j] <= nums[stack[len(stack)-1]][j] {
				bottom[stack[len(stack)-1]] = i // 出栈时 确定右边界
				stack = stack[:len(stack)-1]
			}
			if 0 == len(stack) { // 入栈时 确认左边界
				top[i] = -1
			} else {
				top[i] = stack[len(stack)-1]
			}

			stack = append(stack, i)
		}

		for i := 0; i < m; i++ {
			res = max(res, nums[i][j]*(bottom[i]-top[i]-1))
		}
	}

	return res
}
