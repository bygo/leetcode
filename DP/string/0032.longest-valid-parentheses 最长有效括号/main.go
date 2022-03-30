package main

// https://leetcode-cn.com/problems/longest-valid-parentheses

// f(n) = f(n-2) + 2 || f(n-1) + 2 + f(n-f(n-1)-2)

func longestValidParentheses(s string) int {
	sL := len(s)
	var cntMax int
	f := make([]int, sL)
	for idx := 1; idx < sL; idx++ {
		if s[idx] == ')' {
			if s[idx-1] == '(' {
				if 2 <= idx {
					f[idx] = f[idx-2] + 2
				} else {
					f[idx] = 2
				}
			} else {
				idxLeft := idx - f[idx-1] - 1
				if 0 <= idxLeft && s[idxLeft] == '(' {
					if 2 <= idxLeft {
						f[idx] = f[idxLeft-1] + 2 + f[idx-1]
					} else {
						f[idx] = f[idx-1] + 2
					}
				}
			}
			if cntMax < f[idx] {
				cntMax = f[idx]
			}
		}
	}
	return cntMax
}
