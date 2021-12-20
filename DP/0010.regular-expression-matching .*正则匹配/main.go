package main

// . any char
// * pre(0~n) match
// https://leetcode-cn.com/problems/regular-expression-matching

// f[i][j] = f[i][j-2] || f[i-1][j-1]
func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	f := make([][]bool, l1+1)
	for i := range f {
		f[i] = make([]bool, l2+1)
	}

	f[0][0] = true
	for i := 0; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if p[j-1] == '*' {
				if f[i][j-2] {
					f[i][j] = true
				} else if i != 0 && f[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.') {
					f[i][j] = true
				}
			} else if i != 0 && f[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.') {
				f[i][j] = true
			}
		}
	}

	return f[l1][l2]
}

// 递归
func isMatch(s string, p string) bool {
	return dfs(s, p)
}

func dfs(s string, p string) bool {
	m, n := len(s), len(p)
	if n == 0 {
		return m == 0
	}
	if 2 <= n && p[1] == '*' {
		return dfs(s, p[2:]) || 0 < m && (s[0] == p[0] || p[0] == '.') && dfs(s[1:], p)
	} else {
		return 0 < m && (s[0] == p[0] || p[0] == '.') && dfs(s[1:], p[1:])
	}
}
