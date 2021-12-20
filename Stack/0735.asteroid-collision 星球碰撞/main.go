package main

// https://leetcode-cn.com/problems/asteroid-collision

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, v := range asteroids {
		if 0 < v { // 正数栈
			stack = append(stack, v)
		} else {
			top := len(stack) - 1
			p2 := abs(v)

			for -1 < top && 0 < stack[top] && stack[top] < p2 { // 负的大, 撞
				stack = stack[:top]
				top--
			}

			if -1 == top || stack[top] < 0 {
				stack = append(stack, v)
			} else if stack[top] == p2 {
				stack = stack[:top]
			}
		}
	}

	return stack
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
