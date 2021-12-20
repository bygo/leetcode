package main

// https://leetcode-cn.com/problems/distinct-subsequences

func numDistinct(s string, t string) int {
	l1, l2 := len(s), len(t)
	f := make([][]int, l1+1)
	for i := range f {
		f[i] = make([]int, l2+1)
		f[i][0] = 1
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			f[i][j] = f[i-1][j]   // 少一个字符 的路径数
			if s[i-1] == t[j-1] { // 如果相等 加 上次普通路径数
				f[i][j] += f[i-1][j-1]
			}
		}
	}

	return f[l1][l2]
}
