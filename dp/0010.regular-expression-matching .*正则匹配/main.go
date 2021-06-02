package main

//Link: https://leetcode-cn.com/problems/regular-expression-matching

// 递归
func isMatch(s string, p string) bool {
	return dfs(s, p)
}

func dfs(s string, p string) bool {
	m, n := len(s), len(p)
	if n == 0 {
		return m == 0
	}
	if 2 <= n && p[1] == '*' { //长度大于2，获取下一个元素且为*
		//println(m,n)
		//1.去除x*（p[2:]）,继续递归匹配
		//2.去除字符(s[1:]),继续递归匹配
		return dfs(s, p[2:]) || 0 < m && (s[0] == p[0] || p[0] == '.') && dfs(s[1:], p) //只要一种情况成功，计算匹配成功
	} else {
		//普通匹配
		return 0 < m && (s[0] == p[0] || p[0] == '.') && dfs(s[1:], p[1:])
	}
}

// dp

//S = abbbbc
//P = ab*d*c
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if i != 0 && dp[i-1][j] {
					if s[i-1] == p[j-2] || p[j-2] == '.' {
						dp[i][j] = true
					}
				}
			} else if i != 0 && dp[i-1][j-1] {
				if s[i-1] == p[j-1] || p[j-1] == '.' {
					dp[i][j] = true
				}
			}
		}
	}

	return dp[m][n]
}
