package main

// https://leetcode-cn.com/problems/wildcard-matching

// f[i][j] = f[i-1][j] || f[i][j-1] || f[i-1][j-1]
// ? 单个字符
// * 任意字符

func isMatch(s string, p string) bool {
	sL := len(s)
	pL := len(p)

	f := make([][]bool, sL+1)
	for i := range f {
		f[i] = make([]bool, pL+1)
	}
	f[0][0] = true
	for j := 1; j <= pL && p[j-1] == '*'; j++ {
		f[0][j] = true
	}

	for i := 1; i <= sL; i++ {
		f[i][0] = false
		for j := 1; j <= pL; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j-1] || f[i-1][j]
			} else if p[j-1] == s[i-1] || p[j-1] == '?' {
				f[i][j] = f[i-1][j-1]
			}
		}
	}

	return f[sL][pL]
}
