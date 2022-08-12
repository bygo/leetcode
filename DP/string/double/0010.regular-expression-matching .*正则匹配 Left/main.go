package main

// https://leetcode.cn/problems/regular-expression-matching

func isMatch(s string, p string) bool {
	sL, pL := len(s), len(p)
	f := make([][]bool, sL+1)
	for i := range f {
		f[i] = make([]bool, pL+1)
	}

	f[0][0] = true
	for i := 0; i <= sL; i++ {
		for j := 1; j <= pL; j++ {
			// 以x*为整体， *为关联关系
			if p[j-1] == '*' {
				// 忽略
				if f[i][j-2] {
					f[i][j] = true
				} else if 0 < i && f[i-1][j] && (p[j-2] == s[i-1] || p[j-2] == '.') {
					// f[i-1][j] 可能是忽略 可能是贪心
					f[i][j] = true
				}
			} else if 0 < i && f[i-1][j-1] && (p[j-1] == s[i-1] || p[j-1] == '.') {
				f[i][j] = true
			}
		}
	}
	return f[sL][pL]
}
