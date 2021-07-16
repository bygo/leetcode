package main

import "strconv"

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation

func evalRPN(tokens []string) int {
	var stack []int
	for _, t := range tokens {
		l := len(stack)
		if t == "+" {
			stack[l-2] = stack[l-2] + stack[l-1]
			stack = stack[:l-1]
		} else if t == "-" {
			stack[l-2] = stack[l-2] - stack[l-1]
			stack = stack[:l-1]
		} else if t == "*" {
			stack[l-2] = stack[l-2] * stack[l-1]
			stack = stack[:l-1]
		} else if t == "/" {
			stack[l-2] = stack[l-2] / stack[l-1]
			stack = stack[:l-1]
		} else {
			c, _ := strconv.Atoi(t)
			stack = append(stack, c)
		}
	}
	return stack[0]
}
