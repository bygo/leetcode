package main

// https://leetcode.cn/problems/maximum-nesting-depth-of-the-parentheses

// ❓ 括号的最大嵌套深度

func maxDepth(s string) int {
	var cnt int
	var cntMax int
	for i := range s {
		ch := s[i]
		if ch == ')' {
			cnt--
		} else if ch == '(' {
			cnt++
			if cntMax < cnt {
				cntMax = cnt
			}
		}
	}
	return cntMax
}
