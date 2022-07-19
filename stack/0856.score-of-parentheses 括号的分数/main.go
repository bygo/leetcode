package main

// https://leetcode.cn/problems/score-of-parentheses

func scoreOfParentheses(s string) int {
	var res, level int
	for i, v := range s {
		if v == '(' {
			level++
		} else {
			if s[i-1] == '(' { // 抵消计算层级
				res += 1 << (level - 1)
			}
			level--
		}
	}
	return res
}
