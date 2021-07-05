package main

// Link: https://leetcode-cn.com/problems/daily-temperatures

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	var res = make([]int, len(temperatures))
	for k, v := range temperatures {
		top := len(stack) - 1
		for -1 < top && temperatures[stack[top]] < v {
			res[stack[top]] = k - stack[top]
			stack = stack[:top]
			top--
		}

		stack = append(stack, k)
	}

	for _, index := range stack {
		res[index] = 0
	}

	return res
}
