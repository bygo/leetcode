package main

// https://leetcode.cn/problems/distinct-subsequences

// 选择或不选择 f[i][j] = f[i-1][j] + f[i-1][j-1]

func numDistinct(s string, t string) int {
	sL, tL := len(s), len(t)
	f := make([][]int, sL+1)
	for i := range f {
		f[i] = make([]int, tL+1)
		f[i][0] = 1
	}

	for i := 1; i <= sL; i++ {
		for j := 1; j <= tL; j++ {
			// 不选择当前字符 的 路径数
			f[i][j] += f[i-1][j]

			// 选择当前字符 的 路径数 ，必须相等
			if s[i-1] == t[j-1] {
				// 相当于消元
				f[i][j] = f[i-1][j-1]
			}
		}
	}

	return f[sL][tL]
}
