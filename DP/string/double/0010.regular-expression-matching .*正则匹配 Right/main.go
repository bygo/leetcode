package main

// https://leetcode.cn/problems/regular-expression-matching

func isMatch(s string, p string) bool {
	sL, pL := len(s), len(p)

	f := make([][]bool, sL+1)
	for i := range f {
		f[i] = make([]bool, pL+1)
	}
	f[sL][pL] = true

	for i := sL; 0 <= i; i-- {
		for j := pL - 1; 0 <= j; j-- {
			// 以x*为整体， x为关联关系
			if j <= pL-2 && p[j+1] == '*' {
				if f[i][j+2] {
					f[i][j] = true
				} else if i < sL && f[i+1][j] && (p[j] == s[i] || p[j] == '.') {
					f[i][j] = true
				}
			} else if i < sL && f[i+1][j+1] && (p[j] == s[i] || p[j] == '.') {
				f[i][j] = true
			}
		}
	}
	return f[0][0]
}
