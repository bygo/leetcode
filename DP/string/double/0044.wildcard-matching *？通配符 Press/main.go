package main

// https://leetcode-cn.com/problems/wildcard-matching

// f[i][j] = f[i-1][j] || f[i][j-1] || f[i-1][j-1]
// ? 单个字符
// * 任意字符
func isMatch(s string, p string) bool {
	sL, pL := len(s), len(p)
	f := make([]bool, pL+1)

	// 空字符匹配距离
	f[0] = true
	for j := 1; j <= pL && p[j-1] == '*'; j++ {
		f[j] = true
	}

	for i := 1; i <= sL; i++ {
		// 首个
		angle := f[0]

		// 有字符的正则匹配必定为false
		f[0] = false
		for j := 1; j <= pL; j++ {
			pre := f[j]
			if p[j-1] == '*' {
				// 忽略 | 贪心
				f[j] = f[j-1] || f[j]
				// 普通匹配
			} else {
				f[j] = angle && (p[j-1] == s[i-1] || p[j-1] == '?')
			}
			angle = pre
		}
	}
	return f[pL]
}
