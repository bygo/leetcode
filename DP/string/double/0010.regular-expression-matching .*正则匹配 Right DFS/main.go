package main

// https://leetcode.cn/problems/regular-expression-matching

func isMatch(s string, p string) bool {
	sL, pL := len(s), len(p)
	// 空字符匹配空字符
	if pL == 0 {
		return sL == 0
	}
	if 2 <= pL && p[1] == '*' {
		if isMatch(s, p[2:]) {
			return true
		}
		// 贪心
		return 0 < sL && (p[0] == s[0] || p[0] == '.') && isMatch(s[1:], p)
	} else {
		// 普通
		return 0 < sL && (p[0] == s[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}
}
