package main

// Link: https://leetcode-cn.com/problems/wildcard-matching

// 二维
func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	f := make([][]bool, l1+1)
	for i := range f {
		f[i] = make([]bool, l2+1)
	}
	f[0][0] = true
	for j := 1; j <= l2; j++ {
		if p[j-1] != '*' {
			break
		}
		f[0][j] = true
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i-1][j] || f[i][j-1]
			} else if s[i-1] == p[j-1] || '?' == p[j-1] {
				f[i][j] = f[i-1][j-1]
			}
		}
	}
	return f[l1][l2]
}
