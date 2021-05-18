package main

import "strconv"

//Link: https://leetcode-cn.com/problems/evaluate-reverse-polish-notation

func evalRPN(tokens []string) int {
	var stack []int
	for _, t := range tokens {
		if t == "+" {
			stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if t == "-" {
			stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if t == "*" {
			stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if t == "/" {
			stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			c, _ := strconv.Atoi(t)
			stack = append(stack, c)
		}
	}
	return stack[0]
}
