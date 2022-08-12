package main

// https://leetcode.cn/problems/regular-expression-matching

func isMatch(s string, p string) bool {
	sL, pL := len(s), len(p)

	if pL == 0 {
		return sL == 0
	}

	sTop := sL - 1
	pTop := pL - 1
	if p[pTop] == '*' {
		// 忽略
		if isMatch(s, p[:pTop-1]) {
			return true
		}
		// 贪心
		return -1 < sTop && (p[pTop-1] == s[sTop] || p[pTop-1] == '.') && isMatch(s[:sTop], p)
	} else {
		// 普通
		return -1 < sTop && (p[pTop] == s[sTop] || p[pTop] == '.') && isMatch(s[:sTop], p[:pTop])
	}
}
