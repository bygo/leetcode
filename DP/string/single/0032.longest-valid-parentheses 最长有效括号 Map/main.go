package main

// https://leetcode.cn/problems/longest-valid-parentheses

// f(n) = f(n-2) + 2 || f(n-1) + 2 + f(n-f(n-1)-2)

// map
func longestValidParentheses(s string) int {
	var cntMax int
	var sL = len(s)
	var f = map[int]int{}
	for idx := 1; idx < sL; idx++ {
		if s[idx] == ')' {
			if s[idx-1] == '(' {
				f[idx] = f[idx-2] + 2
			} else {
				idxLeft := idx - f[idx-1] - 1
				if 0 <= idxLeft && s[idxLeft] == '(' {
					f[idx] = f[idx-1] + f[idxLeft-1] + 2
				}
			}
			if cntMax < f[idx] {
				cntMax = f[idx]
			}
		}
	}
	return cntMax
}
